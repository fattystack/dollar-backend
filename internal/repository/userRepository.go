package repository

import (
	"context"
	"errors"

	"github.com/fattystack/dollar-backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("user not found")

type UserRepository struct {
	db *pgxpool.Pool
}

type createUserParams struct {
	SupabaseID  string
	Email       string
	DisplayName string
}

func (r *UserRepository) getUserRepoReturn(ctx context.Context, q string, supabaseID string) (models.User, error) {
	var u models.User

	err := r.db.QueryRow(ctx, q, supabaseID).Scan(
		&u.ID,
		&u.SupabaseID,
		&u.Email,
		&u.DisplayName,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return models.User{}, err
	}
	return u, err
}

func (r *UserRepository) GetByID(ctx context.Context, supabaseID string) (models.User, error) {
	const q = ` 
		SELECT 
     		ID, 
       		supabase_id as SupabaseID,
       		Email,
       		display_name as DisplayName,
       		created_at as CreatedAt,
       		updated_at as UpdatedAt,
       from public.users
       where supabase_id = $1`

	return r.getUserRepoReturn(ctx, q, supabaseID)
}

func (r *UserRepository) Create(ctx context.Context, p createUserParams) (models.User, error) {

	const q = `
  INSERT INTO public.users (supabase_id, email, display_name)
  VALUES ($1, $2, $3)
  RETURNING ID, supabase_id, email, display_name, created_at, updated_at;
`
	return r.getUserRepoReturn(ctx, q, p.SupabaseID)
}
