# 基础镜像
FROM alpine
RUN echo ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir -p /opt/myddns
ADD  myddns /opt/myddns/myddns
ADD  config.yaml /opt/myddns/config.yaml
# 解决控制台乱码 start
ENV LANG C.UTF-8
ENV LANGUAGE zh_CN.UTF-8
ENV LC_ALL C.UTF-8
ENV TZ Asia/Shanghai
WORKDIR /opt/myddns
RUN chmod -R 755 /opt/myddns
CMD ["./myddns"]