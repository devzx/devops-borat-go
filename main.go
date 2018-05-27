package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	tweetFileEnvVarName      = "TWEET_FILE_NAME"
	slackWebhookEnvVarName   = "SLACK_WEBHOOK"
	discordWebhookEnvVarName = "DISCORD_WEBHOOK"
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
	payload     io.Reader
}

func (s *slack) createPayload(text, iconUrl, username string) {
	s.payload = strings.NewReader(fmt.Sprintf(`{"text": "%s", "icon_url": "%s", "username": "%s"}`, text, iconUrl, username))
}

func (s *slack) post() {
	resp, err := http.Post(s.webhook, s.contentType, s.payload)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

type discord struct {
	contentType string
	webhook     string
	payload     io.Reader
}

func (d *discord) createPayload(text, username string) {
	d.payload = strings.NewReader(fmt.Sprintf(`{"content": "%s", "username": "%s"}`, text, username))
}

func (d *discord) post() {
	resp, err := http.Post(d.webhook, d.contentType, d.payload)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp)
}

func main() {
	tweetFile, err := getEnvVar(tweetFileEnvVarName, errTweetFilePathEnvVarNotFound)
	if err != nil {
		log.Fatal(err)
	}
	slackWebhook, err := getEnvVar(slackWebhookEnvVarName, errWebhookEnvVarNotFound)
	if err != nil {
		log.Fatal(err)
	}
	discordWebhook, err := getEnvVar(discordWebhookEnvVarName, errWebhookEnvVarNotFound)
	if err != nil {
		log.Printf("discord webhook not found, continuing")
	}
	openTweetFile, err := openTweetFile(tweetFile)
	if err != nil {
		log.Fatal(err)
	}
	t := NewTweets(openTweetFile)
	tweet := t.getTweet()

	slack := &slack{
		contentType: "application/json",
		webhook:     slackWebhook,
	}
	discord := &discord{
		contentType: "application/json",
		webhook:     discordWebhook,
	}
	slack.createPayload(tweet, "https://pbs.twimg.com/profile_images/1079908235/borat_855_18535194_0_0_12672_300_400x400.jpg", "DevOps Borat")
	slack.post()
	discord.createPayload(tweet, "DevOps Borat")
	discord.post()
}
