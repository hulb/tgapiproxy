build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags '-extldflags -static -s -w' -o proxy