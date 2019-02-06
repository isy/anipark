FROM golang:1.10.2-alpine3.7

COPY src/be/ /go/src/be/

WORKDIR /go/src/be/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

CMD gin -i run