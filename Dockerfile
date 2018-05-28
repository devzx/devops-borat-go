FROM golang:alpine3.7 as builder
COPY . /src
WORKDIR /src
RUN go test -v
RUN go build *.go

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /src/main .
COPY --from=builder /src/devops_borat_tweets.txt .
ENV TWEET_FILE ./devops_borat_tweets.txt
ENV SLACK_WEBHOOK REPLACE_WITH_SLACK_WEBHOOK
ENV DISCORD_WEBHOOK REPLACE_WITH_DISCORD_WEBHOOK
RUN touch crontab.tmp \
    && echo '0 8 * * * /root/main' > crontab.tmp \
    && echo '0 15 * * * /root/main' >> crontab.tmp \
    && crontab crontab.tmp \
    && rm -rf crontab.tmp
CMD ["crond", "-f"]