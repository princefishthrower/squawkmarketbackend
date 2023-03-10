package twitter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func StreamTweetsOfUser(twitterHandle string) {

	// Set up your Twitter API credentials
	bearerToken := os.Getenv("TWITTER_BEARER_TOKEN")

	// Set up the tweet stream request
	u, err := url.Parse(fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", twitterHandle))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	q := u.Query()
	q.Set("expansions", "author_id")
	q.Set("tweet.fields", "created_at,text")
	q.Set("user.fields", "username")
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	// Make the tweet stream request and read the response line by line
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			break
		}

		// Parse the tweet JSON and check if it's from the user we're interested in
		var tweet struct {
			Data struct {
				Text      string `json:"text"`
				CreatedAt string `json:"created_at"`
				AuthorID  string `json:"author_id"`
			} `json:"data"`
			Includes struct {
				Users []struct {
					ID       string `json:"id"`
					Username string `json:"username"`
				} `json:"users"`
			} `json:"includes"`
		}
		err = json.Unmarshal(line, &tweet)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if strings.EqualFold(tweet.Includes.Users[0].Username, twitterHandle) {
			fmt.Printf("%s: %s\n", tweet.Data.CreatedAt, tweet.Data.Text)
		}
	}
}
