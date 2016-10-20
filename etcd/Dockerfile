FROM        alpine:latest
RUN         apk add --update ca-certificates openssl tar && \
            wget https://github.com/coreos/etcd/releases/download/v3.0.12/etcd-v3.0.12-linux-amd64.tar.gz && \
            tar xzvf etcd-v3.0.12-linux-amd64.tar.gz && \
            mv etcd-v3.0.12-linux-amd64/etcd* /bin/ && \
            apk del --purge tar openssl && \
            rm -Rf etcd-v3.0.12-linux-amd64* /var/cache/apk/*
VOLUME      /data
EXPOSE      2379 2380 4001 7001
ENTRYPOINT ["/bin/etcd"]
