// Package controllers centraliza todos os métodos relacionados ao fluxo da aplicação web
package controllers

import (
	"fmt"
	"net/http"

	"webapp/src/config"
	"webapp/src/requisicoes"
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

// CarregarPaginaPrincipal carrega a pagina principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	fmt.Println(response, erro)
	utils.ExecutarTemplate(w, "home.html", nil)
}
