// Package controllers centraliza todos os métodos relacionados ao fluxo da aplicação web
package controllers

import (
	"net/http"

	"webapp/src/utils"
)

// CarregarTelaDeLogin renderiza a tela de login da aplicação web
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarPaginaDeCadastroDeUsuario renderiza a tela de cadastro de usuários na aplicação
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
