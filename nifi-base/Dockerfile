FROM maven:3.3.9-jdk-8
MAINTAINER Jaehue Jang <jaehue@jang.io>

RUN apt-get update && apt-get -qq -y install build-essential nodejs
RUN apt-get -qq -y install npm

ENV MAVEN_OPTS -Xms1024m -Xmx3076m -XX:MaxPermSize=256m
ENV NIFI_VERSION 1.1.0-RC1

RUN wget -q -O - https://github.com/apache/nifi/archive/nifi-$NIFI_VERSION.tar.gz | tar -xzf - -C /usr/local \
  && cd /usr/local/nifi-nifi-$NIFI_VERSION \
  && mvn -T C2.0 clean install
