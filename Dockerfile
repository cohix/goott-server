FROM golang:1.11.2 as builder

RUN mkdir -p /go/src/github.com/cohix/goott-server
COPY . /go/src/github.com/cohix/goott-server

WORKDIR /go/src/github.com/cohix/goott-server
RUN go build

FROM cohix/goott-server-base:latest

RUN mkdir /goott
COPY --from=builder /go/src/github.com/cohix/goott-server/goott-server /goott/server
RUN chmod +x /goott/server

EXPOSE 3687

CMD /goott/server