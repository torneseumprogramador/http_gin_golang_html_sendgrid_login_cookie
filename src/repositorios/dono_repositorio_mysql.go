package repositorios

import (
	"database/sql"
	"errors"
	"fmt"
	"http_gin/src/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type DonoRepositorioMySql struct {
	DB *sql.DB
}

// Lista todos os donos
func (dr *DonoRepositorioMySql) Lista() ([]models.Dono, error) {
	var donos []models.Dono

	rows, err := dr.DB.Query("SELECT id, nome, telefone FROM donos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dono models.Dono
		if err := rows.Scan(&dono.Id, &dono.Nome, &dono.Telefone); err != nil {
			return nil, err
		}
		donos = append(donos, dono)
	}

	return donos, nil
}

func (ar *DonoRepositorioMySql) Where(filtros map[string]string) ([]models.Dono, error) {
	var donos []models.Dono

	// Verifica se há filtros para aplicar
	if len(filtros) == 0 {
		return nil, fmt.Errorf("nenhum filtro fornecido")
	}

	// Constrói a cláusula WHERE dinamicamente
	var whereClauses []string
	var valores []interface{}

	for chave, valor := range filtros {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", chave))
		valores = append(valores, valor)
	}

	queryBase := "SELECT id, nome, email, senha FROM donos"
	queryWhere := " WHERE " + strings.Join(whereClauses, " AND ")

	// Executa a consulta com os filtros aplicados
	rows, err := ar.DB.Query(queryBase+queryWhere, valores...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados
	for rows.Next() {
		var dono models.Dono
		if err := rows.Scan(&dono.Id, &dono.Nome, &dono.Telefone); err != nil {
			return nil, err
		}
		donos = append(donos, dono)
	}

	return donos, nil
}

// Adiciona um novo dono
func (dr *DonoRepositorioMySql) Adicionar(dono models.Dono) error {
	if dono.Id == "" {
		dono.Id = uuid.New().String()
	}

	erro := dr.validaCampos(&dono)
	if erro != nil {
		return erro
	}

	_, err := dr.DB.Exec("INSERT INTO donos (id, nome, telefone) VALUES (?, ?, ?)",
		dono.Id, dono.Nome, dono.Telefone)

	return err
}

func (dr *DonoRepositorioMySql) BuscarPorId(id string) (*models.Dono, error) {
	var dono models.Dono

	// Prepara a consulta SQL para buscar o dono pelo ID
	query := "SELECT id, nome, telefone FROM donos WHERE id = ?"
	err := dr.DB.QueryRow(query, id).Scan(&dono.Id, &dono.Nome, &dono.Telefone)

	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return nil, nil
		}
		// Algum outro erro ocorreu
		return nil, err
	}

	return &dono, nil
}

// Altera um dono existente
func (dr *DonoRepositorioMySql) Alterar(dono models.Dono) error {
	erro := dr.validaCampos(&dono)
	if erro != nil {
		return erro
	}

	_, err := dr.DB.Exec("UPDATE donos SET nome = ?, telefone = ? WHERE id = ?",
		dono.Nome, dono.Telefone, dono.Id)

	return err
}

// Exclui um dono pelo ID
func (dr *DonoRepositorioMySql) Excluir(id string) error {
	_, err := dr.DB.Exec("DELETE FROM donos WHERE id = ?", id)
	return err
}

func (dr *DonoRepositorioMySql) validaCampos(dono *models.Dono) error {
	if dono.Id == "" {
		return errors.New("O ID de identificação não pode ser vazio")
	}

	if strings.TrimSpace(dono.Nome) == "" {
		return errors.New("O nome é obrigatório")
	}

	if strings.TrimSpace(dono.Telefone) == "" {
		return errors.New("O telefone é obrigatório")
	}

	return nil
}
