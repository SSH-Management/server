
ARG VARIANT="1.17-bullseye"
ARG MIGRATE_TAG=

FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

ENV CC=gcc
ENV CXX=g++

RUN apt update \
    && apt install iputils-ping mlocate \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /bin \
	&& cd $(go env GOPATH) && \
	rm -rf $(go env GOPATH)/src/github.com/golang-migrate/migrate && \
	go get -u -d github.com/golang-migrate/migrate/cmd/migrate && \
	cd $(go env GOPATH)/src/github.com/golang-migrate/migrate && \
	git checkout ${MIGRATE_TAG} && \
	cd cmd/migrate && \
	go build -tags 'postgres sqlite3 mysql github file' -ldflags="-X main.Version=${MIGRATE_TAG}" -o $(go env GOPATH)/bin/migrate $(go env GOPATH)/src/github.com/golang-migrate/migrate/cmd/migrate \
	&& updatedb

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
	&& go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
	&& go install github.com/ramya-rao-a/go-outline@latest \
	&& go install github.com/cweill/gotests/gotests@latest \
	&& go install github.com/fatih/gomodifytags@latest \
	&& go install github.com/josharian/impl@latest \
	&& go install github.com/haya14busa/goplay/cmd/goplay@latest \
	&& go install github.com/go-delve/delve/cmd/dlv@master \
	&& go install mvdan.cc/gofumpt@latest
