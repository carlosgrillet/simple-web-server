# Builder
FROM golang:1.24.5 AS builder
WORKDIR /opt
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# Runner
FROM scratch
WORKDIR /opt
COPY --from=builder /opt/server .
CMD ["/opt/server"]

