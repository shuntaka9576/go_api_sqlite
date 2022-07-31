FROM golang:alpine AS build-stage
RUN apk add alpine-sdk
ADD . /app
WORKDIR /app
RUN go build -o app .

FROM alpine:latest
COPY --from=build-stage /app/app /usr/local/bin/app
ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz
COPY entrypoint.sh /usr/local/bin
COPY litestream.yml /etc/litestream.yml
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
