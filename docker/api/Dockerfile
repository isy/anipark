FROM golang:1.11.5-alpine3.8

COPY src/be/ /go/src/github.com/isy/ownedia/src/be/

WORKDIR /go/src/github.com/isy/ownedia/src/be/

RUN apk update \
  && apk add --no-cache git gcc libc-dev \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && go get -u bitbucket.org/liamstask/goose/cmd/goose \
  && dep ensure

CMD gin -i run