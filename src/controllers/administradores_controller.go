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

func admRepositorio() *repositorios.AdministradorRepositorioMySql {
	db, _ := database.GetDB()
	return &repositorios.AdministradorRepositorioMySql{DB: db}
}

// func admRepositorio() *repositorios.AdministradorRepositorioJson {
// 	return &repositorios.AdministradorRepositorioJson{}
// }

type AdministradoresController struct{}

func (pc *AdministradoresController) Index(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())

	adm, _ := c.Get("admin")
	administradores, _ := servico.Repo.Lista()
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Administradores",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/index.tmpl.html",
					map[string][]models.Administrador{
						"administradores": administradores,
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Novo(c *gin.Context) {
	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"administrador": models.Administrador{},
						"titulo":        "Registro de Administrador",
						"action":        "/administradores/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Cadastrar(c *gin.Context) {
	administrador := models.Administrador{
		Id:    "",
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
		Senha: c.Request.FormValue("senha"),
	}

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	erro := servico.Repo.Adicionar(administrador)

	if erro == nil {
		c.Redirect(302, "/administradores")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Registro de Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          erro,
						"administrador": administrador,
						"titulo":        "Registro de Administrador",
						"action":        "/administradores/cadastrar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Excluir(c *gin.Context) {
	id := c.Param("id")

	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	servico.Repo.Excluir(id)
	c.Redirect(302, "/administradores")
}

func (pc *AdministradoresController) Editar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	administrador, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/administradores")
		return
	}

	if administrador == nil {
		c.Redirect(302, "/administradores")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          nil,
						"administrador": administrador,
						"titulo":        "Alterando um Administrador",
						"action":        "/administradores/" + administrador.Id + "/alterar",
					},
				),
			),
		},
	)
}

func (pc *AdministradoresController) Alterar(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
	administrador, erro := servico.Repo.BuscarPorId(c.Param("id"))

	if erro != nil {
		fmt.Println("Erro ao executar instrução sql ", erro.Error())
		c.Redirect(302, "/administradores")
		return
	}

	if administrador == nil {
		c.Redirect(302, "/administradores")
		return
	}

	administrador.Nome = c.Request.FormValue("nome")
	administrador.Email = c.Request.FormValue("email")
	administrador.Senha = c.Request.FormValue("senha")

	erroAlterar := servico.Repo.Alterar(*administrador)

	if erroAlterar == nil {
		c.Redirect(302, "/administradores")
		return
	}

	adm, _ := c.Get("admin")
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"adm":          adm,
			"title":        "Alterando um Administrador",
			"currentRoute": "administradores",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/administradores/salvar.tmpl.html",
					map[string]interface{}{
						"erro":          erroAlterar,
						"administrador": administrador,
						"titulo":        "Alterando um Administrador",
						"action":        "/administradores/" + administrador.Id + "/alterar",
					},
				),
			),
		},
	)
}
