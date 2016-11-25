FROM alpine:3.4
MAINTAINER JangJaehue <jaehue@jang.io>

ENV SSH_AUTH_SOCK /ssh-agent

# Install the SSH-client
RUN apk add --update openssh-client && rm -rf /var/cache/apk/*

VOLUME ["/ssh-agent"]
EXPOSE 2222 
ENTRYPOINT ["/usr/bin/ssh", "-T", "-N", "-o", "StrictHostKeyChecking=false", "-o", "ServerAliveInterval=3600", "-L"]
