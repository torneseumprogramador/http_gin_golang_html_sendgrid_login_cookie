package servicos

type CrudServico[T any] struct {
	Repo CRUDServico[T]
}

func NovoCrudServico[T any](repo CRUDServico[T]) *CrudServico[T] {
	return &CrudServico[T]{
		Repo: repo,
	}
}
