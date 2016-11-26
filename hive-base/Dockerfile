FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# Setup env
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
ENV HIVE_VERSION 2.1.0

RUN wget -q -O - http://mirror.navercorp.com/apache/hive/hive-$HIVE_VERSION/apache-hive-$HIVE_VERSION-bin.tar.gz | tar -xzf - -C /usr/local \
 && ln -s /usr/local/apache-hive-$HIVE_VERSION-bin /usr/local/apache-hive-bin