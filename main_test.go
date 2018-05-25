package main

import "testing"

func TestGetWebhook(t *testing.T) {
	t.Run("error if webhook is not found", func(t *testing.T) {
		_, err := getWebhook("BORAT_WEBHOOK")
		if err == nil {
			t.Error("want an error got none")
		}
	})

	t.Run("correct error is raised if webhook env var not found", func(t *testing.T) {
		want := "webhook env var not found"
		_, err := getWebhook("BORAT_WEBHOOK")
		if err.Error() != want {
			t.Errorf("got '%s' want '%s'", err, want)
		}
	})
}

func TestImportTweets(t *testing.T) {
	t.Run("error if tweet csv not found", func(t *testing.T) {
		csvPath := "./devops_borat_tweets.csv"
		_, err := importTweets(csvPath)
		if err == nil {
			t.Error("want an error got none")
		}
	})
	t.Run("correct error is raised if tweet csv not found", func(t *testing.T) {
		csvPath := "./devops_borat_tweets.csv"
		want := "tweet csv not found"
		_, err := importTweets(cs)

	})
}
