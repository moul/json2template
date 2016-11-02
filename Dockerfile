FROM golang:1.7.1
COPY . /go/src/github.com/moul/json2template
WORKDIR /go/src/github.com/moul/json2template
RUN go install .
ENTRYPOINT ["json2template"]
