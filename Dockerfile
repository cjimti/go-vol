FROM golang:latest AS builder

RUN mkdir -p /go/src/github.com/cjimti/go-vol
COPY . /go/src/github.com/cjimti/go-vol

RUN go get ...
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /go/bin/vol ./src/github.com/cjimti/go-vol

FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin/vol /vol

WORKDIR /

ENTRYPOINT ["/vol"]
