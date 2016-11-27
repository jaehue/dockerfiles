FROM java:openjdk-7u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
USER root
ENV ZEPPELIN_VERSION 0.6.2

RUN wget -q -O - http://apache.tt.co.kr/zeppelin/zeppelin-$ZEPPELIN_VERSION/zeppelin-$ZEPPELIN_VERSION-bin-all.tgz | tar -xzf - -C /usr/local \
 && ln -s zeppelin-$ZEPPELIN_VERSION-bin-all /usr/local/zeppelin-bin-all

RUN groupadd -r zeppelin && useradd -r -g zeppelin -G zeppelin zeppelin \
 && chown -R zeppelin:zeppelin /usr/local/zeppelin-bin-all \
 && chmod -R 775 /usr/local/zeppelin-bin-all/bin

USER zeppelin