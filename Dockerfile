FROM golang:alpine as builder

WORKDIR /workspace
COPY . .
# swagger 文档
RUN go install github.com/swaggo/swag/cmd/swag@latest \
    && swag init
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="123456@qq.com"

WORKDIR /workspace

COPY --from=0 /workspace/server ./
COPY --from=0 /workspace/resource ./resource/
COPY --from=0 /workspace/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
