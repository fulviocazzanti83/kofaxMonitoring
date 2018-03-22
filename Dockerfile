FROM golang:1.8-alpine
RUN apk update && apk update && apk add git
ADD . /go/src/oversinerco/erp/workers
RUN cd /go/src/oversinerco/erp/workers && go get && go install
ENTRYPOINT /go/bin/workers
