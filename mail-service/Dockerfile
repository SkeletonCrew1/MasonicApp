FROM golang:1.26.5-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o mail-service .

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/mail-service .

ENTRYPOINT [ "./mail-service" ]