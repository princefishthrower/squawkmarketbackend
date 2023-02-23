package supabase

import (
	"context"
	"os"

	supa "github.com/nedpals/supabase-go"
)

func SendUserMagicLinkToEmail(email string) error {
	s := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_SERVICE_KEY"))

	ctx := context.Background()
	err := s.Auth.SendMagicLink(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
