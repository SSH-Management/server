ARG VARIANT="1.17-bullseye"

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ENV CC=gcc
ENV CXX=g++


RUN apt update && \
	apt upgrade -y && \
	apt install iputils-ping mlocate -y && \
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin && \
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest && \
	go install github.com/ramya-rao-a/go-outline@latest && \
	go install github.com/cweill/gotests/gotests@latest && \
	go install github.com/fatih/gomodifytags@latest && \
	go install github.com/josharian/impl@latest && \
	go install github.com/haya14busa/goplay/cmd/goplay@latest && \
	go install github.com/go-delve/delve/cmd/dlv@latest && \
	go install -v golang.org/x/tools/cmd/guru@latest && \
	go install mvdan.cc/gofumpt@latest && \
	go install -tags 'postgres github file' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	updatedb
