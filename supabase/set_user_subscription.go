package supabase

import (
	"os"

	"github.com/supabase/postgrest-go"
)

type Profile struct {
	IsSubscribed bool   `json:"is_subscribed"`
	Interval     string `json:"interval"`
}

func SetUserSubscription(id string, isSubscribed bool, interval string) error {
	client := postgrest.NewClient(os.Getenv("SUPABASE_REST_URL"), "public", map[string]string{"apikey": os.Getenv("SUPABASE_SERVICE_KEY"), "Authorization": "Bearer " + os.Getenv("SUPABASE_SERVICE_KEY")})
	if client.ClientError != nil {
		panic(client.ClientError)
	}

	_, _, err := client.From("profile").Update(map[string]bool{"is_subscribed": isSubscribed}, "", "").Eq("id", id).Execute()
	if err != nil {
		return err
	}
	return nil
}
