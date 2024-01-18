FROM golang:1.21-alpine AS builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
COPY main.go .

COPY app/ app/

RUN go mod download

RUN go build -o mygrpcapp .

FROM alpine AS final

USER nobody:nobody

COPY --from=builder /go/src/app/mygrpcapp /usr/local/bin/mygrpcapp

EXPOSE 50051

CMD ["/usr/local/bin/mygrpcapp"]
