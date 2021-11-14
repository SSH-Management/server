FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

RUN apk update && apk add make gcc && make build-server

FROM alpine:3.13

WORKDIR /app

COPY ./docker/entrypoint.sh .
COPY --from=build /app/bin .

RUN apk update && apk add bash && mkdir -p /var/log/ssh_management && chmod +x entrypoint.sh

EXPOSE 8080

CMD [ "/app/entrypoint.sh" ]
