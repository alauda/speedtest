FROM golang:1.8

ADD . $GOPATH/src/app

RUN cd $GOPATH/src/app && go install

CMD ["app"]

