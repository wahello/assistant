FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME workflow
COPY ./dist/workflow_linux_amd64/workflow app
COPY ./configs/rpc.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
