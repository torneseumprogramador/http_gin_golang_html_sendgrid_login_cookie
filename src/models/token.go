package models

//lint:ignore U1000 reason: Used by ORM to specify table name
type Token struct {
	tableName struct{} `table:"tokens"`

	Id    string `field:"id"`
	Token string `field:"token"`
	Email string `field:"email"`
}
