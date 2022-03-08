# 采源代码编译形式
#FROM golang:alpine AS build-env
#MAINTAINER yzc "ylzlcl@163.com"
#
#ENV GO111MODULE=on \
#    GOPROXY=https://goproxy.cn,direct \
#    GIN_MODE=release
#
#WORKDIR /upload_task
## golang 静态文件不打包进去
##ADD config /controller_minio/config
#
#ADD . /upload_task
#
#RUN go build -o task ./main.go
#
#
#EXPOSE 9000
#
#CMD /upload_task/task



## linux 直接采用二进制文件运行
FROM  ubuntu:latest

## 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
## 设置编码
ENV LANG C.UTF-8
## 添加当前目录二进制文件进容器当中
ADD ./task ./task
CMD ["./task"]

