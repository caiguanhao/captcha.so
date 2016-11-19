FROM golang:1.7.1

WORKDIR /go/src/github.com/caiguanhao/captcha.so
ADD . .
RUN go build -buildmode=c-shared -v -ldflags="-s -w" -o /captcha.so
