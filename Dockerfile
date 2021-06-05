FROM scratch

WORKDIR $GOPATH/src/github.com/qiaocco/go-gin-example
COPY . $GOPATH/src/github.com/qiaocco/go-gin-example

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]