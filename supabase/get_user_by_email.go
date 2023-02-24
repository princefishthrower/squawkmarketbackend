package supabase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/supabase/postgrest-go"
)

type User struct {
	ID string `json:"id"`
}

func GetUserIdByEmail(email string) (string, error) {
	client := postgrest.NewClient(os.Getenv("SUPABASE_REST_URL"), "public", map[string]string{"apikey": os.Getenv("SUPABASE_SERVICE_KEY"), "Authorization": "Bearer " + os.Getenv("SUPABASE_SERVICE_KEY")})
	if client.ClientError != nil {
		panic(client.ClientError)
	}

	res, _, err := client.From("profile").Select("id", "", false).Eq("email", email).Single().Execute()
	if err != nil {
		fmt.Println(err)
		if err.Error() == "(PGRST116) JSON object requested, multiple (or no) rows returned" {
			return "", nil
		}
		return "", err
	}

	// unmarshal res to supabaseTypes.User
	var user User
	err = json.Unmarshal(res, &user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return user.ID, nil
}
