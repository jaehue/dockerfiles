FROM jenkins:2.19.4
MAINTAINER jang.jaehue@eland.co.kr

USER root

# install docker
RUN \
  curl -sSL https://get.docker.com/ | sh
RUN curl -L https://github.com/docker/compose/releases/download/1.8.1/docker-compose-`uname -s`-`uname -m` > /usr/bin/docker-compose && chmod +x /usr/bin/docker-compose

# install etc
RUN apt-get update && apt-get install -y libsystemd-journal0 jq

USER jenkins

COPY env/plugins.txt /usr/share/jenkins/ref/
RUN /usr/local/bin/plugins.sh /usr/share/jenkins/ref/plugins.txt
