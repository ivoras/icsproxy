FROM golang:1.21-alpine

COPY . /srv/icsproxy
RUN rm .env
RUN go build

CMD /srv/icsproxy/icsproxy
