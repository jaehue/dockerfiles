FROM golang:latest

RUN go get github.com/Masterminds/glide && \
    git clone https://github.com/btcsuite/btcd $GOPATH/src/github.com/btcsuite/btcd && \
    cd $GOPATH/src/github.com/btcsuite/btcd && \
    glide install && \
    go install . ./cmd/...