package repositorios

import (
	"encoding/json"
	"errors"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"os"
	"strings"

	"github.com/google/uuid"
)

const CAMINHO_JSON_PATS = "db/pets.json"

type PetRepositorioJson struct{}

func (pr *PetRepositorioJson) Lista() ([]models.Pet, error) {
	var pets []models.Pet

	bytes, err := os.ReadFile(CAMINHO_JSON_PATS)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Pet{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	return pets, nil
}

func (pr *PetRepositorioJson) ListaPetView(dr DonoRepositorioJson) ([]model_views.PetView, error) {
	var pets []models.Pet
	var pets_views []model_views.PetView

	bytes, err := os.ReadFile(CAMINHO_JSON_PATS)
	if err != nil {
		if os.IsNotExist(err) {
			return []model_views.PetView{}, nil // Retorna slice vazia se o arquivo não existir
		}
		return nil, err
	}

	err = json.Unmarshal(bytes, &pets)
	if err != nil {
		return nil, err
	}

	for _, pet := range pets {

		dono, _ := dr.BuscarPorId(pet.DonoId)
		petView := model_views.PetView{}
		petView.Id = pet.Id
		petView.Nome = pet.Nome
		petView.DonoId = pet.DonoId
		petView.Dono = dono.Nome

		pets_views = append(pets_views, petView)
	}
	return pets_views, nil
}

func (pr *PetRepositorioJson) BuscarPorId(id string) (*models.Pet, error) {
	pets, erro := pr.Lista()
	pet, _ := pr.buscarPorIdStruct(pets, id)
	return pet, erro
}

func (pr *PetRepositorioJson) Salvar(pets []models.Pet) error {
	bytes, err := json.Marshal(pets)
	if err != nil {
		return err
	}

	return os.WriteFile(CAMINHO_JSON_PATS, bytes, 0644)
}

func (pr *PetRepositorioJson) Adicionar(pet models.Pet) error {
	pets, err := pr.Lista()
	if err != nil {
		return err
	}

	if pet.Id == "" {
		pet.Id = uuid.New().String()
	}

	erro := pr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets = append(pets, pet)
	return pr.Salvar(pets)
}

func (pr *PetRepositorioJson) Alterar(pet models.Pet) error {
	pets, err := pr.Lista()
	if err != nil {
		return err
	}

	index, err := pr.buscarPorId(pets, pet.Id)
	if err != nil {
		return err
	}

	erro := pr.validaCampos(&pet)
	if erro != nil {
		return erro
	}

	pets[index].Nome = pet.Nome
	pets[index].DonoId = pet.DonoId
	pets[index].Tipo = pet.Tipo

	return pr.Salvar(pets)
}

func (pr *PetRepositorioJson) Excluir(id string) error {
	pets, err := pr.Lista()
	if err != nil {
		return err
	}

	index, err := pr.buscarPorId(pets, id)
	if err != nil {
		return err
	}

	pets = append(pets[:index], pets[index+1:]...)
	return pr.Salvar(pets)
}

func (ps *PetRepositorioJson) buscarPorId(pets []models.Pet, id string) (int, error) {
	for i, pet := range pets {
		if pet.Id == id {
			return i, nil
		}
	}

	return -1, errors.New("Pet não encontrado")
}

func (ps *PetRepositorioJson) BuscarPorDonoId(id string) []models.Pet {
	pets := []models.Pet{}

	lista, _ := ps.Lista()

	for _, pet := range lista {
		if pet.Id == id {
			pets = append(pets, pet)
		}
	}

	return pets
}

func (ps *PetRepositorioJson) buscarPorIdStruct(pets []models.Pet, id string) (*models.Pet, int) {
	for i, pet := range pets {
		if pet.Id == id {
			return &pet, i
		}
	}

	return nil, -1
}

func (ps *PetRepositorioJson) validaCampos(pet *models.Pet) error {
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
