FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY base.env .

COPY app/ app/

RUN go mod download

RUN CGO_ENABLED=0 GOFLAGS=-mod=mod GOOS=linux go build -ldflags="-w -s" -a -o /mygrpcapp .

FROM alpine AS final

USER nobody:nobody

COPY --chown=nobody:nobody --from=builder /mygrpcapp /mygrpcapp

EXPOSE 50051

CMD ["/mygrpcapp"]
