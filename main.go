package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	errWebhookEnvVarNotFound = errors.New("webhook env var not found")
	errTweetFileEmpty        = errors.New("tweet file is empty")
)

func getWebhook(name string) (webhook string, err error) {
	if _, ok := os.LookupEnv(name); !ok {
		return "", errWebhookEnvVarNotFound
	}
	return
}

func openTweetFile(file string) (*os.File, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	} else if size, _ := f.Stat(); size.Size() == 0 {
		return nil, errTweetFileEmpty
	}
	return f, nil
}

type tweets struct {
	tweets []string
}

func NewTweets(tweetFile *os.File) *tweets {
	tweetsRead, err := ioutil.ReadAll(tweetFile)
	defer tweetFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	t := &tweets{tweets: strings.Split(strings.TrimSpace(string(tweetsRead)), "\n")}
	return t
}

// INIT
//   Get tweet file name
//   Open tweet file
//   Read tweets in to struct
//   Get channel name
// MAIN
//   Select a random tweet -- Struct tweet -- struct contains JSON data
//   Post tweet to channel  -- Method on tweet struct - compile payload and post to channel
