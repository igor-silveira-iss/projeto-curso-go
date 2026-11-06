// Package repositorios
package repositorios

import (
	"database/sql"
	"fmt"

	"api/src/modelos"
)

// Usuarios representa um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	stmt, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?);",
	)
	if erro != nil {
		return 0, erro
	}
	defer stmt.Close()

	resultado, erro := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Buscar realiza a busca dos usuários do sistema que atendem ao parâmetro enviado
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?;",
		nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// BuscarPorID realiza a busca de um usuário cadastrado no banco de dados através do parâmetro usuarioID
func (repositorio Usuarios) BuscarPorID(usuarioID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = ?",
		usuarioID,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar altera os dados de um usuário cadastrado no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	stmt, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui os dados de um usuário cadastrado no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro := stmt.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail efetua a busca de um usuário do banco de dados pelo email
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query(
		"select id, senha from usuarios where email = ?;",
		email,
	)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Seguir permite que um usuario siga outros usuarios
func (repositorio Usuarios) Seguir(usuarioID uint64, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (usuario_id, seguidor_id) values(?, ?);",
	)
	if erro != nil {
		return erro
	}
	defer stmt.Close()
	if _, erro = stmt.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// PararDeSeguir permite que o usuário possa parar de seguir outros usuários
func (repositorio Usuarios) PararDeSeguir(usuarioID uint64, seguidorID uint64) error {
	stmt, erro := repositorio.db.Prepare(
		"delete from seguidores where usuario_id = ? and seguidor_id = ?;",
	)
	if erro != nil {
		return erro
	}
	defer stmt.Close()
	if _, erro = stmt.Exec(usuarioID, seguidorID); erro != nil {
		return erro
	}
	return nil
}

// BuscarSeguidores retorna os usuários seguidos
func (repositorio Usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u
		inner join seguidores s on u.id = s.seguidor_id
		where s.usuario_id = ?;
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro = linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

// BuscarSeguindo retorna quem o usuário segue
func (repositorio Usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from usuarios u
		inner join seguidores s on u.id = s.usuario_id
		where s.seguidor_id = ?;
	`, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosSeguidos []modelos.Usuario
	for linhas.Next() {
		var usuarioSeguido modelos.Usuario
		if erro = linhas.Scan(
			&usuarioSeguido.ID,
			&usuarioSeguido.Nome,
			&usuarioSeguido.Nick,
			&usuarioSeguido.Email,
			&usuarioSeguido.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuariosSeguidos = append(usuariosSeguidos, usuarioSeguido)
	}

	return usuariosSeguidos, nil
}

// BuscarSenha traz a senha do usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query(`
		select u.senha from usuarios u where u.id = ?;
	`, usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", nil
		}
	}

	return usuario.Senha, nil
}

// AtualizarSenha efetua a atualização da senha do usuário no banco de dados
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	stmt, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?;")
	if erro != nil {
		return erro
	}
	defer stmt.Close()

	if _, erro = stmt.Exec(senha, usuarioID); erro != nil {
		return erro
	}

	return nil
}
