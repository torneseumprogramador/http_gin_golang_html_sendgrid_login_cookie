// config/cnn.go

package database

import (
	"database/sql"
	"fmt"
	"http_gin/src/libs"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetDB retorna uma conexão com o banco de dados
func GetDB() (*sql.DB, error) {
	dbUser := libs.GetEnv("DB_USER", "root")
	dbPassword := libs.GetEnv("DB_PASSWORD", "root")
	dbHost := libs.GetEnv("DB_HOST", "127.0.0.1:3306")
	dbName := libs.GetEnv("DB_NAME", "desafio_go")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Erro ao abrir a conexão com o banco de dados: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	return db, nil
}
