# Onur is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
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

.PHONY: imports grab vet test lint install deps coverage
.DEFAULT_GOAL := test

RUNNER ?= podman

OS :=linux
ARCH := amd64

NAME := onur
MAIN := ./cmd/${NAME}/main.go
DEST := ${HOME}/.local/bin

IMAGE_BUILD := ${USER}/${NAME}:$(shell cat .version)

# ================================= CONTAINER

image-build:
	@${RUNNER} build --file ./Containerfile --tag ${IMAGE_BUILD}

image-repl:
	@${RUNNER} run --rm -it -v ${PWD}:/app -w /app golang:1 bash


# ================================= UTILS

deps:
	@go mod download

imports:
	@goimports -l -w .

coverage:
	@go test --cover ./... -coverprofile=coverage.out

build: test
	@GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${NAME} ${MAIN}

install: build
	@mv -v ./${NAME} ${DEST}/${NAME}

lint:
	@golangci-lint run --enable-all internal cmd/pak

clean:
	@go clean
	@rm ${NAME}
