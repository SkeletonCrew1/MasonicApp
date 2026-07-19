FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go mod tidy && go build -o ban-service .
EXPOSE 8081
CMD ["./ban-service"]