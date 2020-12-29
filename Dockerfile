FROM golang:1.15.3-buster AS builder

WORKDIR /app

# Download libraries
COPY go.mod .
COPY go.sum .

RUN go mod download

# Build executable
COPY cmd cmd
COPY internal internal

RUN CGO_ENABLED=0 GOOS=linux go build github.com/zerodays/sistem-payments/cmd/payments

FROM debian:buster

# Get certificates.
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates
RUN update-ca-certificates

# Copy executable
WORKDIR /app
COPY --from=builder /app/payments .

# Entry point
ENTRYPOINT ["./payments"]
CMD ["serve"]
