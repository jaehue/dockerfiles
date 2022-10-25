FROM jenkins/jenkins:lts

USER root

# install docker
RUN apt-get update && \
    apt-get -y --no-install-recommends install apt-transport-https \
      ca-certificates \
      curl \
      gnupg2 \
      software-properties-common && \
    curl -sSL https://get.docker.com/ | sh && \
    chmod +x /usr/local/bin/docker-compose

USER jenkins
