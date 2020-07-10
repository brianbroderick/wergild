FROM golang:1.13-alpine as builder
RUN apk update && apk upgrade && \
  apk add --no-cache git openssh

WORKDIR /go/src/github.com/brianbroderick/wergild
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . . 

RUN go install -v /go/src/github.com/brianbroderick/wergild/cmd/wergild/ 

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
RUN mkdir welcome
COPY --from=builder /go/src/github.com/brianbroderick/wergild/cmd/wergild/welcome/ welcome
COPY --from=builder /go/bin/wergild .

EXPOSE 2222
# CMD ["./wergild"]
