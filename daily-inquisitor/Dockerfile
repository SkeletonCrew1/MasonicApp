FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o inquisitor-server ./cmd/inquisitor

# COPY entrypoint.sh ./
# RUN chmod +x entrypoint.sh

CMD ["./inquisitor-server"]