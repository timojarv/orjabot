FROM golang:1 AS builder

WORKDIR /go/src/github.com/timojarv/orjabot

ADD . .

RUN go get
RUN go build -o orja .

FROM golang:1

COPY --from=builder /go/src/github.com/timojarv/orjabot/ /

CMD ["/orja"]
