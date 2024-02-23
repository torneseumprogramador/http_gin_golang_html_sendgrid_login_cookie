package controllers

import (
	"fmt"
	"html/template"
	"http_gin/src/database"
	"http_gin/src/enums"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func petRepo() *repositorios.PetRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.PetRepositorioMySql{DB: db}
}

// func petRepo() *repositorios.PetRepositorioJson {
// 	return &repositorios.PetRepositorioJson{}
// }

type PetsController struct{}

func (pc *PetsController) Index(c *gin.Context) {
	servico := servicos.NovoPetServico(petRepo())
	pets, _ := servico.ListaPetView()

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Pets",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/index.tmpl.html",
					map[string][]model_views.PetView{
						"pets": pets,
					},
				),
			),
		},
	)
}

func donos() []models.Dono {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donos, _ := servico.Repo.Lista()
	return donos
}

func (pc *PetsController) Novo(c *gin.Context) {
	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"pet":    models.Pet{},
						"titulo": "Registro de Pet",
						"action": "/pets/cadastrar",
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Cadastrar(c *gin.Context) {
	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))

	pet := models.Pet{
		Id:     "",
		Nome:   c.Request.FormValue("nome"),
		DonoId: c.Request.FormValue("dono_id"),
		Tipo:   enums.Tipo(tipoInt),
	}

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	erro := servico.Repo.Adicionar(pet)

	if erro == nil {
		c.Redirect(302, "/pets")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"pet":    pet,
						"titulo": "Registro de Pet",
						"action": "/pets/cadastrar",
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	servico.Repo.Excluir(id)
	c.Redirect(302, "/pets")
}

func (pc *PetsController) Editar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	pet, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/pets")
		return
	}

	if pet == nil {
		c.Redirect(302, "/pets")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   nil,
						"pet":    pet,
						"titulo": "Alterando um Pet",
						"action": "/pets/" + pet.Id + "/alterar",
						"donos":  donos(),
					},
				),
			),
		},
	)
}

func (pc *PetsController) Alterar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Pet](petRepo())
	pet, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/pets")
		return
	}

	if pet == nil {
		c.Redirect(302, "/pets")
		return
	}

	pet.Nome = c.Request.FormValue("nome")
	pet.DonoId = c.Request.FormValue("dono_id")

	tipoInt, _ := strconv.Atoi(c.Request.FormValue("tipo"))
	pet.Tipo = enums.Tipo(tipoInt)

	erroAlterar := servico.Repo.Alterar(*pet)

	if erroAlterar == nil {
		c.Redirect(302, "/pets")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando um Pet",
			"currentRoute": "pets",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/pets/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erroAlterar,
						"pet":    pet,
						"titulo": "Alterando um Pet",
						"action": "/pets/" + pet.Id + "/alterar",
						"donos":  donos(),
					},
				),
			),
		},
	)
}
