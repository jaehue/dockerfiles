FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV ZEPPELIN_VERSION 0.6.2

RUN wget -q -O - http://apache.tt.co.kr/zeppelin/zeppelin-$ZEPPELIN_VERSION/zeppelin-$ZEPPELIN_VERSION-bin-all.tgz | tar -xzf - -C /usr/local \
 && ln -s zeppelin-$ZEPPELIN_VERSION-bin-all /usr/local/zeppelin-bin-all
