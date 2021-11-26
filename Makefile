MIGRATE_TAG = v4.15.1
RACE ?= 0
ENV ?= development
VERSION ?= dev
GOPATH ?= ${HOME}/go
DOCKER ?= 0
TAG ?= 0.0.3
PLATFORM ?= linux/arm64,linux/amd64
ARCH ?= arm64

DATABASE_URL="mysql://server:server@tcp(mysql:3306)/ssh_management?charset=utf8mb4&checkConnLiveness=true&collation=utf8mb4_general_ci&interpolateParams=true&loc=UTC&multiStatements=true&parseTime=true"

CC = gcc
CXX = g++

.PHONY: all
all: clean build


.PHONY: build
build:
ifeq ($(ENV),production)
	CGO_ENABLED=0 CXX=g++ CC=gcc go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ./bin/ssh_management ./cmd/*.go
else ifeq ($(ENV),development)
	CXX=g++ CC=gcc go build -o ./bin/ssh_management ./cmd/*.go
else
	echo "Target ${ENV} is not supported"
endif
	cp ssh_management.example.yml ./bin/ssh_management.yml

.PHONY: run
run:
	CXX=g++ CC=gcc go run ./cmd/*.go

#.PHONY: copy-files
#copy-files: config.yml
#	mkdir -p ./bin/migrations
#	mkdir -p ./bin/public
#ifeq ($(DOCKER),1)
#	cp config.docker.yml ./bin/config.yml
#else
#	cp config.yml ./bin/config.yml
#endif
#	cp -r ./database/migrations ./bin
#	cp -r ./public ./bin/

.PHONY: test
test:
ifeq ($(RACE), 1)
	CC=gcc CXX=g++ go test ./... -race -covermode=atomic -coverprofile=coverage.txt -timeout 5m
else
	CC=gcc CXX=g++ go test ./... -covermode=atomic -coverprofile=coverage.txt -timeout 1m
endif

ssh_management.yml:
	@cp ssh_management.example.yml bin/ssh_management.yml

.PHONY: tidy
tidy:
	rm -f go.sum
	go mod tidy

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: migrate
migrate:
	@migrate -source file://$(shell pwd)/migrations -database $(DATABASE_URL) up

.PHONY: migrate-down
migrate-down:
	@migrate -source file://$(shell pwd)/migrations -database $(DATABASE_URL) down

.PHONY: migrate-create
migration-create:
	@migrate -database $(DATABASE_URL) create -dir ./migrations -seq -ext sql $(M_NAME)

.PHONY: install-migrate-cli
install-migrate-cli:/
ifneq ($(findstring migrate,$(shell ls $(GOPATH)/bin)),migrate)
	CC=gcc CXX=g++ cd $(GOPATH) && \
	rm -rf $(GOPATH)/src/github.com/golang-migrate/migrate && \
	go get -u -d github.com/golang-migrate/migrate/cmd/migrate && \
	cd $(GOPATH)/src/github.com/golang-migrate/migrate && \
	git checkout $(MIGRATE_TAG) && \
	cd cmd/migrate && \
	go build -tags 'postgres sqlite3 mysql github file' -ldflags="-X main.Version=${MIGRATE_TAG}" -o $(GOPATH)/bin/migrate ${GOPATH}/src/github.com/golang-migrate/migrate/cmd/migrate
endif

.PHONY: buildx
buildx:
	docker buildx build --platform "$(PLATFORM)" -t "malusevd99/ssh-management:server-$(TAG)" --push --file ./docker/server/Dockerfile .
	docker buildx build --platform "$(PLATFORM)" -t "malusevd99/ssh-management:queue-$(TAG)" --push --file ./docker/queue/Dockerfile .

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: fmt
fmt:
	@gofumpt -l -w .
