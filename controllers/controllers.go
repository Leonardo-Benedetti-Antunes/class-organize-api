package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cristiantebaldi/class-organize-api/infra"
	"github.com/cristiantebaldi/class-organize-api/models"
	"github.com/cristiantebaldi/class-organize-api/repositories"

	"github.com/gorilla/mux"
)

// ProfessorController gerencia as requisições relacionadas a professores
type ProfessorController struct {
	Repo *repositories.ProfessorRepository
}

// SalaController gerencia as requisições relacionadas a salas
type SalaController struct {
	Repo *repositories.SalaRepository
}

// TurmaController gerencia as requisições relacionadas a turmas
type TurmaController struct {
	Repo *repositories.TurmaRepository
}

// AlocacaoController gerencia as requisições relacionadas a alocações
type AlocacaoController struct {
	Repo *repositories.AlocacaoRepository
}

// NewProfessorController cria um novo controlador de professores
func NewProfessorController(db *sql.DB) *ProfessorController {
	return &ProfessorController{
		Repo: repositories.NewProfessorRepository(db),
	}
}

// NewSalaController cria um novo controlador de salas
func NewSalaController(db *sql.DB) *SalaController {
	return &SalaController{
		Repo: repositories.NewSalaRepository(db),
	}
}

// NewTurmaController cria um novo controlador de turmas
func NewTurmaController(db *sql.DB) *TurmaController {
	return &TurmaController{
		Repo: repositories.NewTurmaRepository(db),
	}
}

// NewAlocacaoController cria um novo controlador de alocações
func NewAlocacaoController(db *sql.DB) *AlocacaoController {
	return &AlocacaoController{
		Repo: repositories.NewAlocacaoRepository(db),
	}
}

// SetupRoutes configura todas as rotas da API
func SetupRoutes(r *mux.Router, db *sql.DB) {
	// Inicializar controladores
	professorController := NewProfessorController(db)
	salaController := NewSalaController(db)
	turmaController := NewTurmaController(db)
	alocacaoController := NewAlocacaoController(db)

	// Rotas para professores
	r.HandleFunc("/api/professores", professorController.GetAllProfessores).Methods("GET")
	r.HandleFunc("/api/professores/{id}", professorController.GetProfessor).Methods("GET")
	r.HandleFunc("/api/professores", professorController.CreateProfessor).Methods("POST")
	r.HandleFunc("/api/professores/{id}", professorController.UpdateProfessor).Methods("PUT")
	r.HandleFunc("/api/professores/{id}", professorController.DeleteProfessor).Methods("DELETE")

	// Rotas para salas
	r.HandleFunc("/api/salas", salaController.GetAllSalas).Methods("GET")
	r.HandleFunc("/api/salas/{id}", salaController.GetSala).Methods("GET")
	r.HandleFunc("/api/salas", salaController.CreateSala).Methods("POST")
	r.HandleFunc("/api/salas/{id}", salaController.UpdateSala).Methods("PUT")
	r.HandleFunc("/api/salas/{id}", salaController.DeleteSala).Methods("DELETE")

	// Rotas para turmas
	r.HandleFunc("/api/turmas", turmaController.GetAllTurmas).Methods("GET")
	r.HandleFunc("/api/turmas/{id}", turmaController.GetTurma).Methods("GET")
	r.HandleFunc("/api/turmas", turmaController.CreateTurma).Methods("POST")
	r.HandleFunc("/api/turmas/{id}", turmaController.UpdateTurma).Methods("PUT")
	r.HandleFunc("/api/turmas/{id}", turmaController.DeleteTurma).Methods("DELETE")

	// Rotas para alocações
	r.HandleFunc("/api/alocacoes", alocacaoController.GetAllAlocacoes).Methods("GET")
	r.HandleFunc("/api/alocacoes/{id}", alocacaoController.GetAlocacao).Methods("GET")
	r.HandleFunc("/api/alocacoes", alocacaoController.CreateAlocacao).Methods("POST")
	r.HandleFunc("/api/alocacoes/{id}", alocacaoController.UpdateAlocacao).Methods("PUT")
	r.HandleFunc("/api/alocacoes/{id}", alocacaoController.DeleteAlocacao).Methods("DELETE")
	r.HandleFunc("/api/alocacoes/automatico", alocacaoController.OrganizarAlocacoesAutomaticas).Methods("POST")

	// Rotas especiais para alocações
	r.HandleFunc("/api/alocacoes/sala/{id}", alocacaoController.GetAlocacoesBySala).Methods("GET")
	r.HandleFunc("/api/alocacoes/professor/{id}", alocacaoController.GetAlocacoesByProfessor).Methods("GET")
	r.HandleFunc("/api/alocacoes/turma/{id}", alocacaoController.GetAlocacoesByTurma).Methods("GET")
}

// ===== Métodos do ProfessorController =====

// GetAllProfessores retorna todos os professores
func (c *ProfessorController) GetAllProfessores(w http.ResponseWriter, r *http.Request) {
	professores, err := c.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professores)
}

// GetProfessor retorna um professor pelo ID
func (c *ProfessorController) GetProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	professor, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Professor não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(professor)
}

// CreateProfessor cria um novo professor
func (c *ProfessorController) CreateProfessor(w http.ResponseWriter, r *http.Request) {
	var professor models.Professor
	err := json.NewDecoder(r.Body).Decode(&professor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	professor, err = c.Repo.Create(professor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(professor)
}

// UpdateProfessor atualiza um professor existente
func (c *ProfessorController) UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var professor models.Professor
	err = json.NewDecoder(r.Body).Decode(&professor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	professor.ID = id
	err = c.Repo.Update(professor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProfessor remove um professor pelo ID
func (c *ProfessorController) DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ===== Métodos do SalaController =====

// GetAllSalas retorna todas as salas
func (c *SalaController) GetAllSalas(w http.ResponseWriter, r *http.Request) {
	salas, err := c.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(salas)
}

// GetSala retorna uma sala pelo ID
func (c *SalaController) GetSala(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	sala, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Sala não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sala)
}

// CreateSala cria uma nova sala
func (c *SalaController) CreateSala(w http.ResponseWriter, r *http.Request) {
	var sala models.Sala
	err := json.NewDecoder(r.Body).Decode(&sala)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sala, err = c.Repo.Create(sala)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sala)
}

// UpdateSala atualiza uma sala existente
func (c *SalaController) UpdateSala(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var sala models.Sala
	err = json.NewDecoder(r.Body).Decode(&sala)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sala.ID = id
	err = c.Repo.Update(sala)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteSala remove uma sala pelo ID
func (c *SalaController) DeleteSala(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ===== Métodos do TurmaController =====

// GetAllTurmas retorna todas as turmas
func (c *TurmaController) GetAllTurmas(w http.ResponseWriter, r *http.Request) {
	turmas, err := c.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turmas)
}

// GetTurma retorna uma turma pelo ID
func (c *TurmaController) GetTurma(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	turma, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Turma não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(turma)
}

// CreateTurma cria uma nova turma
func (c *TurmaController) CreateTurma(w http.ResponseWriter, r *http.Request) {
	var turma models.Turma
	err := json.NewDecoder(r.Body).Decode(&turma)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	turma, err = c.Repo.Create(turma)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(turma)
}

// UpdateTurma atualiza uma turma existente
func (c *TurmaController) UpdateTurma(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var turma models.Turma
	err = json.NewDecoder(r.Body).Decode(&turma)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	turma.ID = id
	err = c.Repo.Update(turma)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteTurma remove uma turma pelo ID
func (c *TurmaController) DeleteTurma(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ===== Métodos do AlocacaoController =====

// GetAllAlocacoes retorna todas as alocações
func (c *AlocacaoController) GetAllAlocacoes(w http.ResponseWriter, r *http.Request) {
	alocacoes, err := c.Repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alocacoes)
}

// GetAlocacao retorna uma alocação pelo ID
func (c *AlocacaoController) GetAlocacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	alocacao, err := c.Repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Alocação não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alocacao)
}

// CreateAlocacao cria uma nova alocação
func (c *AlocacaoController) CreateAlocacao(w http.ResponseWriter, r *http.Request) {
	var alocacao models.Alocacao
	err := json.NewDecoder(r.Body).Decode(&alocacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alocacao, err = c.Repo.Create(alocacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go infra.SendEmailOnAlocacaoSuccess(alocacao)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alocacao)
}

// UpdateAlocacao atualiza uma alocação existente
func (c *AlocacaoController) UpdateAlocacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var alocacao models.Alocacao
	err = json.NewDecoder(r.Body).Decode(&alocacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	alocacao.ID = id
	err = c.Repo.Update(alocacao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAlocacao remove uma alocação pelo ID
func (c *AlocacaoController) DeleteAlocacao(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.Repo.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAlocacoesBySala retorna todas as alocações de uma sala específica
func (c *AlocacaoController) GetAlocacoesBySala(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	alocacoes, err := c.Repo.GetBySalaID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alocacoes)
}

// GetAlocacoesByProfessor retorna todas as alocações de um professor específico
func (c *AlocacaoController) GetAlocacoesByProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	alocacoes, err := c.Repo.GetByProfessorID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alocacoes)
}

// GetAlocacoesByTurma retorna todas as alocações de uma turma específica
func (c *AlocacaoController) GetAlocacoesByTurma(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	alocacoes, err := c.Repo.GetByTurmaID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alocacoes)
}

// OrganizarAlocacoesAutomaticas organiza alocações automaticamente para um dia e horário específicos
func (c *AlocacaoController) OrganizarAlocacoesAutomaticas(w http.ResponseWriter, r *http.Request) {
	// Estrutura para receber os dados da requisição
	type AlocacaoAutomaticaRequest struct {
		DiaSemana     string `json:"dia_semana"`
		HorarioInicio string `json:"horario_inicio"`
		HorarioFim    string `json:"horario_fim"`
	}

	// Decodificar o corpo da requisição
	var req AlocacaoAutomaticaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validar os dados recebidos
	if req.DiaSemana == "" || req.HorarioInicio == "" || req.HorarioFim == "" {
		http.Error(w, "Os campos dia_semana, horario_inicio e horario_fim são obrigatórios", http.StatusBadRequest)
		return
	}

	// Chamar o método do repositório para organizar as alocações automaticamente
	alocacoes, err := c.Repo.OrganizarAlocacoesAutomaticas(req.DiaSemana, req.HorarioInicio, req.HorarioFim)
	if err != nil {
		http.Error(w, "Erro ao organizar alocações automaticamente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar as alocações criadas
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alocacoes)
}
