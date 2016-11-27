FROM java:openjdk-7u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
USER root
ENV ZEPPELIN_USER zeppelin
ENV ZEPPELIN_HOME /usr/local/zeppelin-bin-all
ENV ZEPPELIN_LOG_DIR /log
ENV ZEPPELIN_PID_DIR /run
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV ZEPPELIN_VERSION 0.6.2

RUN wget -q -O - http://apache.tt.co.kr/zeppelin/zeppelin-$ZEPPELIN_VERSION/zeppelin-$ZEPPELIN_VERSION-bin-all.tgz | tar -xzf - -C /usr/local \
 && ln -s zeppelin-$ZEPPELIN_VERSION-bin-all /usr/local/zeppelin-bin-all \
 && groupadd -r $ZEPPELIN_USER && useradd -r -g $ZEPPELIN_USER -G zeppelin $ZEPPELIN_USER

RUN chown -R $ZEPPELIN_USER:zeppelin /usr/local/zeppelin-bin-all \
 && chmod -R 775 /usr/local/zeppelin-bin-all/bin

USER $ZEPPELIN_USER