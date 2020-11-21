FROM golang:1.15 AS builder

WORKDIR /app
COPY . .
RUN make build

FROM alpine:3.7
COPY --from=builder /app/nodelabd .
ENV TZ Asia/Seoul