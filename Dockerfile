FROM golang:1.15 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o nodelabd ./cmd/...

FROM alpine:3.7
COPY --from=builder /app/nodelabd .
ENV TZ Asia/Seoul