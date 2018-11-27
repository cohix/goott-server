FROM golang:1.12.2 as builder

RUN mkdir -p /go/src/github.com/cohix/goott-server
COPY . /go/src/github.com/cohix/goott-server

WORKDIR /go/src/github.com/cohix/goott-server
RUN go build

FROM cohix/goott-base:latest

RUN mkdir /goot
COPY --from=builder /go/src/github.com/cohix/goott-server /goot/server

CMD /goot/server