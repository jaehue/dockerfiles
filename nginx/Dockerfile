FROM ubuntu:16.04

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    curl perl make build-essential procps \
    libreadline-dev libncurses5-dev libpcre3-dev libssl-dev \
    git wget unzip \
 && rm -rf /var/lib/apt/lists/*

ENV OPENRESTY_VERSION 1.9.15.1
ENV OPENRESTY_PREFIX /opt/openresty
ENV NGINX_PREFIX /opt/openresty/nginx
ENV VAR_PREFIX /var/nginx

# NginX prefix is automatically set by OpenResty to $OPENRESTY_PREFIX/nginx
# look for $ngx_prefix in https://github.com/openresty/ngx_openresty/blob/master/util/configure

RUN cd /root \
 && echo "==> Downloading OpenResty..." \
 && curl -sSL http://openresty.org/download/ngx_openresty-${OPENRESTY_VERSION}.tar.gz | tar -xvz \
 && echo "==> Configuring OpenResty..." \
 && cd /root/openresty-${OPENRESTY_VERSION} \
 && readonly NPROC=$(grep -c ^processor /proc/cpuinfo 2>/dev/null || 1) \
 && echo "using upto $NPROC threads" \
 && ./configure \
    --prefix=$OPENRESTY_PREFIX \
    --http-client-body-temp-path=$VAR_PREFIX/client_body_temp \
    --http-proxy-temp-path=$VAR_PREFIX/proxy_temp \
    --http-log-path=$VAR_PREFIX/access.log \
    --error-log-path=$VAR_PREFIX/error.log \
    --pid-path=$VAR_PREFIX/nginx.pid \
    --lock-path=$VAR_PREFIX/nginx.lock \
    --with-luajit \
    --with-pcre-jit \
    --with-ipv6 \
    --with-http_ssl_module \
    --with-http_stub_status_module \
    --without-http_ssi_module \
    --without-http_userid_module \
    --without-http_fastcgi_module \
    --without-http_uwsgi_module \
    --without-http_scgi_module \
    --without-http_memcached_module \
    -j${NPROC} \
 && echo "==> Building OpenResty..." \
 && make -j${NPROC} \
 && echo "==> Installing OpenResty..." \
 && make install \
 && echo "==> Finishing..." \
 && ln -sf $NGINX_PREFIX/sbin/nginx /usr/local/bin/nginx \
 && ln -sf $NGINX_PREFIX/sbin/nginx /usr/local/bin/openresty \
 && ln -sf $OPENRESTY_PREFIX/bin/resty /usr/local/bin/resty \
 && ln -sf $OPENRESTY_PREFIX/luajit/bin/luajit-* $OPENRESTY_PREFIX/luajit/bin/lua \
 && ln -sf $OPENRESTY_PREFIX/luajit/bin/luajit-* /usr/local/bin/lua \
 && rm -rf /root/ngx_openresty*

# install consul-template
RUN cd /tmp && \
    wget https://releases.hashicorp.com/consul-template/0.15.0/consul-template_0.15.0_linux_amd64.zip --no-check-certificate && \
    unzip consul-template_0.15.0_linux_amd64.zip && \
    chmod a+x consul-template && \
    mv consul-template /usr/bin/consul-template && \
    rm consul-template_0.15.0_linux_amd64.zip

# install nginx-jwt
RUN cd /tmp && mkdir -p /opt/openresty/nginx/lua && \
    wget https://github.com/auth0/nginx-jwt/releases/download/v1.0.1/nginx-jwt.tar.gz --no-check-certificate && \
    tar -C /opt/openresty/nginx/lua -xf nginx-jwt.tar.gz && \
    rm nginx-jwt.tar.gz
ENV JWT_SECRET="JWTs are the best!"

## install golang 1.6.2
ENV GOLANG_VERSION 1.6.2
RUN cd /tmp && \
    wget https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz --no-check-certificate && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $NGINX_PREFIX/

ONBUILD RUN rm -rf conf/* html/*
#ONBUILD COPY nginx $NGINX_PREFIX/

EXPOSE 80

CMD ["nginx", "-g", "daemon off; error_log /dev/stderr info;"]
