package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int32
	Nome  string
	Email string
}

// POST METHOD
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Pegando o corpo da requisição
	req, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requsição"))
		return
	}

	// Criando um usuário e copiando as informações passadas no body da requisição
	var usuario usuario
	if erro := json.Unmarshal(req, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao converter o usuário da requisição para struct"))
		return
	}

	// Criando uma conexão com o DB
	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// Criando um Prepare Statement
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	// Executando o Prepare Statement e recebendo o retorno do DB
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	// Pegando o Id o Id da ultima inserção
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao obter o Id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

// GET ALL METHOD
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao executar a query no banco de dados"))
		return
	}
	defer linhas.Close()

	// Criando um slice de usuários e o populando.
	// A função Scan é utilizada para, durante a iteração das linhas da tabela usuário, ele clonar os valores e por no struct usuário.
	// o Scan faz essa copia dos valores respeitando a sequência da tabela do banco de dados.
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

// GET ONE METHOD
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseUint(param["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar a query no banco de dados"))
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Erro ao encontar o usuário. ID inválido"))
		return
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao converter o usuário para JSON"))
		return
	}

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseInt(param["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	req, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requsição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(req, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao converter o usuário da requisição para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseInt(param["id"], 10, 32)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao conectar com o banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
