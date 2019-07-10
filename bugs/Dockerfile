FROM golang

RUN mkdir -p /go/src/github.com/najamsk/kitplay/first

ADD . /go/src/github.com/najamsk/kitplay/first
WORKDIR /go/src/github.com/najamsk/kitplay/first/bugs

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/najamsk/kitplay/first/bugs/cmd  -watch github.com/najamsk/kitplay/first/bugs
