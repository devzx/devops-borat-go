# devops-borat-go
[![Build Status](https://travis-ci.org/devzx/devops-borat-go.svg?branch=master)](https://travis-ci.org/devzx/devops-borat-go)
----

Rewrite of https://github.com/devzx/devops-borat in [Golang](https://golang.org/).  
A fun little app that posts a random [DevOps Borat](https://twitter.com/DEVOPS_BORAT) tweet to a given chat channel.  

## Why rewrite?
- Become more familiar with Go
- Become more familiar with TDD
- Implement CI with Travis
- Incorporate automated building of Docker images and pushing to Docker Hub on successful CI runs

## Dependencies
```docker```
## Usage

To build your own image run the following. At least one valid webhook (Slack or Discord) is required.
``` sh
$ git clone git@github.com:devzx/devops-borat-go.git
$ cd devops-borat-go
$ docker build -t borat .
$ docker run -d -e TWEET_FILE='./devops_borat_tweets.txt' -e SLACK_WEBHOOK='<Replace with your webhook>' -e DISCORD_WEBHOOK='<Replace with your webhook>' --name borat borat
```
If you don't fancy building your own image; here is one I created earlier :) .
```
$ docker run -d -e TWEET_FILE='./devops_borat_tweets.txt' -e SLACK_WEBHOOK='<Replace with your webhook>' -e DISCORD_WEBHOOK='<Replace with your webhook>' --name borat devzx/devops-borat
```

The application by default will post to the channel twice a day at 9AM and 4PM. You can customise this to your liking by editing the cronjob which also resides within the Dockerfile.

### Supports
- Slack
- Discord
