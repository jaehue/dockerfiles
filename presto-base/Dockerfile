FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env

ENV PRESTO_USER presto
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV PRESTO_VERSION 0.157

USER root

RUN wget -q -O - https://repo1.maven.org/maven2/com/facebook/presto/presto-server/$PRESTO_VERSION/presto-server-$PRESTO_VERSION.tar.gz | tar -xzf - -C /usr/local \
 && ln -s presto-server-$PRESTO_VERSION /usr/local/presto-server

WORKDIR /usr/local/presto-server

RUN wget -q https://repo1.maven.org/maven2/com/facebook/presto/presto-cli/$PRESTO_VERSION/presto-cli-$PRESTO_VERSION-executable.jar \
  && mv presto-cli-$PRESTO_VERSION-executable.jar /usr/local/presto-server/bin/presto-cli \
  && chmod +x /usr/local/presto-server/bin/presto-cli \
  && chown -R $PRESTO_USER:presto /usr/local/presto-server \
  && chmod -R 775 /usr/local/presto-server/bin
  