FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME worker
COPY ./dist/worker_linux_amd64/worker app
COPY ./configs/worker.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
