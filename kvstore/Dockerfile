FROM golang:1.22.5-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go test -cover -v ./...
RUN go build -o kvstore
RUN chmod +x /src/kvstore


FROM alpine:3.16

WORKDIR /app

COPY --from=builder /src/kvstore /app/
EXPOSE 8010

CMD ["/app/kvstore"]
