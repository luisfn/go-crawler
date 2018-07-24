FROM golang:1.10

WORKDIR /go/src/github.com/luisfn/crawler/

RUN go get -u github.com/golang/dep/cmd/dep

CMD ["dep", "ensure"]
CMD ["go", "run", "/go/src/github.com/luisfn/crawler/crawler.go"]
#CMD ["tail", "-f", "/dev/null"]