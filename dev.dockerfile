# 使用官方的 Golang 镜像作为基础镜像
FROM golang:latest AS build

# 创建一个工作目录
WORKDIR /work

# 将本地代码拷贝到容器中
COPY . .

# 编译应用程序
RUN GOOS=linux GOARCH=amd64 go build -o fund-stock-api serve/binary-api/main.go && \
    mv fund-stock-api binary/fund-stock-api


# 运行阶段 使用简洁的作为基础镜像
FROM --platform=linux/amd64 alpine:latest

# 设置工作目录
WORKDIR /home/yezhiming/binary/api

# 将打包后的程序拷贝到容器中
COPY --from=build /work/binary/api /home/yezhiming/binary/api

# 创建配置文件挂载目录
RUN mkdir -p /home/yezhiming/binary/api/config

# 暴露应用程序使用的端口
EXPOSE 7104

# 运行应用程序
CMD ["./fund-stock-api", "--APP_PORT=7104", ">/dev/null", "2>error.log", "&"]
