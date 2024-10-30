package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/northwindman/book-shop/internal/app/common/server"
)

func (h HttpServer) SignUp(w http.ResponseWriter, r *http.Request) {
	var authRequest AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := authRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	hashedPassword, err := hashPassword(authRequest.Password)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	user, err := toDomainUser(authRequest.Username, hashedPassword)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	_, err = h.userService.CreateUser(r.Context(), user)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]bool{"ok": true}, w, r)
}

func (h HttpServer) SignIn(w http.ResponseWriter, r *http.Request) {
	var authRequest AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := authRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	user, err := h.userService.GetUser(r.Context(), authRequest.Username)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	if !checkPasswordHash(authRequest.Password, user.Password()) {
		server.BadRequest("invalid-password", nil, w, r)
		return
	}

	token, err := h.tokenService.GenerateToken(user)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]string{"token": token}, w, r)
}
