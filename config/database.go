package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// SetupDB configura e conecta ao banco de dados PostgreSQL
func SetupDB() (*sql.DB, error) {
	// Obter variáveis de ambiente para conexão com o banco de dados
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Construir string de conexão
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Abrir conexão com o banco de dados
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Verificar conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado com sucesso ao banco de dados")
	return db, nil
}
