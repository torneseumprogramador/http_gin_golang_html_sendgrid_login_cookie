package models

import "http_gin/src/enums"

//lint:ignore U1000 reason: Used by ORM to specify table name
type Pet struct {
	tableName struct{} `table:"pets"`

	Id     string     `field:"id"`
	Nome   string     `field:"nome"`
	DonoId string     `field:"donoId"`
	Tipo   enums.Tipo `field:"tipo"`
}
