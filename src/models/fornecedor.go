package models

//lint:ignore U1000 reason: Used by ORM to specify table name
type Fornecedor struct {
	tableName struct{} `table:"fornecedores"`

	Id    string `field:"id"`
	Nome  string `field:"nome"`
	Email string `field:"email"`
}
