FROM phase2/jenkins-docker
MAINTAINER jang.jaehue@eland.co.kr

USER root

RUN apt-get update && \
    apt-get -y install libsystemd-journal0 jq

COPY env/plugins.txt  /usr/share/jenkins/plugins.txt

RUN /usr/local/bin/plugins.sh /usr/share/jenkins/plugins.txt
