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
MAIN := ./cmd/${NAME}/main.go
DEST := ${HOME}/.local/bin

RUNNER ?= podman
VERSION := $(shell cat .version meson.build)
CONTAINER_IMAGE := registry.gitlab.com/${USER}/${NAME}:${VERSION}

# ================================= MANAGEMENT

.PHONY: deps
deps:
	@go mod download

.PHONY: imports
imports:
	@goimports -l -w .

.PHONY: imports
coverage:
	@go test --cover ./... -coverprofile=coverage.out

.PHONY: lint
lint:
	@golangci-lint run --enable-all internal cmd/pak

.PHONY: loca.build
local.build:
	GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${NAME} ${MAIN}

.PHONY: loca.install
local.install: local.build
	@mv -v ./${NAME} ${DEST}/${NAME}

.PHONY: loca.clean
local.clean:
	@go clean
	@rm ${NAME}

# ================================= CONTAINER

.PHONY: image.build
image.build:
	${RUNNER} build --file ./Containerfile --tag ${CONTAINER_IMAGE} --env ONUR_VERSION=${VERSION}

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


test:
	go test -v ./internal/... ./cmd/...
