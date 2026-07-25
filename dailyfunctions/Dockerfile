FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build both binaries from their respective cmd directories
RUN go build -o password-server ./cmd/password
RUN go build -o inquisitor-server ./cmd/inquisitor

# Copy the startup script
COPY entrypoint.sh ./
RUN chmod +x entrypoint.sh

# Run both services via the script
CMD ["./entrypoint.sh"]