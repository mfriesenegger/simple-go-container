FROM registry.suse.com/bci/golang:1.17 as builder
COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/bin/sample /go/src/server.go

FROM registry.suse.com/bci/bci-micro:latest
WORKDIR /go/bin
COPY --from=builder /go/bin/sample ./
EXPOSE 8080
CMD ["/go/bin/sample"]
