#FROM golang:1.22 AS builder

#WORKDIR /app
#
#COPY . .
#
#RUN go mod tidy
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o telegram-bot ./cmd/main.go

FROM alpine:latest
#COPY --from=builder /app/telegram-bot /usr/bin

COPY telegram-bot /usr/bin

COPY .env .
RUN mkdir config
COPY /config/dev_config.yaml /config

CMD ["./usr/bin/telegram-bot"]