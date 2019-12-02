FROM alpine:3.7

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && mkdir /app

ADD linux /app

WORKDIR /app