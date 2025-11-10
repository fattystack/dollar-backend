package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fattystack/dollar-backend/internal/middleware"
	"github.com/fattystack/dollar-backend/internal/models"
	"github.com/fattystack/dollar-backend/internal/repository"
)

type MeHandler struct {
	UserRepo *repository.UserRepository
}

func NewMeHander(repo *repository.UserRepository) *MeHandler {
	return &MeHandler{UserRepo: repo}
}

func (h *MeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO(1): get the request context (you'll need this for repo calls and for reading auth info)
	context := r.Context()
	// TODO(2): read the Supabase user ID (or whatever your auth middleware stored) from the context
	//           - if it's missing, return 401
	//           - if it's not a string / empty, also return 401
	//           - this step makes sure only authenticated users can hit /me
	userID := context.Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// TODO(3): using the user ID, try to fetch the user from the database via the repository
	//           - call your repo method (e.g. GetBySupabaseID)
	//           - this is the "does this user already exist?" check
	supabaseID, ok := userID.(string)
	if !ok || supabaseID == "" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	user, err := h.UserRepo.GetByID(context, supabaseID)
	if err != nil {
		if err == repository.ErrNotFound {
			createParams := map[string]string{"supabase_id": supabaseID, "email", context.Value(middleware.SupabaseJWT)}
		}
	}

	// TODO(4): if the repository says "not found", create a new user
	//           - build the params the repo needs (supabaseID, maybe email/display name)
	//           - call repo.Create(...)
	//           - if creation fails, return 500
	//           - this makes /me idempotent: first call creates, later calls just read

	// TODO(5): if the repository returned some OTHER error (not "not found"), return 500
	//           - this is for DB/network/unexpected errors
	//           - don't leak internal errors to the client

	// TODO(6): at this point you have a user (either fetched or newly created)
	//           - write it back as JSON
	//           - set Content-Type to application/json
	//           - encode the user struct into the response

	// TODO(7): optional: if you have a shared JSON writing helper, use it here to keep things tidy
}
