FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o ban-service . 

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ban-service .
EXPOSE 8081
CMD ["./ban-service"]
