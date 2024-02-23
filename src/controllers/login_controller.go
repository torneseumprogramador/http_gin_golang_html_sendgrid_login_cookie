package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"http_gin/src/database"
	"http_gin/src/helpers"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"http_gin/src/models"
	"http_gin/src/repositorios"
	"http_gin/src/servicos"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LoginController struct{}

func (hc *LoginController) Index(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Página de login",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/index.tmpl.html",
					nil,
				),
			),
		},
	)
}

func (hc *LoginController) Esqueci(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Esqueci senha",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/esqueci.tmpl.html",
					nil,
				),
			),
		},
	)
}

func (pc *LoginController) Recuperar(c *gin.Context) {
	var erroReturn error

	email := c.Request.FormValue("email")

	if email != "" {
		servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
		params := make(map[string]string)
		params["email"] = email

		adms, erro := servico.Repo.Where(params)

		if erro != nil {
			erroReturn = fmt.Errorf(erro.Error())
		} else {
			if len(adms) == 0 {
				erroReturn = fmt.Errorf("email não encontrado")
			} else {
				admDb := adms[0]
				erro = helpers.EnviarEmailRecuperarSenha(admDb)
				if erro == nil {
					c.HTML(
						http.StatusOK,
						"main.tmpl.html",
						gin.H{
							"title":        "Esqueci senha",
							"currentRoute": "login",
							"content": template.HTML(
								libs.Render(
									"src/templates/pages/Login/esqueci.tmpl.html",
									map[string]interface{}{
										"sucesso": "Email enviado para " + email + ", confira na caixa de entrada",
									},
								),
							),
						},
					)

					return
				} else {
					erroReturn = erro
				}
			}
		}
	} else {
		erroReturn = fmt.Errorf("email é obrigatório")
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Esqueci senha",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/esqueci.tmpl.html",
					map[string]interface{}{
						"erro": erroReturn,
					},
				),
			),
		},
	)
}

func (pc *LoginController) RecadastrarSenha(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		db, _ := database.GetDB()
		tokenRepo := repositorios.GenericoRepositorioMySql[models.Token]{DB: db}
		params := make(map[string]string)
		params["token"] = token
		tokens, _ := tokenRepo.Where(params)
		if len(tokens) > 0 {
			c.HTML(
				http.StatusOK,
				"main.tmpl.html",
				gin.H{
					"title":        "Registrar senha",
					"currentRoute": "login",
					"content": template.HTML(
						libs.Render(
							"src/templates/pages/Login/registrar_senha.tmpl.html",
							map[string]interface{}{
								"token": token,
							},
						),
					),
				},
			)
			return
		}
	}

	c.Redirect(302, "/login")
}

func (pc *LoginController) RegistrarSenha(c *gin.Context) {
	token := c.Request.FormValue("token")
	if token != "" {

		db, _ := database.GetDB()
		tokenRepo := repositorios.GenericoRepositorioMySql[models.Token]{DB: db}
		params := make(map[string]string)
		params["token"] = token
		tokens, _ := tokenRepo.Where(params)
		if len(tokens) > 0 {
			tokenDb := tokens[0]
			var erro error

			if c.Request.FormValue("senha") != "" && c.Request.FormValue("senha") == c.Request.FormValue("csenha") {
				servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
				params := make(map[string]string)
				params["email"] = tokenDb.Email
				adms, _ := servico.Repo.Where(params)
				if len(adms) > 0 {
					adm := adms[0]
					adm.Senha = c.Request.FormValue("senha")
					servico.Repo.Alterar(adm)

					tokenRepo.ApagarPorId(tokenDb.Id)
					registrarLogin(c, adms[0])
					return
				} else {
					erro = fmt.Errorf("email de recuperação inválido")
				}
			} else {
				erro = fmt.Errorf("senha não é igual a confirmação de senha")
			}

			c.HTML(
				http.StatusOK,
				"main.tmpl.html",
				gin.H{
					"title":        "Registrar senha",
					"currentRoute": "login",
					"content": template.HTML(
						libs.Render(
							"src/templates/pages/Login/registrar_senha.tmpl.html",
							map[string]interface{}{
								"erro":  erro,
								"token": token,
							},
						),
					),
				},
			)
			return
		}
	}
	c.Redirect(302, "/login")
}

func (hc *LoginController) Registrar(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registrar Usuário",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/registrar.tmpl.html",
					map[string]interface{}{
						"administrador": models.Administrador{},
					},
				),
			),
		},
	)
}

func (pc *LoginController) Cadastrar(c *gin.Context) {
	var erro error

	administrador := models.Administrador{
		Id:    "",
		Nome:  c.Request.FormValue("nome"),
		Email: c.Request.FormValue("email"),
		Senha: c.Request.FormValue("senha"),
	}

	if c.Request.FormValue("senha") != "" && c.Request.FormValue("senha") == c.Request.FormValue("csenha") {
		servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())
		erro = servico.Repo.Adicionar(administrador)

		if erro == nil {
			registrarLogin(c, administrador)
			return
		}
	} else {
		erro = fmt.Errorf("senha não é igual a confirmação de senha")
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Registrar Usuário",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/registrar.tmpl.html",
					map[string]interface{}{
						"erro":          erro,
						"administrador": administrador,
					},
				),
			),
		},
	)
}

func (hc *LoginController) Sair(c *gin.Context) {
	c.SetCookie("adminInfo", "", -1, "/", "", false, true)
	c.Redirect(302, "/login")
}

func (hc *LoginController) Login(c *gin.Context) {
	servico := servicos.NovoCrudServico[models.Administrador](admRepositorio())

	credenciais := make(map[string]string)
	credenciais["email"] = c.Request.FormValue("email")

	adms, erro := servico.Repo.Where(credenciais)

	if len(adms) > 0 && libs.CryptoEq(c.Request.FormValue("senha"), adms[0].Senha) {
		adm := adms[0]
		registrarLogin(c, adm)
		return
	}

	if erro == nil {
		erro = fmt.Errorf("Login ou senha inválido")
	}

	c.HTML(
		http.StatusOK,
		"main.tmpl.html",
		gin.H{
			"title":        "Página de login",
			"currentRoute": "login",
			"content": template.HTML(
				libs.Render(
					"src/templates/pages/Login/index.tmpl.html",
					map[string]interface{}{
						"erro": erro,
					},
				),
			),
		},
	)
}

func registrarLogin(c *gin.Context, adm models.Administrador) {
	admView := model_views.AdmView{
		Id:    adm.Id,
		Email: adm.Email,
		Nome:  adm.Nome,
		Super: adm.Super,
	}

	cookieValueBytes, err := json.Marshal(admView)
	if err != nil {
		c.Redirect(302, "/login")
		return
	}

	cookieValue := string(cookieValueBytes)
	encodedValue := url.QueryEscape(cookieValue)

	tempoExpiracao := time.Now().Add(time.Hour * 1)

	token := jwt.New(jwt.SigningMethodHS256)

	// Define claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = encodedValue
	claims["exp"] = tempoExpiracao.Unix()

	chave := libs.GetEnv("JWT_TOKEN", "desafio_go")
	tokenString, err := token.SignedString([]byte(chave))
	if err != nil {
		c.Redirect(302, "/login")
		return
	}

	duracao := int(tempoExpiracao.Unix() - time.Now().Unix())

	c.SetCookie("adminInfo", tokenString, duracao, "/", "", false, true)

	c.Redirect(302, "/pets")
}
