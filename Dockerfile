FROM alpine:3.7

WORKDIR /www
COPY   ./bin/process-server /usr/bin/process-server

RUN chmod +x /usr/bin/process-server

ENTRYPOINT ["process-server"]
