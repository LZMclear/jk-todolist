# Multi-stage Dockerfile (simple, builds from local files)
FROM golang:1.23.5-alpine AS builder
# 设置国内代理
ENV GOPROXY=https://goproxy.cn,direct
RUN apk add --no-cache ca-certificates
WORKDIR /src

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/todo ./cmd/todo

# 运行时镜像
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=builder /app/todo /usr/local/bin/todo
COPY --from=builder /src/web /app/web

ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/todo"]
