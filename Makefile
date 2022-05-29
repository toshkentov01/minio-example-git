CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

IMG_NAME=${APP}
REGISTRY=${REGISTRY:-861701250313.dkr.ecr.us-east-1.amazonaws.com}
TAG=latest
ENV_TAG=latest

ifneq (,$(wildcard ./.env))
	include .env
endif

ifdef CI_COMMIT_BRANCH
        include .build_info
endif

make create-env:
	cp ./.env.example ./.env

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

update-go-deps:
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod tidy

ifneq (,$(wildcard vendor))
	go mod vendor
endif

.PHONY: vendor
