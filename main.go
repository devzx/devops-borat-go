package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	errWebhookEnvVarNotFound       = errors.New("webhook env var not found")
	errTweetFilePathEnvVarNotFound = errors.New("tweet file path env var not found")
	errTweetFileEmpty              = errors.New("tweet file is empty")
)

func getEnvVar(name string, err error) (string, error) {
	if webhook, ok := os.LookupEnv(name); ok {
		return webhook, nil
	}
	return "", err
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
	data := strings.Split(string(tweetsRead), "\n")
	// When spliting an emtpy new line is added as an
	// element to the slice, this removes that
	data = data[:(len(data) - 1)]
	t := &tweets{tweets: data}
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
