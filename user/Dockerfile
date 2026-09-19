# user service Dockerfile
FROM golang:1.22-alpine AS builder

LABEL maintainer="Alex"

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

# Copy go.work and module files for dependency caching
COPY go.work ./
COPY common/go.mod ./common/
COPY user/go.mod ./user/

# Download dependencies
WORKDIR /app/user
RUN go mod download

# Copy source code
COPY common/ ./../common/
COPY user/ ./

# Build
RUN CGO_ENABLED=0 go build -o /user ./cmd/main.go

# Runtime stage
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /user /user
COPY user/configs/ /configs/

ENV CONFIG_PATH=/configs/config.yaml

EXPOSE 8082 9000

ENTRYPOINT ["/user"]