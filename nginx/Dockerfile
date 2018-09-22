#
# nginx (http://nginx.org)
#
# build:
#   docker build --force-rm=true -t jaehue/nginx:1.14.0 .
# run:
#   docker run -p 80:80 -v /data/nginx/sites:/usr/local/nginx/conf/sites -d jaehue/nginx:1.14.0
# reload
#   docker exec [container id] /usr/local/sbin/nginx -s reload
#

FROM nginx:latest

MAINTAINER jang.jaehue@eland.co.kr

RUN apt-get update && apt-get -y upgrade
RUN apt-get -qq -y install build-essential zlib1g-dev libpcre3-dev unzip wget uuid-dev curl
ADD ./pagespeed_install.sh /tmp/
RUN chmod 755 /tmp/pagespeed_install.sh && /tmp/pagespeed_install.sh --nginx-version latest
# RUN chmod 755 /tmp/pagespped_install.sh
# RUN curl -f -L -sS https://ngxpagespeed.com/install > pagespped_install.sh
# RUN chmod 755 pagespped_install.sh
# RUN ./pagespped_install.sh --nginx-version latest