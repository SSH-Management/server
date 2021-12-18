FROM golang:1.17-bullseye as development

RUN apt update && \
	apt install iputils-ping mlocate -y && \
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin && \
	updatedb

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
	go install github.com/ramya-rao-a/go-outline@latest && \
	go install github.com/cweill/gotests/gotests@latest && \
	go install github.com/fatih/gomodifytags@latest && \
	go install github.com/josharian/impl@latest && \
	go install github.com/haya14busa/goplay/cmd/goplay@latest && \
	go install github.com/go-delve/delve/cmd/dlv@latest && \
	go install github.com/go-delve/delve/cmd/dlv@master && \
	go install mvdan.cc/gofumpt@latest

WORKDIR /app

COPY . .

CMD ["air", "-c", ".air.server.toml"]

FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

ARG VERSION=

RUN apk add make && make "VERSION=${VERSION}" ENV=production build

FROM ubuntu:20.04 as production

WORKDIR /app

COPY --from=build /app/bin .

RUN apt update && apt upgrade -y && mkdir -p /var/log/ssh_management

CMD ["./ssh_management", "--env", "production", "serve" ]
