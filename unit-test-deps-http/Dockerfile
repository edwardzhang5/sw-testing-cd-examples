FROM golang:1.17
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -v github.com/rakyll/gotest
RUN go build -o main .
# CMD ["go", "test", "-v"]
CMD ["go", "test", "-v", "--tags=integration"]
# CMD ["/app/main"]

