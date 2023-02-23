package supabase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

func CreateUserWithEmail(email string) (string, error) {
	s := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"))

	ctx := context.Background()
	_, err := s.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: uuid.New().String(),
	})
	if err != nil {
		return "", err
	}

	// sleep 2 seconds to wait for the user to be created (trigger)
	time.Sleep(2 * time.Second)

	userId, err := GetUserIdByEmail(email)
	if err != nil {
		return "", err
	}
	fmt.Println("CreateUserWithEmail, userID")
	fmt.Println(userId)
	return userId, nil
}
