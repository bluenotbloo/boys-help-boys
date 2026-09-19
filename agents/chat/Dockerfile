# chat agent Dockerfile
FROM golang:1.22-alpine AS builder

LABEL maintainer="Alex"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY go.work ./
COPY common/go.mod ./common/
COPY agents/chat/go.mod ./agents/chat/

WORKDIR /app/agents/chat
RUN go mod download

COPY common/ ./../../common/
COPY agents/chat/ ./

RUN CGO_ENABLED=0 go build -o /chat ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /chat /chat
COPY agents/chat/configs/ /configs/

ENV CONFIG_PATH=/configs/config.yaml

EXPOSE 8083 9100

ENTRYPOINT ["/chat"]