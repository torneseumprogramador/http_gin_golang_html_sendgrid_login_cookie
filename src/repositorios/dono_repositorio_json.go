package repositorios

import (
	"encoding/json"
	"errors"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_DONOS = "db/donos.json"

type DonoRepositorioJson struct{}

func (dr *DonoRepositorioJson) Lista() ([]models.Dono, error) {
	var pets []models.Dono

	bytes, err := os.ReadFile(CAMINHO_JSON_DONOS)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Dono{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (dr *DonoRepositorioJson) BuscarPorId(id string) (*models.Dono, error) {
	pets, erro := dr.Lista()
	pet, _ := dr.buscarPorIdStruct(pets, id)
	return pet, erro
}

func (dr *DonoRepositorioJson) Salvar(pets []models.Dono) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_DONOS, bytes, 0644)
}

func (dr *DonoRepositorioJson) Adicionar(pet models.Dono) error {
	pets, err := dr.Lista()
	if err != nil {
		return err
	}

	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := dr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets = append(pets, pet)
	return dr.Salvar(pets)
}

func (dr *DonoRepositorioJson) Alterar(pet models.Dono) error {
	pets, err := dr.Lista()
	if err != nil {
		return err
	}

	index, err := dr.buscarPorId(pets, pet.Id)
	if err != nil {
		return err
	}

	erro := dr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets[index].Nome = pet.Nome
	pets[index].Telefone = pet.Telefone

	return dr.Salvar(pets)
}

func (dr *DonoRepositorioJson) Excluir(id string) error {
	pets, err := dr.Lista()
	if err != nil {
		return err
	}

	index, err := dr.buscarPorId(pets, id)
	if err != nil {
		return err
	}

	pets = append(pets[:index], pets[index+1:]...)
	return dr.Salvar(pets)
}

func (dr *DonoRepositorioJson) buscarPorId(pets []models.Dono, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (dr *DonoRepositorioJson) buscarPorIdStruct(pets []models.Dono, id string) (*models.Dono, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (dr *DonoRepositorioJson) validaCampos(pet *models.Dono) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome é obrigatório")
	}

	if strings.TrimSpace(pet.Telefone) == "" {
		return errors.New("O telefone obrigatório")
	}

	return nil
}
