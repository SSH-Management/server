FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

RUN apk update && apk add make gcc && make build

FROM alpine:3.13

WORKDIR /app

COPY --from=build /app/bin .

RUN apk update \
    && apk add bash \
    && mkdir -p /var/log/ssh_management

EXPOSE 8080

CMD ["./ssh_management", "--env", "production", "queue:worker" ]
