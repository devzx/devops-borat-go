package main

import (
	"os"
	"testing"
)

func TestGetWebhook(t *testing.T) {
	t.Run("error if webhook is not found", func(t *testing.T) {
		// Set this env var to nothing in case it already exists
		_, err := getWebhook("BORAT_WEBHOOK")
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("correct error is raised if webhook env var not found", func(t *testing.T) {
		// Set this env var to nothing in case it already exists
		want := "webhook env var not found"
		_, err := getWebhook("BORAT_WEBHOOK")
		if err.Error() != want {
			t.Errorf("got '%s' want '%s'", err, want)
		}
	})
}

func TestImportTweets(t *testing.T) {
	csvPath := "./devops_borat_tweets_test.csv"
	os.Remove(csvPath)

	t.Run("error if tweet csv not found", func(t *testing.T) {
		_, err := importTweets(csvPath)
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("correct error is raised if tweet csv not found", func(t *testing.T) {
		want := "tweet csv not found"
		_, err := importTweets(csvPath)
		if err.Error() != want {
			t.Errorf("got '%s' want '%s'", err, want)
		}
	})

	t.Run("correct data is returned from tweet csv", func(t *testing.T) {
		// Create mock tweet data save to disk -- Read in and verify correct columns are there -- possibly don't actually write to disk but mock it
	})
}
