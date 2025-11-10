package models

import "time"

type User struct {
	ID          string    `json:"id"`
	SupabaseID  string    `json:"supabase_id"`
	Email       string    `json:"email,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
