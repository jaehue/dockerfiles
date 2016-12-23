FROM java:openjdk-8u111-jdk
MAINTAINER Jaehue Jang <jaehue@jang.io>

# install mvn
ARG MAVEN_VERSION=3.3.9
ARG USER_HOME_DIR="/root"

RUN mkdir -p /usr/share/maven /usr/share/maven/ref \
  && curl -fsSL http://apache.osuosl.org/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz \
    | tar -xzC /usr/share/maven --strip-components=1 \
  && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn

# install npm
RUN apt-get update && apt-get -qq -y install build-essential nodejs
RUN apt-get -qq -y install npm

# build nifi
ENV NIFI_VERSION 1.1.1
ENV NIFI_HOME /usr/local/nifi-1.1.1

RUN wget -q -O - https://github.com/apache/nifi/archive/rel/nifi-$NIFI_VERSION.tar.gz | tar -xzf - -C /tmp \
  && cd /tmp/nifi-rel-nifi-$NIFI_VERSION \
  && mvn clean install -Dmaven.test.skip=true \
  && tar xzf /tmp/nifi-rel-nifi-$NIFI_VERSION/nifi-assembly/target/nifi-1.1.1-bin.tar.gz -C /usr/local \
  && rm -rf /tmp/nifi-rel-nifi-$NIFI_VERSION

RUN groupadd -r nifi && useradd -r -g nifi -G nifi nifi \
 && chown nifi:nifi -R $NIFI_HOME

# remove npm
RUN apt-get -qq -y remove build-essential nodejs npm

USER nifi

EXPOSE 8080

CMD ["/usr/local/nifi-1.1.1/bin/nifi.sh", "run"]
