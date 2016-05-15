FROM golang:1.6
MAINTAINER jaehue@jang.io

# add application
WORKDIR /go/src/whoami
ADD . /go/src/whoami
RUN go install

EXPOSE 3000

CMD ["/go/bin/whoami"]
