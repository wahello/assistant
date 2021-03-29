FROM alpine:latest

RUN apk update && apk add bash

ENV APP_NAME cron
COPY ./dist/cron_linux_amd64/cron app
COPY ./configs/cron.yml.example config.yml
CMD ["./app", "-f", "config.yml"]
