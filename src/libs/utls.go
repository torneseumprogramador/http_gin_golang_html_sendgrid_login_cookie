package libs

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

func IsCrypto(senha string) bool {
	return len(senha) == 60
}

func Crypto(senha string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func CryptoEq(senhaUnCrypto, senhaHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaUnCrypto))
	return err == nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func Render(templateFile string, data interface{}) string {
	// Carrega o arquivo de template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatalf("Error parsing template files: %v", err)
	}

	// Cria um buffer para armazenar a saída do template
	var tmplBytes bytes.Buffer

	// Executa o template com os dados fornecidos e escreve a saída no buffer
	err = tmpl.Execute(&tmplBytes, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Retorna a string resultante
	return tmplBytes.String()
}
