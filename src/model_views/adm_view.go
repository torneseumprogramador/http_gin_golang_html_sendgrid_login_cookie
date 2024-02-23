package model_views

type AdmView struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
	Super bool   `json:"super"`
}
