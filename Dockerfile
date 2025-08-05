# Build stage
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

WORKDIR /app/cmd/server
RUN go build -o rpc-server

# Final image
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/cmd/server/rpc-server .

CMD ["./rpc-server"]