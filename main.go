package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/cristiantebaldi/class-organize-api/config"
	"github.com/cristiantebaldi/class-organize-api/controllers"
	"github.com/cristiantebaldi/class-organize-api/models"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Inicializar conexão com o banco de dados
	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Criar tabelas se não existirem
	models.MigrateTables(db)

	// Inicializar o router
	r := mux.NewRouter()

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Configurar rotas
	controllers.SetupRoutes(r, db)

	// Iniciar o servidor
	port := "8080"
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, c.Handler(r)))
}