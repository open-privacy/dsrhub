FROM golang:1.14
WORKDIR /go/src/github.com/dsrhub/dsrhub/mocks
ADD . .
RUN go build
EXPOSE 9999
CMD ./mocks
