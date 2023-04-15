FROM golang:1.20 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOOS=linux GOARCH=amd64
RUN go build -o ./qas ./cmd/qas/main.go

FROM golang:1.20-alpine
MAINTAINER EAS Barbosa <easbarba@outlook.com>
COPY --from=build /app/qas /opt/qas
COPY examples /root/.config/qas
CMD [ "/opt/qas" ]
