package model_views

import (
	"http_gin/src/enums"
)

type PetView struct {
	Id     string
	Nome   string
	DonoId string
	Dono   string
	Tipo   enums.Tipo
}
