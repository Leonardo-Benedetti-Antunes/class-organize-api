package models

import (
	"database/sql"
	"fmt"
	"log"
)

// Professor representa um professor no sistema
type Professor struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Email      string `json:"email"`
	Formacao   string `json:"formacao"`
	Disciplina string `json:"disciplina"`
}

// Sala representa uma sala de aula no sistema
type Sala struct {
	ID         int    `json:"id"`
	Numero     string `json:"numero"`
	Capacidade int    `json:"capacidade"`
	Bloco      string `json:"bloco"`
	Tipo       string `json:"tipo"` // Laboratório, Sala comum, etc.
}

// Turma representa uma turma no sistema
type Turma struct {
	ID          int    `json:"id"`
	Nome        string `json:"nome"`
	Curso       string `json:"curso"`
	Periodo     string `json:"periodo"`
	QuantAlunos int    `json:"quant_alunos"`
}

// Alocacao representa a associação entre professor, sala e turma
type Alocacao struct {
	ID            int       `json:"id"`
	ProfessorID   int       `json:"professor_id"`
	SalaID        int       `json:"sala_id"`
	TurmaID       int       `json:"turma_id"`
	DiaSemana     string    `json:"dia_semana"`
	HorarioInicio string    `json:"horario_inicio"`
	HorarioFim    string    `json:"horario_fim"`
	Professor     Professor `json:"professor,omitempty"`
	Sala          Sala      `json:"sala,omitempty"`
	Turma         Turma     `json:"turma,omitempty"`
}

// MigrateTables cria as tabelas no banco de dados se não existirem
func MigrateTables(db *sql.DB) {
	// Criar tabela de professores
	createProfessorTable := `
	CREATE TABLE IF NOT EXISTS professores (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		formacao VARCHAR(100),
		disciplina VARCHAR(100) NOT NULL
	);
	`
	_, err := db.Exec(createProfessorTable)
	if err != nil {
		log.Fatalf("Erro ao criar tabela de professores: %v", err)
	}

	// Criar tabela de salas
	createSalaTable := `
	CREATE TABLE IF NOT EXISTS salas (
		id SERIAL PRIMARY KEY,
		numero VARCHAR(20) NOT NULL,
		capacidade INT NOT NULL,
		bloco VARCHAR(50),
		tipo VARCHAR(50)
	);
	`
	_, err = db.Exec(createSalaTable)
	if err != nil {
		log.Fatalf("Erro ao criar tabela de salas: %v", err)
	}

	// Criar tabela de turmas
	createTurmaTable := `
	CREATE TABLE IF NOT EXISTS turmas (
		id SERIAL PRIMARY KEY,
		nome VARCHAR(100) NOT NULL,
		curso VARCHAR(100) NOT NULL,
		periodo VARCHAR(50),
		quant_alunos INT
	);
	`
	_, err = db.Exec(createTurmaTable)
	if err != nil {
		log.Fatalf("Erro ao criar tabela de turmas: %v", err)
	}

	// Criar tabela de alocações
	createAlocacaoTable := `
	CREATE TABLE IF NOT EXISTS alocacoes (
		id SERIAL PRIMARY KEY,
		professor_id INT REFERENCES professores(id),
		sala_id INT REFERENCES salas(id),
		turma_id INT REFERENCES turmas(id),
		dia_semana VARCHAR(20) NOT NULL,
		horario_inicio VARCHAR(10) NOT NULL,
		horario_fim VARCHAR(10) NOT NULL,
		CONSTRAINT unique_alocacao UNIQUE(sala_id, dia_semana, horario_inicio)
	);
	`
	_, err = db.Exec(createAlocacaoTable)
	if err != nil {
		log.Fatalf("Erro ao criar tabela de alocações: %v", err)
	}

	fmt.Println("Tabelas criadas com sucesso")
}
