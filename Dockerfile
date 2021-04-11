FROM registry.suse.com/suse/sle15:latest

RUN echo "V0.01"
RUN mv /etc/zypp /etc/zypp.pristine

RUN zypper -n rm -y container-suseconnect
COPY zypp /etc/zypp
RUN zypper -n in -y go1.16

COPY src /go/src
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/bin/sample /go/src/server.go

RUN rm -rf /etc/zypp && mv /etc/zypp.pristine /etc/zypp

EXPOSE 8080
CMD ["/go/bin/sample"]
