package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const telegramHost = "https://api.telegram.org"

var tgUrl *url.URL

func init() {
	var err error
	tgUrl, err = url.Parse(telegramHost)
	if err != nil {
		panic(err)
	}
}

func main() {
	proxy := httputil.NewSingleHostReverseProxy(tgUrl)
	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = tgUrl.Scheme
		r.URL.Host = tgUrl.Host
		r.Host = tgUrl.Host
		r.Header.Del("X-Forwarded-For")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8093"
	}

	log.Println("server listen on ", port)
	if err := http.ListenAndServe(":"+port, proxy); err != nil {
		panic(err)
	}
}
