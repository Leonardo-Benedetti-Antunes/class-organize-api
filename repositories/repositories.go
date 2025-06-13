package repositories

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/cristiantebaldi/class-organize-api/models"
)

// ProfessorRepository gerencia operações de banco de dados para professores
type ProfessorRepository struct {
	DB *sql.DB
}

// SalaRepository gerencia operações de banco de dados para salas
type SalaRepository struct {
	DB *sql.DB
}

// TurmaRepository gerencia operações de banco de dados para turmas
type TurmaRepository struct {
	DB *sql.DB
}

// AlocacaoRepository gerencia operações de banco de dados para alocações
type AlocacaoRepository struct {
	DB *sql.DB
}

// NewProfessorRepository cria um novo repositório de professores
func NewProfessorRepository(db *sql.DB) *ProfessorRepository {
	return &ProfessorRepository{DB: db}
}

// NewSalaRepository cria um novo repositório de salas
func NewSalaRepository(db *sql.DB) *SalaRepository {
	return &SalaRepository{DB: db}
}

// NewTurmaRepository cria um novo repositório de turmas
func NewTurmaRepository(db *sql.DB) *TurmaRepository {
	return &TurmaRepository{DB: db}
}

// NewAlocacaoRepository cria um novo repositório de alocações
func NewAlocacaoRepository(db *sql.DB) *AlocacaoRepository {
	return &AlocacaoRepository{DB: db}
}

// ===== Métodos do ProfessorRepository =====

// GetAll retorna todos os professores
func (r *ProfessorRepository) GetAll() ([]models.Professor, error) {
	rows, err := r.DB.Query("SELECT id, nome, email, formacao, disciplina FROM professores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var professores []models.Professor
	for rows.Next() {
		var p models.Professor
		err := rows.Scan(&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina)
		if err != nil {
			return nil, err
		}
		professores = append(professores, p)
	}

	return professores, nil
}

// GetByID retorna um professor pelo ID
func (r *ProfessorRepository) GetByID(id int) (models.Professor, error) {
	var p models.Professor
	err := r.DB.QueryRow("SELECT id, nome, email, formacao, disciplina FROM professores WHERE id = $1", id).Scan(
		&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
	)
	if err != nil {
		return models.Professor{}, err
	}
	return p, nil
}

// Create cria um novo professor
func (r *ProfessorRepository) Create(p models.Professor) (models.Professor, error) {
	query := `INSERT INTO professores (nome, email, formacao, disciplina) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.DB.QueryRow(query, p.Nome, p.Email, p.Formacao, p.Disciplina).Scan(&p.ID)
	if err != nil {
		return models.Professor{}, err
	}

	return p, nil
}

// Update atualiza um professor existente
func (r *ProfessorRepository) Update(p models.Professor) error {
	query := `UPDATE professores SET nome = $1, email = $2, formacao = $3, disciplina = $4 
			WHERE id = $5`

	_, err := r.DB.Exec(query, p.Nome, p.Email, p.Formacao, p.Disciplina, p.ID)
	return err
}

// Delete remove um professor pelo ID
func (r *ProfessorRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM professores WHERE id = $1", id)
	return err
}

// ===== Métodos do SalaRepository =====

// GetAll retorna todas as salas
func (r *SalaRepository) GetAll() ([]models.Sala, error) {
	rows, err := r.DB.Query("SELECT id, numero, capacidade, bloco, tipo FROM salas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salas []models.Sala
	for rows.Next() {
		var s models.Sala
		err := rows.Scan(&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo)
		if err != nil {
			return nil, err
		}
		salas = append(salas, s)
	}

	return salas, nil
}

// GetByID retorna uma sala pelo ID
func (r *SalaRepository) GetByID(id int) (models.Sala, error) {
	var s models.Sala
	err := r.DB.QueryRow("SELECT id, numero, capacidade, bloco, tipo FROM salas WHERE id = $1", id).Scan(
		&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
	)
	if err != nil {
		return models.Sala{}, err
	}
	return s, nil
}

// Create cria uma nova sala
func (r *SalaRepository) Create(s models.Sala) (models.Sala, error) {
	query := `INSERT INTO salas (numero, capacidade, bloco, tipo) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.DB.QueryRow(query, s.Numero, s.Capacidade, s.Bloco, s.Tipo).Scan(&s.ID)
	if err != nil {
		return models.Sala{}, err
	}

	return s, nil
}

// Update atualiza uma sala existente
func (r *SalaRepository) Update(s models.Sala) error {
	query := `UPDATE salas SET numero = $1, capacidade = $2, bloco = $3, tipo = $4 
			WHERE id = $5`

	_, err := r.DB.Exec(query, s.Numero, s.Capacidade, s.Bloco, s.Tipo, s.ID)
	return err
}

// Delete remove uma sala pelo ID
func (r *SalaRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM salas WHERE id = $1", id)
	return err
}

// ===== Métodos do TurmaRepository =====

// GetAll retorna todas as turmas
func (r *TurmaRepository) GetAll() ([]models.Turma, error) {
	rows, err := r.DB.Query("SELECT id, nome, curso, periodo, quant_alunos FROM turmas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turmas []models.Turma
	for rows.Next() {
		var t models.Turma
		err := rows.Scan(&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos)
		if err != nil {
			return nil, err
		}
		turmas = append(turmas, t)
	}

	return turmas, nil
}

// GetByID retorna uma turma pelo ID
func (r *TurmaRepository) GetByID(id int) (models.Turma, error) {
	var t models.Turma
	err := r.DB.QueryRow("SELECT id, nome, curso, periodo, quant_alunos FROM turmas WHERE id = $1", id).Scan(
		&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
	)
	if err != nil {
		return models.Turma{}, err
	}
	return t, nil
}

// Create cria uma nova turma
func (r *TurmaRepository) Create(t models.Turma) (models.Turma, error) {
	query := `INSERT INTO turmas (nome, curso, periodo, quant_alunos) 
			VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.DB.QueryRow(query, t.Nome, t.Curso, t.Periodo, t.QuantAlunos).Scan(&t.ID)
	if err != nil {
		return models.Turma{}, err
	}

	return t, nil
}

// Update atualiza uma turma existente
func (r *TurmaRepository) Update(t models.Turma) error {
	query := `UPDATE turmas SET nome = $1, curso = $2, periodo = $3, quant_alunos = $4 
			WHERE id = $5`

	_, err := r.DB.Exec(query, t.Nome, t.Curso, t.Periodo, t.QuantAlunos, t.ID)
	return err
}

// Delete remove uma turma pelo ID
func (r *TurmaRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM turmas WHERE id = $1", id)
	return err
}

// ===== Métodos do AlocacaoRepository =====

// GetAll retorna todas as alocações com detalhes
func (r *AlocacaoRepository) GetAll() ([]models.Alocacao, error) {
	query := `
		SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id
	`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alocacoes []models.Alocacao
	for rows.Next() {
		var a models.Alocacao
		var p models.Professor
		var s models.Sala
		var t models.Turma

		err := rows.Scan(
			&a.ID, &a.ProfessorID, &a.SalaID, &a.TurmaID, &a.DiaSemana, &a.HorarioInicio, &a.HorarioFim,
			&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
			&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
			&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
		)
		if err != nil {
			return nil, err
		}

		a.Professor = p
		a.Sala = s
		a.Turma = t
		alocacoes = append(alocacoes, a)
	}

	return alocacoes, nil
}

// GetByID retorna uma alocação pelo ID com detalhes
func (r *AlocacaoRepository) GetByID(id int) (models.Alocacao, error) {
	query := `
		SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id
		WHERE a.id = $1
	`

	var a models.Alocacao
	var p models.Professor
	var s models.Sala
	var t models.Turma

	err := r.DB.QueryRow(query, id).Scan(
		&a.ID, &a.ProfessorID, &a.SalaID, &a.TurmaID, &a.DiaSemana, &a.HorarioInicio, &a.HorarioFim,
		&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
		&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
		&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
	)
	if err != nil {
		return models.Alocacao{}, err
	}

	a.Professor = p
	a.Sala = s
	a.Turma = t

	return a, nil
}

// Create cria uma nova alocação
func (r *AlocacaoRepository) Create(a models.Alocacao) (models.Alocacao, error) {
	// Verificar se a sala está disponível no horário solicitado
	var count int
	query := `
		SELECT COUNT(*) FROM alocacoes 
		WHERE sala_id = $1 AND dia_semana = $2 AND 
		((horario_inicio <= $3 AND horario_fim > $3) OR 
		(horario_inicio < $4 AND horario_fim >= $4) OR
		(horario_inicio >= $3 AND horario_fim <= $4))
	`

	err := r.DB.QueryRow(query, a.SalaID, a.DiaSemana, a.HorarioInicio, a.HorarioFim).Scan(&count)
	if err != nil {
		return models.Alocacao{}, err
	}

	if count > 0 {
		return models.Alocacao{}, fmt.Errorf("sala já está alocada neste horário")
	}

	// Inserir a alocação
	insertQuery := `
		INSERT INTO alocacoes (professor_id, sala_id, turma_id, dia_semana, horario_inicio, horario_fim) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`

	err = r.DB.QueryRow(insertQuery, a.ProfessorID, a.SalaID, a.TurmaID, a.DiaSemana, a.HorarioInicio, a.HorarioFim).Scan(&a.ID)
	if err != nil {
		return models.Alocacao{}, err
	}

	// Buscar a alocação completa com os detalhes
	return r.GetByID(a.ID)
}

// Update atualiza uma alocação existente
func (r *AlocacaoRepository) Update(a models.Alocacao) error {
	// Verificar se a sala está disponível no horário solicitado (excluindo a própria alocação)
	var count int
	query := `
		SELECT COUNT(*) FROM alocacoes 
		WHERE sala_id = $1 AND dia_semana = $2 AND 
		((horario_inicio <= $3 AND horario_fim > $3) OR 
		(horario_inicio < $4 AND horario_fim >= $4) OR
		(horario_inicio >= $3 AND horario_fim <= $4)) AND
		id != $5
	`

	err := r.DB.QueryRow(query, a.SalaID, a.DiaSemana, a.HorarioInicio, a.HorarioFim, a.ID).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("sala já está alocada neste horário")
	}

	// Atualizar a alocação
	updateQuery := `
		UPDATE alocacoes SET 
		professor_id = $1, sala_id = $2, turma_id = $3, 
		dia_semana = $4, horario_inicio = $5, horario_fim = $6 
		WHERE id = $7
	`

	_, err = r.DB.Exec(updateQuery, a.ProfessorID, a.SalaID, a.TurmaID, a.DiaSemana, a.HorarioInicio, a.HorarioFim, a.ID)
	return err
}

// Delete remove uma alocação pelo ID
func (r *AlocacaoRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM alocacoes WHERE id = $1", id)
	return err
}

// GetBySalaID retorna todas as alocações de uma sala específica
func (r *AlocacaoRepository) GetBySalaID(salaID int) ([]models.Alocacao, error) {
	query := `
		SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id
		WHERE a.sala_id = $1
	`

	rows, err := r.DB.Query(query, salaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alocacoes []models.Alocacao
	for rows.Next() {
		var a models.Alocacao
		var p models.Professor
		var s models.Sala
		var t models.Turma

		err := rows.Scan(
			&a.ID, &a.ProfessorID, &a.SalaID, &a.TurmaID, &a.DiaSemana, &a.HorarioInicio, &a.HorarioFim,
			&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
			&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
			&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
		)
		if err != nil {
			return nil, err
		}

		a.Professor = p
		a.Sala = s
		a.Turma = t
		alocacoes = append(alocacoes, a)
	}

	return alocacoes, nil
}

// GetByProfessorID retorna todas as alocações de um professor específico
func (r *AlocacaoRepository) GetByProfessorID(professorID int) ([]models.Alocacao, error) {
	query := `
		SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id
		WHERE a.professor_id = $1
	`

	rows, err := r.DB.Query(query, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alocacoes []models.Alocacao
	for rows.Next() {
		var a models.Alocacao
		var p models.Professor
		var s models.Sala
		var t models.Turma

		err := rows.Scan(
			&a.ID, &a.ProfessorID, &a.SalaID, &a.TurmaID, &a.DiaSemana, &a.HorarioInicio, &a.HorarioFim,
			&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
			&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
			&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
		)
		if err != nil {
			return nil, err
		}

		a.Professor = p
		a.Sala = s
		a.Turma = t
		alocacoes = append(alocacoes, a)
	}

	return alocacoes, nil
}

// GetByTurmaID retorna todas as alocações de uma turma específica
func (r *AlocacaoRepository) GetByTurmaID(turmaID int) ([]models.Alocacao, error) {
	query := `
		SELECT 
			a.id, a.professor_id, a.sala_id, a.turma_id, a.dia_semana, a.horario_inicio, a.horario_fim,
			p.id, p.nome, p.email, p.formacao, p.disciplina,
			s.id, s.numero, s.capacidade, s.bloco, s.tipo,
			t.id, t.nome, t.curso, t.periodo, t.quant_alunos
		FROM alocacoes a
		JOIN professores p ON a.professor_id = p.id
		JOIN salas s ON a.sala_id = s.id
		JOIN turmas t ON a.turma_id = t.id
		WHERE a.turma_id = $1
	`

	rows, err := r.DB.Query(query, turmaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alocacoes []models.Alocacao
	for rows.Next() {
		var a models.Alocacao
		var p models.Professor
		var s models.Sala
		var t models.Turma

		err := rows.Scan(
			&a.ID, &a.ProfessorID, &a.SalaID, &a.TurmaID, &a.DiaSemana, &a.HorarioInicio, &a.HorarioFim,
			&p.ID, &p.Nome, &p.Email, &p.Formacao, &p.Disciplina,
			&s.ID, &s.Numero, &s.Capacidade, &s.Bloco, &s.Tipo,
			&t.ID, &t.Nome, &t.Curso, &t.Periodo, &t.QuantAlunos,
		)
		if err != nil {
			return nil, err
		}

		a.Professor = p
		a.Sala = s
		a.Turma = t
		alocacoes = append(alocacoes, a)
	}

	return alocacoes, nil
}

// OrganizarAlocacoesAutomaticas organiza alocações automaticamente para um dia e horário específicos
func (r *AlocacaoRepository) OrganizarAlocacoesAutomaticas(diaSemana, horarioInicio, horarioFim string) ([]models.Alocacao, error) {
	// Inicializar o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// 1. Obter todos os professores disponíveis
	professores, err := r.getProfessoresDisponiveis(diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter professores disponíveis: %v", err)
	}

	// 2. Obter todas as salas disponíveis
	salas, err := r.getSalasDisponiveis(diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter salas disponíveis: %v", err)
	}

	// 3. Obter todas as turmas disponíveis
	turmas, err := r.getTurmasDisponiveis(diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, fmt.Errorf("erro ao obter turmas disponíveis: %v", err)
	}

	// 4. Verificar se há recursos suficientes para fazer alocações
	if len(professores) == 0 || len(salas) == 0 || len(turmas) == 0 {
		return nil, fmt.Errorf("não há recursos suficientes para fazer alocações")
	}

	// 5. Criar alocações automaticamente
	var alocacoesCriadas []models.Alocacao

	// Usar o número mínimo entre professores, salas e turmas disponíveis
	numAlocacoes := min(len(professores), min(len(salas), len(turmas)))

	// Embaralhar os arrays de professores, salas e turmas para alocação aleatória
	rand.Shuffle(len(professores), func(i, j int) {
		professores[i], professores[j] = professores[j], professores[i]
	})

	rand.Shuffle(len(salas), func(i, j int) {
		salas[i], salas[j] = salas[j], salas[i]
	})

	rand.Shuffle(len(turmas), func(i, j int) {
		turmas[i], turmas[j] = turmas[j], turmas[i]
	})

	for i := 0; i < numAlocacoes; i++ {
		// Criar uma nova alocação
		novaAlocacao := models.Alocacao{
			ProfessorID:   professores[i].ID,
			SalaID:        salas[i].ID,
			TurmaID:       turmas[i].ID,
			DiaSemana:     diaSemana,
			HorarioInicio: horarioInicio,
			HorarioFim:    horarioFim,
			Professor:     professores[i],
			Sala:          salas[i],
			Turma:         turmas[i],
		}

		// Inserir a alocação no banco de dados
		alocacaoCriada, err := r.Create(novaAlocacao)
		if err != nil {
			// Continuar mesmo se houver erro em uma alocação específica
			continue
		}

		alocacoesCriadas = append(alocacoesCriadas, alocacaoCriada)
	}

	if len(alocacoesCriadas) == 0 {
		return nil, fmt.Errorf("não foi possível criar nenhuma alocação")
	}

	return alocacoesCriadas, nil
}

// getProfessoresDisponiveis retorna professores disponíveis em um determinado dia e horário
func (r *AlocacaoRepository) getProfessoresDisponiveis(diaSemana, horarioInicio, horarioFim string) ([]models.Professor, error) {
	// Obter todos os professores
	professorRepo := &ProfessorRepository{DB: r.DB}
	allProfessores, err := professorRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Consultar professores já alocados no horário especificado
	query := `
		SELECT DISTINCT professor_id FROM alocacoes 
		WHERE dia_semana = $1 AND 
		((horario_inicio <= $2 AND horario_fim > $2) OR 
		(horario_inicio < $3 AND horario_fim >= $3) OR
		(horario_inicio >= $2 AND horario_fim <= $3))
	`

	rows, err := r.DB.Query(query, diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar um mapa de IDs de professores já alocados
	alocadosIDs := make(map[int]bool)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		alocadosIDs[id] = true
	}

	// Filtrar professores disponíveis
	var professoresDisponiveis []models.Professor
	for _, p := range allProfessores {
		if !alocadosIDs[p.ID] {
			professoresDisponiveis = append(professoresDisponiveis, p)
		}
	}

	return professoresDisponiveis, nil
}

// getSalasDisponiveis retorna salas disponíveis em um determinado dia e horário
func (r *AlocacaoRepository) getSalasDisponiveis(diaSemana, horarioInicio, horarioFim string) ([]models.Sala, error) {
	// Obter todas as salas
	salaRepo := &SalaRepository{DB: r.DB}
	allSalas, err := salaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Consultar salas já alocadas no horário especificado
	query := `
		SELECT DISTINCT sala_id FROM alocacoes 
		WHERE dia_semana = $1 AND 
		((horario_inicio <= $2 AND horario_fim > $2) OR 
		(horario_inicio < $3 AND horario_fim >= $3) OR
		(horario_inicio >= $2 AND horario_fim <= $3))
	`

	rows, err := r.DB.Query(query, diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar um mapa de IDs de salas já alocadas
	alocadosIDs := make(map[int]bool)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		alocadosIDs[id] = true
	}

	// Filtrar salas disponíveis
	var salasDisponiveis []models.Sala
	for _, s := range allSalas {
		if !alocadosIDs[s.ID] {
			salasDisponiveis = append(salasDisponiveis, s)
		}
	}

	return salasDisponiveis, nil
}

// getTurmasDisponiveis retorna turmas disponíveis em um determinado dia e horário
func (r *AlocacaoRepository) getTurmasDisponiveis(diaSemana, horarioInicio, horarioFim string) ([]models.Turma, error) {
	// Obter todas as turmas
	turmaRepo := &TurmaRepository{DB: r.DB}
	allTurmas, err := turmaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Consultar turmas já alocadas no horário especificado
	query := `
		SELECT DISTINCT turma_id FROM alocacoes 
		WHERE dia_semana = $1 AND 
		((horario_inicio <= $2 AND horario_fim > $2) OR 
		(horario_inicio < $3 AND horario_fim >= $3) OR
		(horario_inicio >= $2 AND horario_fim <= $3))
	`

	rows, err := r.DB.Query(query, diaSemana, horarioInicio, horarioFim)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Criar um mapa de IDs de turmas já alocadas
	alocadosIDs := make(map[int]bool)
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		alocadosIDs[id] = true
	}

	// Filtrar turmas disponíveis
	var turmasDisponiveis []models.Turma
	for _, t := range allTurmas {
		if !alocadosIDs[t.ID] {
			turmasDisponiveis = append(turmasDisponiveis, t)
		}
	}

	return turmasDisponiveis, nil
}

// Função auxiliar para obter o mínimo entre dois inteiros
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
