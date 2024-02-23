# Bem-vindo ao Projeto Go Web com Renderização Server-Side!

Este é um projeto desenvolvido em Go (Golang) que oferece uma aplicação web com renderização server-side, autenticação por cookie e envio de e-mails utilizando SendGrid. Esta aplicação é perfeita para construir sites dinâmicos e interativos com toda a potência e desempenho do Go.

## Comunidade

Este projeto é uma iniciativa da comunidade "Torne-se um Programador". Junte-se a nós para aprender e colaborar em mais projetos emocionantes. Visite [torneseumprogramador.com.br](https://www.torneseumprogramador.com.br) para mais informações e confira o curso [Desafio Go Lang](https://www.torneseumprogramador.com.br/cursos/desafio_go_lang).

## Destaques do Projeto

- **Renderização Server-Side**: Utilizando templates HTML, o projeto proporciona uma experiência de usuário fluida e responsiva, com páginas renderizadas diretamente no servidor para uma rápida inicialização e navegação suave.

- **Autenticação por Cookie**: Implementamos uma robusta autenticação por cookie, garantindo segurança e controle de acesso aos recursos da aplicação.

- **Envio de E-mails com SendGrid**: Integrado com o serviço de e-mails SendGrid, a aplicação oferece recursos avançados de comunicação, permitindo enviar e-mails transacionais de forma confiável e escalável.

## Estrutura do Projeto

```
.
├── README.md
├── TODO.md
├── db
│   ├── administradores.json
│   ├── donos.json
│   ├── pets.json
│   └── query.sql
├── go.mod
├── go.sum
├── main.go
├── migrate.sh
├── mysql.sh
├── run.sh
└── src
    ├── config
    │   └── routes.go
    ├── controllers
    │   ├── administradores_controller.go
    │   ├── donos_controller.go
    │   ├── home_controller.go
    │   ├── login_controller.go
    │   └── pets_controller.go
    ├── database
    │   └── cnn.go
    ├── enums
    │   └── tipo.go
    ├── helpers
    │   └── emails.go
    ├── libs
    │   └── utls.go
    ├── middlewares
    │   └── auth.go
    ├── model_views
    │   ├── adm_view.go
    │   └── pet_view.go
    ├── models
    │   ├── administrador.go
    │   ├── dono.go
    │   ├── fornecedor.go
    │   ├── pet.go
    │   └── token.go
    ├── repositorios
    │   ├── administrador_repositorio_json.go
    │   ├── administrador_repositorio_mysql.go
    │   ├── dono_repositorio_json.go
    │   ├── dono_repositorio_mysql.go
    │   ├── generico_repositorio_mysql.go
    │   ├── pet_repositorio_json.go
    │   └── pet_repositorio_mysql.go
    ├── servicos
    │   ├── crud_interface.go
    │   ├── crud_servico.go
    │   └── pet_servico.go
    └── templates
        ├── email
        │   └── recuperar_senha.tmpl.html
        ├── layouts
        │   └── main.tmpl.html
        ├── pages
        │   ├── administradores
        │   │   ├── index.tmpl.html
        │   │   └── salvar.tmpl.html
        │   ├── donos
        │   │   ├── index.tmpl.html
        │   │   └── salvar.tmpl.html
        │   ├── home
        │   │   └── index.tmpl.html
        │   ├── home.tmpl.html
        │   ├── login
        │   │   ├── esqueci.tmpl.html
        │   │   ├── index.tmpl.html
        │   │   ├── registrar.tmpl.html
        │   │   └── registrar_senha.tmpl.html
        │   └── pets
        │       ├── index.tmpl.html
        │       └── salvar.tmpl.html
        └── partials
            ├── footer.tmpl.html
            └── header.tmpl.html
```

## Instruções de Uso

- **Configuração do Banco de Dados**: Utilize os scripts `migrate.sh` e `mysql.sh` para configurar o banco de dados e migrar os dados necessários.

- **Execução**: Utilize o script `run.sh` para iniciar o servidor web e começar a usar a aplicação.



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
