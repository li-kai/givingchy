FROM golang:alpine

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

COPY . /go/src/app
RUN go build -o bin *.go

FROM alpine:latest

WORKDIR /root/
COPY --from=0 /go/src/app/bin .
CMD ["./bin"]
