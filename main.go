package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	botName                  = "DevOps Borat"
	tweetFileEnvVarName      = "TWEET_FILE"
	slackWebhookEnvVarName   = "SLACK_WEBHOOK"
	discordWebhookEnvVarName = "DISCORD_WEBHOOK"
	iconURL                  = "https://pbs.twimg.com/profile_images/1079908235/borat_855_18535194_0_0_12672_300_400x400.jpg"
)

var (
	errWebhookEnvVarNotFound       = errors.New("webhook env var not found")
	errTweetFilePathEnvVarNotFound = errors.New("tweet file path env var not found")
	errTweetFileEmpty              = errors.New("tweet file is empty")
)

func getEnvVar(name string, err error) (string, error) {
	if envVar, ok := os.LookupEnv(name); ok {
		return envVar, nil
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

func (t *tweets) getTweet() string {
	rand.Seed(time.Now().UnixNano())
	return t.tweets[rand.Intn(len(t.tweets))]
}

type slack struct {
	contentType string
	webhook     string
	payload     struct {
		Username string `json:"username"`
		IconURL  string `json:"icon_url"`
		Text     string `json:"text"`
	}
}

func (s *slack) createPayload(tweet string) (*bytes.Buffer, error) {
	s.payload.Text = tweet
	s.payload.IconURL = iconURL
	s.payload.Username = botName
	b, err := json.Marshal(s.payload)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

type discord struct {
	contentType string
	webhook     string
	payload     struct {
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Content   string `json:"content"`
	}
}

func (d *discord) createPayload(tweet string) (*bytes.Buffer, error) {
	d.payload.Content = tweet
	d.payload.AvatarURL = iconURL
	d.payload.Username = botName
	b, err := json.Marshal(d.payload)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

type service interface {
	createPayload(string) (*bytes.Buffer, error)
}

func main() {
}
