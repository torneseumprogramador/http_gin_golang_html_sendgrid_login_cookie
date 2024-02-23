package middlewares

import (
	"encoding/json"
	"fmt"
	"http_gin/src/libs"
	"http_gin/src/model_views"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthRequired é um middleware que verifica a presença de um cookie específico.
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tenta obter o cookie pelo nome. Substitua "adminInfo" pelo nome real do seu cookie.
		cookieValue, err := c.Cookie("adminInfo")
		if err != nil {
			// Cookie não encontrado, redireciona para a página de login.
			c.Redirect(http.StatusFound, "/login")
			c.Abort() // Aborta a cadeia de manipuladores.
			return
		}

		// Decodifica o valor do cookie
		tokenValue, _ := url.QueryUnescape(cookieValue)

		chave := libs.GetEnv("JWT_TOKEN", "desafio_go")

		// Decodifica o token JWT
		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o algoritmo de assinatura esperado é usado
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
			}

			return []byte(chave), nil
		})

		if err != nil {
			fmt.Println(err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		decodedValue := ""
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			decodedValue = claims["sub"].(string)
		} else {
			fmt.Println("Token JWT inválido")
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		decodedJSON, err := url.QueryUnescape(decodedValue)
		if err != nil {
			fmt.Println("Erro ao decodificar o valor do URL:", err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Supondo que o valor do cookie seja um JSON do objeto Administrador
		adm := model_views.AdmView{}
		if err := json.Unmarshal([]byte(decodedJSON), &adm); err != nil {
			// Handle error
			fmt.Println(err)
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		currentRoute := c.FullPath()
		if !adm.Super && strings.Contains(currentRoute, "administradores") {
			fmt.Println("===[Usuário sem acesso a esta rota]====")
			c.Redirect(302, "/pets")
			c.Abort()
			return
		}

		// Armazena o objeto administrador no contexto do Gin
		c.Set("admin", adm)

		// Cookie encontrado, continua para o próximo manipulador no encadeamento.
		c.Next()
	}
}
