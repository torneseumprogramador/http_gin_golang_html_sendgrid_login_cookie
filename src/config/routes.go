package config

import (
	"http_gin/src/controllers"
	"http_gin/src/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	homeController := controllers.HomeController{}
	r.GET("/", homeController.Index)

	loginController := controllers.LoginController{}
	r.POST("/registrar-senha", loginController.RegistrarSenha)
	r.GET("/recadastrar-senha", loginController.RecadastrarSenha)
	r.GET("/esqueci", loginController.Esqueci)
	r.POST("/esqueci", loginController.Recuperar)
	r.GET("/registrar", loginController.Registrar)
	r.POST("/registrar", loginController.Cadastrar)
	r.GET("/login", loginController.Index)
	r.POST("/login", loginController.Login)
	r.GET("/sair", loginController.Sair)

	protectedRoutes := r.Group("/").Use(middlewares.AuthRequired())
	{
		petsController := controllers.PetsController{}
		protectedRoutes.GET("/pets", petsController.Index)
		protectedRoutes.GET("/pets/novo", petsController.Novo)
		protectedRoutes.POST("/pets/cadastrar", petsController.Cadastrar)
		protectedRoutes.GET("/pets/:id/excluir", petsController.Excluir)
		protectedRoutes.GET("/pets/:id/editar", petsController.Editar)
		protectedRoutes.POST("/pets/:id/alterar", petsController.Alterar)

		donosController := controllers.DonosController{}
		protectedRoutes.GET("/donos", donosController.Index)
		protectedRoutes.GET("/donos/novo", donosController.Novo)
		protectedRoutes.POST("/donos/cadastrar", donosController.Cadastrar)
		protectedRoutes.GET("/donos/:id/excluir", donosController.Excluir)
		protectedRoutes.GET("/donos/:id/editar", donosController.Editar)
		protectedRoutes.POST("/donos/:id/alterar", donosController.Alterar)

		administradoresController := controllers.AdministradoresController{}
		protectedRoutes.GET("/administradores", administradoresController.Index)
		protectedRoutes.GET("/administradores/novo", administradoresController.Novo)
		protectedRoutes.POST("/administradores/cadastrar", administradoresController.Cadastrar)
		protectedRoutes.GET("/administradores/:id/excluir", administradoresController.Excluir)
		protectedRoutes.GET("/administradores/:id/editar", administradoresController.Editar)
		protectedRoutes.POST("/administradores/:id/alterar", administradoresController.Alterar)
	}
}
