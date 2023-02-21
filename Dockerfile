FROM golang:alpine AS builder
WORKDIR /build
COPY app .
RUN go build -o main main.go
FROM alpine
WORKDIR /build
COPY --from=builder /build/main /build/main
CMD ["./main"]