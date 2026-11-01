package controllers

import "net/http"

// CriarUsuario efetua a criação de um usuário dentro do banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// BuscarUsuarios lista todos os usuários existentes dentro do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// BuscarUsuario efetua a busca dos dados de um usuário em específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario atualiza os dados de um usuário específico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeletarUsuario realiza a exclusão dos dados de um usuário cadastrado
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
