package servicos

import (
	"http_gin/src/model_views"
	"http_gin/src/repositorios"
)

type PetServico struct {
	Repo *repositorios.PetRepositorioMySql
}

func NovoPetServico(repo *repositorios.PetRepositorioMySql) *PetServico {
	return &PetServico{
		Repo: repo,
	}
}

func (ps *PetServico) ListaPetView() ([]model_views.PetView, error) {
	return ps.Repo.ListaPetView()
}
