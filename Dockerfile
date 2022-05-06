FROM registry.suse.com/bci/golang:1.17

RUN echo "V0.01"

COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/bin/sample /go/src/server.go

EXPOSE 8080
CMD ["/go/bin/sample"]
