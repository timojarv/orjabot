FROM golang:1

WORKDIR /go/src/orjabot

ADD . .

RUN go get -v
RUN go build -o orja .

CMD ["./orja"]
