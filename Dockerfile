FROM alpine:3.10.1
ADD go-test-server /opt/go-test-server
CMD /opt/go-test-server --help