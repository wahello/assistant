FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME message
COPY ./dist/message_linux_amd64/message app
COPY ./configs/rpc.yml config.yml
CMD ["./app", "-f", "config.yml"]
