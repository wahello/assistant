FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME subscribe
COPY ./dist/subscribe_linux_amd64/subscribe app
COPY ./configs/rpc.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
