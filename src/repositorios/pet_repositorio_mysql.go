package repositorios

import (
	"database/sql"
	"errors"
	"fmt"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type PetRepositorioMySql struct {
	DB *sql.DB
}

// Lista todos os pets
func (pr *PetRepositorioMySql) ListaPetView() ([]model_views.PetView, error) {
	var pets []model_views.PetView

	rows, err := pr.DB.Query(
		"SELECT pets.id, pets.nome, pets.DonoId, donos.nome as Dono, pets.tipo FROM pets " +
			"JOIN donos ON pets.DonoId = donos.id")

	if err != nil {
		fmt.Println("SQL ERROR: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pet model_views.PetView
		if err := rows.Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Dono, &pet.Tipo); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (ar *PetRepositorioMySql) Where(filtros map[string]string) ([]models.Pet, error) {
	var pets []models.Pet

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

	queryBase := "SELECT id, nome, email, senha FROM pets"
	queryWhere := " WHERE " + strings.Join(whereClauses, " AND ")

	// Executa a consulta com os filtros aplicados
	rows, err := ar.DB.Query(queryBase+queryWhere, valores...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Processa os resultados
	for rows.Next() {
		var pet models.Pet
		if err := rows.Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Tipo); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

// Lista todos os pets
func (pr *PetRepositorioMySql) Lista() ([]models.Pet, error) {
	var pets []models.Pet

	rows, err := pr.DB.Query("SELECT id, nome, DonoId, tipo FROM pets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pet models.Pet
		if err := rows.Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Tipo); err != nil {
			return nil, err
		}
		pets = append(pets, pet)
	}

	return pets, nil
}

func (pr *PetRepositorioMySql) BuscarPorId(id string) (*models.Pet, error) {
	var pet models.Pet

	// Prepara a consulta SQL para buscar o pet pelo ID
	query := "SELECT id, nome, donoId, tipo FROM pets WHERE id = ?"
	err := pr.DB.QueryRow(query, id).Scan(&pet.Id, &pet.Nome, &pet.DonoId, &pet.Tipo)

	if err != nil {
		if err == sql.ErrNoRows {
			// Nenhum resultado encontrado
			return nil, nil
		}
		// Algum outro erro ocorreu
		return nil, err
	}

	return &pet, nil
}

// Adiciona um novo pet
func (pr *PetRepositorioMySql) Adicionar(pet models.Pet) error {
	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := pr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	_, err := pr.DB.Exec("INSERT INTO pets (id, nome, donoId, tipo) VALUES (?, ?, ?, ?)",
		pet.Id, pet.Nome, pet.DonoId, pet.Tipo)

	return err
}

// Altera um pet existente
func (pr *PetRepositorioMySql) Alterar(pet models.Pet) error {
	erro := pr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	_, err := pr.DB.Exec("UPDATE pets SET nome = ?, donoId = ?, tipo = ? WHERE id = ?",
		pet.Nome, pet.DonoId, pet.Tipo, pet.Id)

	return err
}

// Exclui um pet pelo ID
func (pr *PetRepositorioMySql) Excluir(id string) error {
	_, err := pr.DB.Exec("DELETE FROM pets WHERE id = ?", id)
	return err
}

func (pr *PetRepositorioMySql) validaCampos(pet *models.Pet) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome do pet é obrigatório")
	}

	if strings.TrimSpace(pet.DonoId) == "" {
		return errors.New("O dono do pet é obrigatório")
	}

	return nil
}
