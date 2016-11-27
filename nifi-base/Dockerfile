FROM maven:3.3.9-jdk-8
MAINTAINER Jaehue Jang <jaehue@jang.io>

RUN apt-get update && apt-get -qq -y install build-essential nodejs
RUN apt-get -qq -y install npm

ENV NIFI_VERSION 1.1.0-RC2

RUN wget -q -O - https://github.com/apache/nifi/archive/nifi-$NIFI_VERSION.tar.gz | tar -xzf - -C /tmp \
  && cd /tmp/nifi-nifi-$NIFI_VERSION \
  && mvn clean install -Dmaven.test.skip=true
# RUN tar xzf /tmp/nifi-nifi-NIFI_VERSION/nifi-assembly/target/nifi-*-bin.tar.gz -C /usr/local \
#   && rm -rf /tmp/nifi-nifi-$NIFI_VERSION

# RUN groupadd -r nifi && useradd -r -g nifi -G nifi nifi
#  && chown nifi:nifi -R /usr/local/nifi-*

# USER nifi