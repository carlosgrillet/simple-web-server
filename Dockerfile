# Builder
FROM golang:1.24.5 AS builder
WORKDIR /opt
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Runner
FROM golang:1.24.5
WORKDIR /opt
COPY --from=builder /opt/server .
CMD ["/opt/server"]

