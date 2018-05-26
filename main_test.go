package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	filePath = "./devops_borat_tweets_test.txt"
)

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("want an error got none")
	}
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Error("got an error, didn't want one")
	}
}
func TestEnvironmentalVariables(t *testing.T) {
	t.Run("webhook env var doesn't exist", func(t *testing.T) {
		_, err := getEnvVar("TEST_BORAT_SLACK_WEBHOOK", errWebhookEnvVarNotFound)
		assertError(t, err, errWebhookEnvVarNotFound)
	})

	t.Run("tweet file path env var doesn't exist", func(t *testing.T) {
		_, err := getEnvVar("TEST_BORAT_TWEET_FILE", errTweetFilePathEnvVarNotFound)
		assertError(t, err, errTweetFilePathEnvVarNotFound)
	})
}

func TestOpenTweetFile(t *testing.T) {
	os.Remove(filePath)
	defer os.Remove(filePath)

	t.Run("tweet file is not found", func(t *testing.T) {
		_, err := openTweetFile(filePath)
		assertError(t, err, err)
	})

	t.Run("tweet file is empty", func(t *testing.T) {
		_, err := os.Create(filePath)
		_, err = openTweetFile(filePath)
		assertError(t, err, errTweetFileEmpty)
	})
}

func TestReadTweets(t *testing.T) {
	os.Remove(filePath)
	defer os.Remove(filePath)
	t.Run("read tweets in to tweets struct", func(t *testing.T) {
		tdTweets := []struct {
			tweets []string
		}{
			{tweets: []string{"random tweet", "yooo"}},
			{tweets: []string{" testing "}},
			{tweets: []string{"this", "is", "", "a", "   "}},
		}
		for _, tt := range tdTweets {
			f, err := os.Create(filePath)
			if err != nil {
				t.Fatal(err)
			}
			for _, tweet := range tt.tweets {
				fmt.Fprintf(f, "%s\n", tweet)
			}

			tweetFileData, _ := openTweetFile(filePath)
			tweets := NewTweets(tweetFileData)
			if !reflect.DeepEqual(tweets.tweets, tt.tweets) {
				t.Errorf("got '%s' want '%s'", tweets.tweets, tt.tweets)
			}
			f.Close()
		}
	})
}
