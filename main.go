package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// MyHandler simple handler
type MyHandler struct{}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	fmt.Print("receive incoming request from ", r.RemoteAddr)
	tghost := "https://api.telegram.org"
	resp, err := http.Post(tghost+r.RequestURI, "application/json", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Println(" -> proxyed, cost:", time.Now().Sub(s).Milliseconds(), "ms")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &MyHandler{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8093"
	}

	server := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		Handler:      mux,
	}

	go func() {
		fmt.Println("server listen at ", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)

	<-quit
	if err := server.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("server close")
}
