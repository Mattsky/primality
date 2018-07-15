FROM golang:alpine3.8

MAINTAINER Matt North

COPY ./server.go /root/server.go

RUN cd ~ && go build

EXPOSE 8000

ENTRYPOINT ["/root/root"]
