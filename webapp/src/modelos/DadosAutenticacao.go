// Package modelos centraliza os modelos de dados da aplicação web
package modelos

// DadosAutenticacao contém o id e o token do usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"tokan"`
}
