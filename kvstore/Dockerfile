FROM golang:1.22.5 AS build
ADD . /src
WORKDIR /src
RUN go mod tidy
RUN go test --cover -v ./...
RUN go build -v -o /kvstore


FROM alpine:3.4
EXPOSE 8010
CMD ["kvstore"]
COPY --from=build /src/kvstore /usr/local/bin/kvstore
RUN chmod +x /usr/local/bin/kvstore
