package api

import (
	"ctf01d/internal/app/repository"
	api_helpers "ctf01d/internal/app/utils"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
)

type RequestLogin struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func LoginSessionHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var req RequestLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn(err.Error(), "handler", "LoginSessionHandler")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userRepo := repository.NewUserRepository(db)
	user, err := userRepo.GetByUserName(r.Context(), req.Username)
	if err != nil || !api_helpers.CheckPasswordHash(req.Password, user.PasswordHash) {
		slog.Warn(err.Error(), "handler", "LoginSessionHandler")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid password or user"})
		return
	}

	sessionRepo := repository.NewSessionRepository(db)
	sessionId, err := sessionRepo.StoreSessionInDB(r.Context(), user.Id)
	if err != nil {
		slog.Warn(err.Error(), "handler", "LoginSessionHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to store session"})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		HttpOnly: true,
		Value:    sessionId,
		Path:     "/",
		MaxAge:   96 * 3600, // fixme, брать из db
	})

	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logged in"})
}

func LogoutSessionHandler(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		slog.Warn(err.Error(), "handler", "LogoutSessionHandler")
		api_helpers.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"error": "No session found"})
		return
	}
	sessionRepo := repository.NewSessionRepository(db)
	err = sessionRepo.DeleteSessionInDB(r.Context(), cookie.Value)
	if err != nil {
		slog.Warn(err.Error(), "handler", "LogoutSessionHandler")
		api_helpers.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": "Failed to delete session"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // Удаление куки
	})
	api_helpers.RespondWithJSON(w, http.StatusOK, map[string]string{"data": "User logout successful"})
}