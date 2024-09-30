FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o telegram-bot ./cmd/main.go

FROM alpine:latest

RUN apk update

COPY --from=builder /app/telegram-bot /usr/bin
#COPY telegram-bot /usr/bin

ENV PATH="/app/usr/bin:${PATH}"

CMD ["telegram-bot"]