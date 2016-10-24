FROM jenkins:2.19.1
MAINTAINER jang.jaehue@eland.co.kr

USER root

# Add the s6 overlay.
ENV S6_VERSION v1.18.1.5
RUN curl -L "https://github.com/just-containers/s6-overlay/releases/download/$S6_VERSION/s6-overlay-amd64.tar.gz" | \
    tar xzvf - -C /
ENV S6_BEHAVIOUR_IF_STAGE2_FAILS 1

# Download confd.
ENV CONFD_VERSION 0.11.0
RUN curl -L "https://github.com/kelseyhightower/confd/releases/download/v$CONFD_VERSION/confd-$CONFD_VERSION-linux-amd64" > /usr/bin/confd && \
    chmod +x /usr/bin/confd
ENV CONFD_OPTS '--backend=env --onetime'

# install docker
RUN \
  curl -sSL https://get.docker.com/ | sh
RUN curl -L https://github.com/docker/compose/releases/download/1.8.1/docker-compose-`uname -s`-`uname -m` > /usr/bin/docker-compose && chmod +x /usr/bin/docker-compose

# install etc
RUN apt-get -y install docker-engine libsystemd-journal0 jq

COPY env/plugins.txt  /usr/share/jenkins/plugins.txt

RUN /usr/local/bin/plugins.sh /usr/share/jenkins/plugins.txt

# Run the s6-based init.
ENTRYPOINT ["/init"]

# Set up a standard volume for logs.
VOLUME ["/var/log/services"]

# Start Jenkins by default
CMD [ "/usr/local/bin/jenkins.sh" ]
