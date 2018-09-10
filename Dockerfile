FROM golang:stretch as builder
COPY . /src
WORKDIR /src
RUN go test -v
RUN go build -o borat *.go

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /src/borat .
COPY --from=builder /src/devops_borat_tweets.txt .
RUN touch crontab.tmp \
    && echo '0 8 * * * /root/borat' > crontab.tmp \
    && echo '0 15 * * * /root/borat' >> crontab.tmp \
    && crontab crontab.tmp \
    && rm -rf crontab.tmp
CMD ["crond", "-f"]
