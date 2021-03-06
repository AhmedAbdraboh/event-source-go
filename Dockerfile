FROM golang:1.12.1-stretch

WORKDIR /go/src/roosh-app
COPY . .

ENV GO111MODULE off

RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8080

CMD ./roosh-app