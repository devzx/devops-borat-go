package main

import (
	"errors"
	"os"
)

var (
	errWebhookEnvVarNotFound = errors.New("webhook env var not found")
	errTweetCSVNotFound      = errors.New("tweet csv not found")
)

type tweet struct {
}

func getWebhook(name string) (webhook string, err error) {
	if _, ok := os.LookupEnv(name); !ok {
		return "", errWebhookEnvVarNotFound
	}
	return
}

func importTweets(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errTweetCSVNotFound
	}
	return f, nil
}

// INIT
//   Import all tweets
//   Get channel name

// MAIN
//   Select a random tweet -- Struct tweet -- struct contains JSON data
//   Post tweet to channel  -- Method on tweet struct - compile payload and post to channel
