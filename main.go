package main

import (
	"errors"
	"os"
)

const (
	webhookEnvVarNotFound = "webhook env var not found"
	tweetCSVNotFound      = "tweet csv not found"
)

func getWebhook(name string) (webhook string, err error) {
	if _, ok := os.LookupEnv(name); !ok {
		return "", errors.New(webhookEnvVarNotFound)
	}
	return
}

func importTweets(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, errors.New(tweetCSVNotFound)
	}
	return f, nil
}
