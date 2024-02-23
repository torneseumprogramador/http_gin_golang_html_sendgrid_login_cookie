package repositorios

import (
	"encoding/json"
	"errors"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_ADMINISTRADORES = "db/administradores.json"

type AdministradorRepositorioJson struct{}

func (ar *AdministradorRepositorioJson) Lista() ([]models.Administrador, error) {
	var pets []models.Administrador

	bytes, err := os.ReadFile(CAMINHO_JSON_ADMINISTRADORES)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Administrador{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (ar *AdministradorRepositorioJson) BuscarPorId(id string) (*models.Administrador, error) {
	pets, erro := ar.Lista()
	pet, _ := ar.buscarPorIdStruct(pets, id)
	return pet, erro
}

func (ar *AdministradorRepositorioJson) Salvar(pets []models.Administrador) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_ADMINISTRADORES, bytes, 0644)
}

func (ar *AdministradorRepositorioJson) Adicionar(pet models.Administrador) error {
	pets, err := ar.Lista()
	if err != nil {
		return err
	}

	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := ar.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets = append(pets, pet)
	return ar.Salvar(pets)
}

func (ar *AdministradorRepositorioJson) Alterar(pet models.Administrador) error {
	pets, err := ar.Lista()
	if err != nil {
		return err
	}

	index, err := ar.buscarPorId(pets, pet.Id)
	if err != nil {
		return err
	}

	erro := ar.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets[index].Nome = pet.Nome
	pets[index].Email = pet.Email
	pets[index].Senha = pet.Senha

	return ar.Salvar(pets)
}

func (ar *AdministradorRepositorioJson) Excluir(id string) error {
	pets, err := ar.Lista()
	if err != nil {
		return err
	}

	index, err := ar.buscarPorId(pets, id)
	if err != nil {
		return err
	}

	pets = append(pets[:index], pets[index+1:]...)
	return ar.Salvar(pets)
}

func (ar *AdministradorRepositorioJson) buscarPorId(pets []models.Administrador, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (ar *AdministradorRepositorioJson) BuscarPorEmailSenha(email string, senha string) *models.Administrador {
	lista, _ := ar.Lista()
	for _, adm := range lista {
		if adm.Email == email && adm.Senha == senha {
			return &adm
		}
	}

	return nil
}

func (ar *AdministradorRepositorioJson) buscarPorIdStruct(pets []models.Administrador, id string) (*models.Administrador, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (ar *AdministradorRepositorioJson) validaCampos(pet *models.Administrador) error {
	if pet.Id == "" {
		return errors.New("O ID de identificação, não pode ser vazio")
	}

	if strings.TrimSpace(pet.Nome) == "" {
		return errors.New("O nome é obrigatório")
	}

	if strings.TrimSpace(pet.Email) == "" {
		return errors.New("O email obrigatório")
	}

	if strings.TrimSpace(pet.Senha) == "" {
		return errors.New("A Senha obrigatória")
	}

	return nil
}
