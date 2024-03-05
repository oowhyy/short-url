FROM golang:1.21.7-bookworm AS builder

WORKDIR /shorturl

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/shorturl cmd/main.go

# base image
FROM alpine:latest
WORKDIR /
COPY --from=builder /bin/shorturl /bin/shorturl
COPY config.yaml /etc/shorturl/config.yaml

EXPOSE 3000
EXPOSE 50051

ENTRYPOINT ["./bin/shorturl", "--config"]
CMD ["/etc/shorturl/config.yaml"]

