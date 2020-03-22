
FROM alpine:3.7

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/assets

WORKDIR /www
COPY   ./bin/xprocess-server /usr/bin/xprocess-server
ADD   ./config.json /www/config.json
ADD ./assets /www/assets

RUN chmod +x /usr/bin/xprocess-server

ENTRYPOINT ["xprocess-server"]
