package controllers

import (
	"net/http"

	"webapp/src/cookies"
)

// FazerLogout efetua o logout do usuário da plataforma
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusFound)
}
