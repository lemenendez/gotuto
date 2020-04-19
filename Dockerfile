FROM golang:1.13.10-buster

WORKDIR /go/gotuto

RUN GOPATH=/go/gotuto

COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

CMD ["app"]