# 构建
FROM golang:1.18-alpine as builder
WORKDIR /home/project
ENV GOPROXY=https://goproxy.cn
COPY ./ ./
RUN go mod tidy && \
    go mod download && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add tzdata
RUN go build -o admin

# 打包
FROM alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /home/project/admin /home/project/
WORKDIR /home/project
CMD ["./admin"]