package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		return nil, fmt.Errorf("Error verifying credentials: %v", err)
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)

	return client, nil
}

func main() {
	fmt.Println("Hello guys, This is TW bot test")
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	client, err := getClient(&creds)

	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}
	tweet, resp, err := client.Statuses.Update("Hello guys, This is TW bot test", nil)

	if err != nil {
		log.Println(err)
	}
	log.Println(tweet)

	log.Printf("%+v\n", resp)

}
