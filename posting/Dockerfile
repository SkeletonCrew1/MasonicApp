FROM ubuntu:24.04

WORKDIR /app

RUN apt-get update && \
    apt-get install -y golang-go ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

EXPOSE 8080

CMD ["./server"]