FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test -cover -v ./...
RUN go build -o kvstore


FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/kvstore .

EXPOSE 8010

CMD ["./kvstore"]
