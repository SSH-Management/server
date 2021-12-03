ARG VERSION=

FROM golang:1.17.3-alpine as build

WORKDIR /app

COPY . .

RUN apk update && apk add make gcc && make VERSION=${VERSION} ENV=production  build

FROM ubuntu:20.04

WORKDIR /app

COPY --from=build /app/bin .

RUN apt update \
    && apt upgrade -y \
    && mkdir -p /var/log/ssh_management

EXPOSE 8080

CMD ["./ssh_management", "--env", "production", "queue:worker" ]
