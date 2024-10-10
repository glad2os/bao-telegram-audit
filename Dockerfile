FROM ubuntu:latest
LABEL authors="glad2os"

WORKDIR /app

FROM golang:1.23.2
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bao-telegram-audit && \
    chmod +x ./bao-telegram-audit

# Добавить запуск от юзера: группа 1000
ENTRYPOINT ["./bao-telegram-audit"]