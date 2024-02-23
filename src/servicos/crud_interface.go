package servicos

type CRUDServico[T any] interface {
	Lista() ([]T, error)
	BuscarPorId(id string) (*T, error)
	Adicionar(entidade T) error
	Alterar(entidade T) error
	Excluir(id string) error
	Where(filtro map[string]string) ([]T, error)
}
