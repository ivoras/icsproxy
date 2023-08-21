FROM golang:1.21-alpine

COPY *.go go.* /srv/icsproxy
WORKDIR /srv/icsproxy
RUN go build

CMD /srv/icsproxy/icsproxy
