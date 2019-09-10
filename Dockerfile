FROM golang:1.12

RUN mkdir -p $GOPATH/src/three
WORKDIR $GOPATH/src/three

COPY . .
RUN cd $GOPATH/src/three/ && go get

EXPOSE 7779

RUN go get
RUN go build

CMD ["./three"]
