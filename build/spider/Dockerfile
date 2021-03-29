FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME spider
COPY ./dist/spider_linux_amd64/spider app
COPY ./configs/spider.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
