FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o statistics-sender ./cmd/main.go

FROM alpine:latest

RUN apk update

COPY --from=builder /app/statistics-sender /usr/bin

ENV PATH="/app/usr/bin:${PATH}"

CMD ["statistics-sender"]
