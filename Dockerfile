FROM golang:alpine3.7 as builder
COPY . /src
WORKDIR /src
RUN go test -v
RUN go build *.go

FROM alpine:3.7
WORKDIR /root/
COPY --from=builder /src/main .
CMD ["./main"]