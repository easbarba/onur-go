# Onur is free software: you can redistribute it and/or modify
# it under  the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Onur is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Onur. If not, see <https://www.gnu.org/licenses/>.

# DEPS: fzf podman bash:5.5

.DEFAULT_GOAL := test

OS :=linux
ARCH := amd64

NAME := onur
DEST := ${HOME}/.local/bin

RUNNER ?= podman
VERSION := $(shell cat .version)
CONTAINER_IMAGE := registry.gitlab.com/${USER}/${NAME}-go:${VERSION}

# ================================= MANAGEMENT

.PHONY: local.deps
local.deps:
	@go mod download

.PHONY: local.imports
local.imports:
	@goimports -l -w .

.PHONY: local.coverage
local.coverage:
	@go test --cover ./... -coverprofile=coverage.out

.PHONY: local.lint
local.lint:
	@golangci-lint run --enable-all internal cmd/pak

.PHONY: local.build
local.build:
	GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${NAME} ./cmd/cli/*

.PHONY: local.install
local.install: local.build
	@mv -v ./${NAME} ${DEST}/${NAME}

.PHONY: local.clean
local.clean:
	@go clean
	@rm ${NAME}

.PHONY: local.test
local.test:
	go test -v ./internal/... ./cmd/...

# ================================= CONTAINER

.PHONY: image.build.dev
image.build.dev:
	${RUNNER} build --file ./Containerfile-dev --tag ${CONTAINER_IMAGE} --env ONUR_VERSION=${VERSION}

.PHONY: image.build.prod
image.build.prod:
	${RUNNER} build --file ./Containerfile-prod --tag ${CONTAINER_IMAGE} --env ONUR_VERSION=${VERSION}

.PHONY: image.repl
image.repl:
	${RUNNER} run --rm -it \
		--volume ${PWD}:/app:Z \
		--workdir /home/easbarba/app \
		${CONTAINER_IMAGE} bash

.PHONY: image.publish
image.publish:
	${RUNNER} push ${CONTAINER_IMAGE}

.PHONY: image.commands
image.commands:
	${RUNNER} run --rm -it \
		--volume ${PWD}:/app:Z \
		--workdir /home/easbarba/app \
		${CONTAINER_IMAGE} bash -c "$(shell cat ./container-commands | fzf)"
