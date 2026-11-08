// Package controllers centraliza todos os métodos relacionados ao fluxo da aplicação web
package controllers

import "net/http"

// CarregarTelaDeLogin renderiza a tela de login da aplicação web
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tela de Login"))
}
