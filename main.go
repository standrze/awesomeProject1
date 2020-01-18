package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

const (
	// Create an app from https://apps.twitter.com/, create an access token
	// and copy-paste the values from there.
	ConsumerKey       = ""
	ConsumerSecret    = ""
	AccessToken       = ""
	AccessTokenSecret = ""
)

func ReturnList(count int, client *twitter.Client, f *os.File) int {
	friends, _, err := client.Friends.List(&twitter.FriendListParams{
		UserID:              0,
		ScreenName:          "PeteButtigieg",
		Cursor:              0,
		Count:               count,
		SkipStatus:          nil,
		IncludeUserEntities: nil,
	})
	if err != nil {
		log.Panic(err)
		return count
	}
	for _, friend := range friends.Users {
		b, err := json.Marshal(friend)
		if err != nil {
			fmt.Println(err)
		}
		f.WriteString(string(b) + "\n")
	}

	fmt.Println(int(friends.NextCursor))
	return ReturnList(int(friends.NextCursor), client, f)
}

func main() {

	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	f, err := os.Create("PeteButtigieg2_following.txt")
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()
	i := ReturnList(0, client, f)
	fmt.Println(i)

}

