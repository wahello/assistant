FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME web
COPY ./dist/web_linux_amd64/web app
COPY ./configs/web.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
