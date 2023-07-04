FROM golang:1.20.5

WORKDIR /src

# get json data
RUN curl -LO http://api.nobelprize.org/v1/prize.json

COPY go.mod go.sum .
COPY v1 v1/
RUN CGO_ENABLED=0 go build -o prize-v1 ./v1

COPY v2 v2/
RUN CGO_ENABLED=0 go build -o prize-v2 ./v2

COPY v3 v3/
RUN CGO_ENABLED=0 go build -o prize-v3 ./v3

COPY v4 v4/
RUN CGO_ENABLED=0 go build -o prize-v4 ./v4
