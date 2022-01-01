FROM docker.io/golang AS build
COPY . /app
RUN cd /app && make build

FROM docker.io/alpine
RUN apk add ca-certificates
COPY --from=build /app/proxy proxy
RUN chmod +x proxy
ENTRYPOINT ./proxy
