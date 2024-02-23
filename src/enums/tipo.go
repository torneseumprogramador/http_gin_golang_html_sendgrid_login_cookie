package enums

type Tipo int

// Definindo os valores do enum
const (
	Cachorro Tipo = iota // iota facilita a atribuição incremental de valores
	Gato
	Outros
)

func (t Tipo) String() string {
	switch t {
	case Cachorro:
		return "Cachorro"
	case Gato:
		return "Gato"
	case Outros:
		return "Outros"
	default:
		return "Desconhecido"
	}
}
