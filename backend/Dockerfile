FROM golang:1.19.3 AS builder

WORKDIR /usr/src/app
ADD . /usr/src/app

RUN go build -o main

FROM alpine:3.17.3

WORKDIR /src
COPY --from=builder /usr/src/app/main /src/main

RUN apk add --no-cache ca-certificates && update-ca-certificates
ENV SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
ENV SSL_CERT_DIR=/etc/ssl/certs

RUN apk add --no-cache libc6-compat

EXPOSE 8000
CMD ["./main"]