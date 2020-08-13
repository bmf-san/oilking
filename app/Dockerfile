FROM golang:1.15.0-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash && \
    go get -u golang.org/x/lint/golint

WORKDIR /go/src/oilking

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app

FROM alpine

COPY --from=builder /go/src/oilking/app ./

CMD ["./app"]