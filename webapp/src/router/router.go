// Package router é responsável por centralizar métodos relacionados as rotas da aplicação
package router

import "github.com/gorilla/mux"

// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
