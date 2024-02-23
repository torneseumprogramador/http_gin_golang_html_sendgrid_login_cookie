package helpers

import (
	"http_gin/src/database"
	"http_gin/src/libs"
	"http_gin/src/models"
	"http_gin/src/repositorios"

	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func EnviarEmailRecuperarSenha(adm models.Administrador) error {
	link := "http://localhost:5000/recadastrar-senha?token=" + gerarToken(adm)

	from := mail.NewEmail("Suporte", "suporte@torneseumprogramador.com.br")
	subject := "Recuperar senha"
	to := mail.NewEmail(adm.Nome, adm.Email)
	plainTextContent := "Olá " + adm.Nome + " segue o link para recadastrar sua senha " + link
	htmlContent := libs.Render(
		"src/templates/email/recuperar_senha.tmpl.html",
		map[string]interface{}{
			"link": link,
		},
	)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	apiKey := libs.GetEnv("API_KEY_SENDGRID", "SEU TOKEN AQUI")
	client := sendgrid.NewSendClient(apiKey)
	_, err := client.Send(message)

	if err != nil {
		return err
	}

	return nil
}

func gerarToken(adm models.Administrador) string {
	token := uuid.New().String()
	db, _ := database.GetDB()

	tokenRepo := repositorios.GenericoRepositorioMySql[models.Token]{DB: db}
	tokenRepo.Adicionar(models.Token{
		Token: token,
		Email: adm.Email,
	})

	return token
}
