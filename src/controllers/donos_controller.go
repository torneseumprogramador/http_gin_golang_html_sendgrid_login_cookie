package controllers

import (
	"fmt"
	"html/template"
	"http_gin/src/database"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func donoRepo() *repositorios.DonoRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.DonoRepositorioMySql{DB: db}
}

// func donoRepo() *repositorios.DonoRepositorioJson {
// 	return &repositorios.DonoRepositorioJson{}
// }

type DonosController struct{}

func (pc *DonosController) Index(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	donos, _ := servico.Repo.Lista()

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Donos",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/index.tmpl.html",
					map[string][]models.Dono{
						"donos": donos,
					},
				),
			),
		},
	)
}

func (pc *DonosController) Novo(c *gin.Context) {
	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"dono":   models.Dono{},
						"titulo": "Registro de Dono",
						"action": "/donos/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Cadastrar(c *gin.Context) {
	dono := models.Dono{
		Id:       "",
		Nome:     c.Request.FormValue("nome"),
		Telefone: c.Request.FormValue("telefone"),
	}

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	erro := servico.Repo.Adicionar(dono)

	if erro == nil {
		c.Redirect(302, "/donos")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erro,
						"dono":   dono,
						"titulo": "Registro de Dono",
						"action": "/donos/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	servico.Repo.Excluir(id)
	c.Redirect(302, "/donos")
}

func (pc *DonosController) Editar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	dono, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/donos")
		return
	}

	if dono == nil {
		c.Redirect(302, "/donos")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   nil,
						"dono":   dono,
						"titulo": "Alterando um Dono",
						"action": "/donos/" + dono.Id + "/alterar",
					},
				),
			),
		},
	)
}

func (pc *DonosController) Alterar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Dono](donoRepo())
	dono, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/donos")
		return
	}

	if dono == nil {
		c.Redirect(302, "/donos")
		return
	}

	dono.Nome = c.Request.FormValue("nome")
	dono.Telefone = c.Request.FormValue("telefone")

	erroAlterar := servico.Repo.Alterar(*dono)

	if erroAlterar == nil {
		c.Redirect(302, "/donos")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando um Dono",
			"currentRoute": "donos",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/donos/salvar.tmpl.html",
					map[string]interface{}{
						"erro":   erroAlterar,
						"dono":   dono,
						"titulo": "Alterando um Dono",
						"action": "/donos/" + dono.Id + "/alterar",
					},
				),
			),
		},
	)
}
