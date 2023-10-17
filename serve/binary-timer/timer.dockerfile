# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest AS build

# 创建一个工作目录
WORKDIR /work

# 将本地代码拷贝到容器中
COPY . .

# 编译应用程序
RUN GOOS=linux GOARCH=amd64 go build -o timer serve/binary-timer/main.go && \
    mv timer binary/timer


# 运行阶段 使用简洁的作为基础镜像
FROM --platform=linux/amd64 alpine:latest

# 设置工作目录
WORKDIR /home/yezhiming/binary/timer

# 将打包后的程序拷贝到容器中
COPY --from=build /work/binary/timer /home/yezhiming/binary/timer

# 创建配置文件挂载目录
RUN mkdir -p /home/yezhiming/binary/timer/config

# 运行应用程序
CMD ["./timer", ">/dev/null", "2>error.log", "&"]
