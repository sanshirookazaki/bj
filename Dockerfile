FROM golang:1.12-stretch AS builder
LABEL maintainer "s.okazaki"

WORKDIR /go/src/app
COPY . .

RUN export GO111MODULE=on && \
    CGO_ENABLED=0 go build

FROM busybox
COPY --from=builder /go/src/app/bj /usr/local/bin/bj
ENTRYPOINT ["/usr/local/bin/bj"]