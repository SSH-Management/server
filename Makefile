RACE ?= 0
ENV ?= development
VERSION ?= v0.3.0
GOPATH ?= ${HOME}/go
PLATFORM ?= linux/arm64,linux/amd64
DOCKER ?= 0

ifeq ($(DOCKER),1)
	DATABASE_URL="postgres://postgres:postgres@postgres:5432/ssh_management?sslmode=disable&application_name=SSHManagement&TimeZone=UTC"
else
	DATABASE_URL="postgres://postgres:postgres@localhost:5432/ssh_management?sslmode=disable&application_name=SSHManagement&TimeZone=UTC"
endif

CC = gcc
CXX = g++

.PHONY: all
all: clean build


.PHONY: build
build:
ifeq ($(ENV),production)
	@CGO_ENABLED=0 CXX=g++ CC=gcc go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -o ./bin/ssh_management ./*.go
else ifeq ($(ENV),development)
	@CXX=g++ CC=gcc go build -o ./bin/ssh_management -gcflags="all=-N -l" ./*.go
else
	@echo "Target ${ENV} is not supported"
endif
	@cp ssh_management.example.yml ./bin/ssh_management.yml

.PHONY: run
run:
	@CXX=g++ CC=gcc go run ./cmd/*.go

.PHONY: test
test:
ifeq ($(RACE), 1)
	@CC=gcc CXX=g++ go test ./... -race -covermode=atomic -coverprofile=coverage.txt -timeout 5m
else
	@CC=gcc CXX=g++ go test ./... -covermode=atomic -coverprofile=coverage.txt -timeout 1m
endif

ssh_management.yml:
	@cp ssh_management.example.yml bin/ssh_management.yml

.PHONY: tidy
tidy:
	@rm -f go.sum
	@go mod tidy

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: migrate
migrate: install-migrate-cli
	@migrate -source file://$(shell pwd)/migrations -database $(DATABASE_URL) up

M_STEP ?= ""

.PHONY: migrate-down
migrate-down: install-migrate-cli
	@migrate -source file://$(shell pwd)/migrations -database $(DATABASE_URL) down $(M_STEP)

.PHONY: migration-create
migration-create: install-migrate-cli
	@migrate -database $(DATABASE_URL) create -dir ./migrations -seq -ext plsql $(M_NAME)

M_VERSION ?= ""

.PHONY: migration-force
migration-force: install-migrate-cli
	@migrate -database $(DATABASE_URL) -source file://$(shell pwd)/migrations force $(M_VERSION)

.PHONY: install-migrate-cli
install-migrate-cli:
ifneq ($(findstring migrate,$(shell ls $(GOPATH)/bin)),migrate)
	@CC=gcc CXX=g++ cd $(HOME) && go install \
		-tags 'postgres sqlite3 mysql github file' \
		github.com/golang-migrate/migrate/v4/cmd/migrate@latest
endif

.PHONY: buildx
buildx: buildx-server buildx-queue

.PHONY: buildx-server
buildx-server:
	@docker buildx build --target production --build-arg VERSION="$(VERSION)" --platform "$(PLATFORM)" -t "malusevd99/ssh-management:server-$(VERSION)" --push --file ./docker/Dockerfile .

.PHONY: buildx-queue
buildx-queue:
	@docker buildx build --target production --build-arg VERSION="$(VERSION)" --platform "$(PLATFORM)" -t "malusevd99/ssh-management:queue-$(VERSION)" --push --file ./docker/Dockerfile .

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: fmt
fmt:
	@gofumpt -l -w .
