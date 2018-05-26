package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	filePath = "./devops_borat_tweets_test.txt"
)

func TestGetWebhook(t *testing.T) {
	t.Run("error raised if webhook is not found", func(t *testing.T) {
		t.Helper()
		_, err := getWebhook("BORAT_WEBHOOK")
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("correct error message is raised if webhook env var not found", func(t *testing.T) {
		t.Helper()
		want := "webhook env var not found"
		_, err := getWebhook("BORAT_WEBHOOK")
		if err.Error() != want {
			t.Errorf("got '%s' want '%s'", err, want)
		}
	})
}

func TestOpenTweetFile(t *testing.T) {
	os.Remove(filePath)
	defer os.Remove(filePath)

	t.Run("error raised if tweet file is not found", func(t *testing.T) {
		t.Helper()
		f, err := openTweetFile(filePath)
		defer f.Close()
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("error raised if tweet file is empty", func(t *testing.T) {
		t.Helper()
		_, err := os.Create(filePath)
		if err != nil {
			t.Fatal(err)
		}
		_, err = openTweetFile(filePath)
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("correct error message if tweet file is empty", func(t *testing.T) {
		t.Helper()
		_, err := os.Create(filePath)
		if err != nil {
			t.Fatal(err)
		}
		want := "tweet file is empty"
		_, err = openTweetFile(filePath)
		if err.Error() != want {
			t.Errorf("got '%s' want '%s'", err, want)
		}
	})
}

func TestReadTweets(t *testing.T) {
	os.Remove(filePath)
	defer os.Remove(filePath)
	t.Run("read tweets in to tweets struct", func(t *testing.T) {
		tdTweets := []struct {
			tweet string
		}{
			{tweet: "random tweet"},
			{tweet: "my tweet"},
			{tweet: "heyyy"},
		}
		for _, tt := range tdTweets {
			f, err := os.Create(filePath)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Fprintf(f, "%s\n", tt.tweet)

			tweetFileData, _ := openTweetFile(filePath)
			tweets := NewTweets(tweetFileData)
			if tweets.tweets[0] != tt.tweet {
				t.Errorf("got '%s' want '%s'", tweets.tweets[0], tt.tweet)
			}
			f.Close()
		}
	})
}
