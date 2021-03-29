FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME gateway
COPY ./dist/gateway_linux_amd64/gateway app
COPY ./configs/gateway.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
