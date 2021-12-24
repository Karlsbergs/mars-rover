FROM golang:alpine as MarsRover

COPY . /src

RUN cd /src && go build -o /go/MarsRover main.go

ENTRYPOINT ["/go/MarsRover"]