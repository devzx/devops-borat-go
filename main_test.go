package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	filePath                   = "./test_devops_borat_tweets.txt"
	testSlackWebhookEnvVarName = "TEST_" + slackWebhookEnvVarName
	testTweetFileEnvVarName    = "TEST_" + tweetFileEnvVarName
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

func assertTrue(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestEnvironmentalVariables(t *testing.T) {
	// Convert to TDT
	t.Run("webhook env var doesn't exist", func(t *testing.T) {
		_, err := getEnvVar(testSlackWebhookEnvVarName, errWebhookEnvVarNotFound)
		assertError(t, err, errWebhookEnvVarNotFound)
	})

	t.Run("tweet file path env var doesn't exist", func(t *testing.T) {
		_, err := getEnvVar(testTweetFileEnvVarName, errTweetFileEnvVarNotFound)
		assertError(t, err, errTweetFileEnvVarNotFound)
	})

	t.Run("webhook env var exists", func(t *testing.T) {
		defer os.Unsetenv(testSlackWebhookEnvVarName)

		os.Setenv(testSlackWebhookEnvVarName, "https://web.hook.com/random/23132")
		slackWebhook, err := getEnvVar(testSlackWebhookEnvVarName, errWebhookEnvVarNotFound)
		assertNoError(t, err)
		assertTrue(t, slackWebhook, "https://web.hook.com/random/23132")
	})

	t.Run("tweet file path env var exist", func(t *testing.T) {
		defer os.Unsetenv(testTweetFileEnvVarName)

		os.Setenv(testTweetFileEnvVarName, filePath)
		tweetFile, err := getEnvVar(testTweetFileEnvVarName, errTweetFileEnvVarNotFound)
		assertNoError(t, err)
		assertTrue(t, tweetFile, filePath)
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
		if err != nil {
			t.Error(err)
		}
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

			tweetFileData, err := openTweetFile(filePath)
			assertNoError(t, err)
			tweets := NewTweets(tweetFileData)
			if !reflect.DeepEqual(tweets.tweets, tt.tweets) {
				t.Errorf("got '%s' want '%s'", tweets.tweets, tt.tweets)
			}
			f.Close()
		}
	})
}

func TestGetRandomTweet(t *testing.T) {
	os.Remove(filePath)
	defer os.Remove(filePath)

	tweetsS := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}
	f, _ := os.Create(filePath)
	defer f.Close()
	for _, tweet := range tweetsS {
		fmt.Fprintf(f, "%s\n", tweet)
	}
	tweetFileData, _ := openTweetFile(filePath)
	tweets := NewTweets(tweetFileData)
	var match int
	for i := 0; i < 100; i++ {
		randomTweet1 := tweets.getTweet()
		randomTweet2 := tweets.getTweet()
		if randomTweet1 == randomTweet2 {
			match++
		}
	}
	if match > 20 {
		t.Fatal("probably not random")
	}
}

func TestCreatePayload(t *testing.T) {
	testCases := []struct {
		desc    string
		want    string
		tweet   string
		service service
	}{
		{
			desc:    "test slack payload",
			want:    `{"username":"` + botName + `","icon_url":"` + iconURL + `","text":"test"}`,
			tweet:   "test",
			service: &slack{},
		},
		{
			desc:    "test discord payload",
			want:    `{"username":"` + botName + `","avatar_url":"` + iconURL + `","content":"test"}`,
			tweet:   "test",
			service: &discord{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := tC.service.createPayload(tC.tweet)
			assertTrue(t, got.String(), tC.want)
			assertNoError(t, err)
		})
	}
}
