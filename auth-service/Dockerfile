FROM alpine:3.23.5 AS build
ARG GOLANG_VERSION=1.26.0


RUN apk update && apk add go gcc bash musl-dev openssl-dev ca-certificates && update-ca-certificates

RUN wget https://dl.google.com/go/go$GOLANG_VERSION.src.tar.gz && tar -C /usr/local -xzf go$GOLANG_VERSION.src.tar.gz

RUN cd /usr/local/go/src && ./make.bash

ENV PATH=$PATH:/usr/local/go/bin

RUN rm go$GOLANG_VERSION.src.tar.gz


RUN apk del go

RUN go version
WORKDIR /auth
COPY go.mod go.sum main.go utils.go models.go ./
RUN go build -o main

FROM alpine:3.23.5
WORKDIR /auth
COPY --from=build /auth/main ./
EXPOSE 8080
CMD ["./main"]