FROM golang AS build
COPY . /app
RUN cd /app && env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -tags netgo main.go

FROM alpine:3.5
RUN apk update && apk add ca-certificates
COPY --from=build /app/main tgapiproxy
RUN chmod +x tgapiproxy
ENTRYPOINT ./tgapiproxy
