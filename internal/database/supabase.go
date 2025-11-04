package database

import (
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func InitSupabase(url, key string) error {
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		return err
	}

	Client = client
	return nil
}
