package tests

import (
	"squawkmarketbackend/supabase"
	"testing"

	"github.com/joho/godotenv"
)

func TestSupabaseFunctions(t *testing.T) {

	// need to load in the env for this test
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error(err)
	}

	// first create a user
	email := "frewin.christopher@gmail.com"
	userId, err := supabase.CreateUserWithEmail(email)
	if err != nil {
		t.Error(err)
	}
	if userId == "" {
		t.Error("userId is empty")
	}

	// then ensure we can get the user's id from their email
	userId, err = supabase.GetUserIdByEmail(email)
	if err != nil {
		t.Error(err)
	}
	if userId == "" {
		t.Error("userId is empty")
	}

	// also ensure we can send a magic link to the user
	err = supabase.SendUserMagicLinkToEmail(email)
	if err != nil {
		t.Error(err)
	}

	// then set the user's subscription
	err = supabase.SetUserSubscription(userId, true, "monthly")
	if err != nil {
		t.Error(err)
	}

	// then unset the user's subscription
	err = supabase.SetUserSubscription(userId, false, "monthly")
	if err != nil {
		t.Error(err)
	}
}
