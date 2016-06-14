FROM gliderlabs/consul:0.6
MAINTAINER Jaehue Jang <jaehue@jang.io>

ADD https://releases.hashicorp.com/consul/0.6.4/consul_0.6.4_web_ui.zip /tmp/webui.zip
RUN mkdir /ui && cd /ui && unzip /tmp/webui.zip && rm /tmp/webui.zip && mv dist/* . && rm -rf dist

EXPOSE 8300 8301 8301/udp 8302 8302/udp 8400 8500 8600 8600/udp

ENV DNS_RESOLVES consul
ENV DNS_PORT 8600

VOLUME /config

ENTRYPOINT ["/bin/consul", "agent", "-config-dir=/config"]