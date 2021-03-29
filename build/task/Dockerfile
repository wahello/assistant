FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME task
COPY ./dist/task_linux_amd64/task app
COPY ./configs/rpc.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
