package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
	"time"
)

const (
	// Create an app from https://apps.twitter.com/, create an access token
	// and copy-paste the values from there.
	ConsumerKey       = ""
	ConsumerSecret    = ""
	AccessToken       = "-"
	AccessTokenSecret = ""
)

func ReturnList(client *twitter.Client, user string, f *os.File) string {
	friends, _, err := client.Friends.IDs(&twitter.FriendIDParams{
		UserID:     0,
		ScreenName: user,
		Cursor:     0,
		Count:      0,
	})
	if err != nil {
		fmt.Println("Sleeping ...")
		time.Sleep(3 * time.Minute)
		return ReturnList(client, user, f)
	}

	b, err := json.Marshal(friends)
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(user + ": " + string(b) + "\n")

	return "Done"
}

func main() {

	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	f, err := os.Create("scrowder.txt")
	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	file, _ := os.Open("users.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var users []string

	for scanner.Scan() {
		users = append(users, scanner.Text())
	}

	for _, user := range users {
		i := ReturnList(client, user, f)
		fmt.Println(user + ": " + i)
	}

}
