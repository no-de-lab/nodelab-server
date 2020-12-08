FROM golang:1.15 AS builder

WORKDIR /app
COPY . .
RUN make build

FROM alpine:3.7
WORKDIR /app

RUN apk add --no-cache curl
COPY --from=builder /app/nodelabd .
COPY --from=builder /app/config.toml .
ENV TZ Asia/Seoul
