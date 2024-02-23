package models

//lint:ignore U1000 reason: Used by ORM to specify table name
type Dono struct {
	tableName struct{} `table:"donos"`

	Id       string `field:"id"`
	Nome     string `field:"nome"`
	Telefone string `field:"telefone"`
}
