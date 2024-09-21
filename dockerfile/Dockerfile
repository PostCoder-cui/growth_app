FROM alpine:3.16
LABEL authors="寒鸦少年"

RUN apk --update add bash
RUN apk add curl
RUN apk add vim

RUN mkdir -p /data/code

COPY ./client /data/code
COPY ./server /data/code
COPY gin_app /data/code
COPY ./grpcurl_linux /bin
RUN chmod +x /data/code/client /data/code/server /data/code/gin_app /bin/grpcurl_linux

WORKDIR /data/code

ENV USER_GROWTH_CONFIG='{"Db":{"Engine":"mysql","Username":"root","Password":"123321","Host":"192.168.137.151","Port":3306,"Database":"user_growth","Charset":"utf8","ShowSql":true,"MaxIdleConns":2,"MaxOpenConns":10,"ConnMaxLifetime":30},"Cache":{}}'

RUN mkdir -p /data/code/logs

COPY startup.sh /data/code
RUN chmod +x /data/code/startup.sh

EXPOSE 80
EXPOSE 8080
EXPOSE 3306

CMD ["/data/code/startup.sh"]