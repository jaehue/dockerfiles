FROM golang:1.13
MAINTAINER jeff@blur.consulting

ENV GO111MODULE=on
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

# install go package
RUN go get github.com/Pallinder/go-randomdata && \
    go get github.com/dgrijalva/jwt-go && \
    go get github.com/gin-gonic/gin && \
    go get github.com/howeyc/fsnotify && \
    go get github.com/jinzhu/gorm && \
    go get github.com/lestrrat-go/file-rotatelogs && \
    go get github.com/lestrrat-go/strftime && \
    go get github.com/pilu/config && \
    go get github.com/pilu/fresh && \
    go get github.com/pkg/errors && \
    go get github.com/rifflock/lfshook && \
    go get github.com/sirupsen/logrus && \
    go get github.com/spf13/viper && \
    go get github.com/stretchr/testify && \
    go get golang.org/x/tools/gopls
