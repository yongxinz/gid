FROM golang:1.16 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /root/gid

COPY go.mod .
RUN go mod download

COPY . /root/gid

RUN go build -o main main.go

# EXPOSE 8080 50051

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从 builder 镜像中把可执行文件拷贝到当前目录
COPY --from=builder /root/gid/main /root/gid/conf/app.toml /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
