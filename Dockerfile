FROM golang:1.7

ADD . /go/src/github.com/stupschwartz/gomin

WORKDIR /go/src/github.com/stupschwartz/gomin
RUN go get ./...
RUN go install github.com/stupschwartz/gomin

ENTRYPOINT /go/bin/gomin

EXPOSE 8080
