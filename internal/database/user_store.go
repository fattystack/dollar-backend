package database

import (
	"context"
	"fmt"

	"github.com/fattystack/dollar-backend/internal/models"
)

type UserStore struct{}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) GetBySupabaseID(ctx context.Context, supabaseID string) (*models.User, error) {
	var users []models.User

	_, err := Client.From("users").
		Select("*", "", false).
		Eq("supabase_id", supabaseID).
		Single().
		ExecuteTo(&users)

	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}
