FROM golang:1.13-alpine as builder
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/brianbroderick/wergild
COPY . . 
RUN go get -d -v ./...
RUN go install -v /go/src/github.com/brianbroderick/wergild/cmd/wergild/ 

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/bin/wergild .

EXPOSE 2222
CMD ["./wergild"]
