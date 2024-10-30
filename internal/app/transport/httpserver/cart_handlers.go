package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/northwindman/book-shop/internal/app/common/server"
)

func (h HttpServer) UpdateCart(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		server.BadRequest("invalid-user", err, w, r)
		return
	}

	var cartRequest CartRequest
	if err := json.NewDecoder(r.Body).Decode(&cartRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	_, err = h.userService.GetUserByID(r.Context(), user.ID())
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	cart, err := toDomainCart(user.ID(), cartRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	updatedCart, err := h.cartService.UpdateCartAndStocks(r.Context(), cart)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseCart(updatedCart)

	server.RespondOK(response, w, r)
}

func (h HttpServer) Checkout(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromContext(r.Context())
	if err != nil {
		server.BadRequest("invalid-user", err, w, r)
		return
	}

	err = h.cartService.Checkout(r.Context(), user.ID())
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]bool{"ok": true}, w, r)
}
