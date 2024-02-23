package main

import (
	"encoding/json"
	"fmt"
	"http_gin/src/config"
	"http_gin/src/database"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func importar_adm_json(done chan bool) {
	db, _ := database.GetDB()

	// Abrindo o arquivo JSON
	arquivo, err := os.Open("db/administradores.json")
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}
	defer arquivo.Close()

	// Lendo o arquivo com io.ReadAll
	bytes, err := io.ReadAll(arquivo)
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	// Deserializando o JSON para uma slice de Administrador
	var adms []models.Administrador
	if err := json.Unmarshal(bytes, &adms); err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	// Iterando sobre a slice de administradores e inserindo no banco de dados
	fornecedorGenericoRepo := repositorios.AdministradorRepositorioMySql{DB: db}
	for _, adm := range adms {
		fornecedorGenericoRepo.Adicionar(adm)
	}

	done <- true // terminei a tarefa (escrevendo)
}

func importar_dono_json(done chan bool) {
	db, _ := database.GetDB()

	// Abrindo o arquivo JSON
	arquivo, err := os.Open("db/donos.json")
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}
	defer arquivo.Close()

	// Lendo o arquivo com io.ReadAll
	bytes, err := io.ReadAll(arquivo)
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	// Deserializando o JSON para uma slice de Dono
	var donos []models.Dono
	if err := json.Unmarshal(bytes, &donos); err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	fornecedorGenericoRepo := repositorios.GenericoRepositorioMySql[models.Dono]{DB: db}
	for _, dono := range donos {
		fornecedorGenericoRepo.Adicionar(dono)
	}

	done <- true // terminei a tarefa (escrevendo)
}

func importar_pet_json(done chan bool) {
	db, _ := database.GetDB()

	// Abrindo o arquivo JSON
	arquivo, err := os.Open("db/pets.json")
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}
	defer arquivo.Close()

	// Lendo o arquivo com io.ReadAll
	bytes, err := io.ReadAll(arquivo)
	if err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	// Deserializando o JSON para uma slice de Pet
	var pets []models.Pet
	if err := json.Unmarshal(bytes, &pets); err != nil {
		panic(err) // ou trate o erro de forma adequada
	}

	// Iterando sobre a slice de Pets e inserindo no banco de dados
	fornecedorGenericoRepo := repositorios.GenericoRepositorioMySql[models.Pet]{DB: db}
	for _, pet := range pets {
		fornecedorGenericoRepo.Adicionar(pet)
	}

	done <- true // terminei a tarefa (escrevendo)
}

func importarDadosJson() {
	canalAdm := make(chan bool)
	canalDono := make(chan bool)
	canalPet := make(chan bool)

	go importar_adm_json(canalAdm)
	go importar_dono_json(canalDono)

	<-canalAdm
	<-canalDono

	go importar_pet_json(canalPet) // Pet é dependente de dono

	<-canalPet

	fmt.Println("Programa terminou")
}

func startWebApp() {
	r := gin.Default()

	// Carrega os templates HTML
	r.LoadHTMLGlob("src/templates/**/*.tmpl.html")

	r.Static("/public/", "./public_assets")

	config.Routes(r)

	r.Run(":5000") // Por padrão, escuta na porta 5000
}

func main() {
	// importarDadosJson()

	startWebApp()
}
