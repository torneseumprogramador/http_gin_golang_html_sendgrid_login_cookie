Como utilizar o repo generico

```go

package main

import (
	"fmt"
	"http_gin/src/database"
	"http_gin/src/models"
	"http_gin/src/repositorios"
)

func main() {
	db, _ := database.GetDB()

	// adminGenericoRepo := repositorios.GenericoRepositorioMySql[models.Administrador]{
	// 	DB:    db,
	// 	Table: "administradores",
	// }

	// adms, _ := adminGenericoRepo.Lista()

	// for _, adm := range adms {
	// 	fmt.Println("--------------------------")
	// 	fmt.Printf("ID: %v\n", adm.Id)
	// 	fmt.Printf("Nome: %v\n", adm.Nome)
	// 	fmt.Printf("Email: %v\n", adm.Email)
	// 	fmt.Printf("Senha: %v\n", adm.Senha)
	// }

	// fmt.Println("------------------")

	// donoGenericoRepo := repositorios.GenericoRepositorioMySql[models.Dono]{
	// 	DB:    db,
	// 	Table: "donos",
	// }

	// donos, _ := donoGenericoRepo.Lista()

	// for _, dono := range donos {
	// 	fmt.Println("--------------------------")
	// 	fmt.Printf("ID: %v\n", dono.Id)
	// 	fmt.Printf("Nome: %v\n", dono.Nome)
	// 	fmt.Printf("Telefone: %v\n", dono.Telefone)
	// }

	// fmt.Println("------------------")

	// petGenericoRepo := repositorios.GenericoRepositorioMySql[models.Pet]{
	// 	DB:    db,
	// 	Table: "pets",
	// }

	// pets, _ := petGenericoRepo.Lista()

	// for _, pet := range pets {
	// 	fmt.Println("--------------------------")
	// 	fmt.Printf("ID: %v\n", pet.Id)
	// 	fmt.Printf("Nome: %v\n", pet.Nome)
	// 	fmt.Printf("DonoId: %v\n", pet.DonoId)
	// 	fmt.Printf("Tipo: %v\n", pet.Tipo.String())
	// }

	fmt.Println("------------------")

	fornecedorGenericoRepo := repositorios.GenericoRepositorioMySql[models.Fornecedor]{DB: db}

	// fornecedorInsert := models.Fornecedor{}
	// fornecedorInsert.Nome = "Um novo adicionado"
	// fornecedorInsert.Email = "novo@teste.com"
	// erro := fornecedorGenericoRepo.Adicionar(fornecedorInsert)

	// if erro != nil {
	// 	fmt.Println(erro)
	// }

	// fornecedorAlterar, _ := fornecedorGenericoRepo.BuscaPorId("2")
	// fornecedorAlterar.Nome = "Empresa SA LTDA"
	// fornecedorGenericoRepo.Alterar(*fornecedorAlterar)

	// fornecedorGenericoRepo.ApagarPorId("ssds2123222")

	fornecedores, erro := fornecedorGenericoRepo.Lista()

	if erro != nil {
		fmt.Println(erro)
	}

	for _, fornecedor := range fornecedores {
		fmt.Println("--------------------------")
		fmt.Printf("ID: %v\n", fornecedor.Id)
		fmt.Printf("Nome: %v\n", fornecedor.Nome)
		fmt.Printf("Email: %v\n", fornecedor.Email)
	}
}

```