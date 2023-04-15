# Qas is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Qas is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Qas. If not, see <https://www.gnu.org/licenses/>.

.DEFAULT_GOAL := build

OS :=linux
ARCH := amd64

BINARY_NAME := qas
MAIN := ./main.go
TO := ${HOME}/.local/bin

all: build test

deps:
	go mod download

build: test
	GOARCH=$(ARCH) GOOS=$(OS) go build -race -ldflags "-extldflags '-static'" -o ${BINARY_NAME} ${MAIN}

install: build
	mv -v ./${BINARY_NAME} ${TO}/${BINARY_NAME}

lint:
	golangci-lint run --enable-all internal cmd/pak

test:
	go test -v ./...

clean:
	go clean
	rm ${BINARY_NAME}

vet:
	go vet ./...

grab:
	go run ./main.go --grab

archive:
	go run ./main.go --archive meh,forevis,tar

imports:
	goimports -l -w .

coverage:
	go test --cover ./... -coverprofile=coverage.out

image:
	podman build --file ./Dockerfile --tag $USER/${BINARY_NAME}:$(shell cat .env)

.PHONY: imports grab vet test lint install deps coverage
