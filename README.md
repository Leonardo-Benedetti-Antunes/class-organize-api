# Class Organize - Sistema de Gerenciamento de Salas, Turmas e Professores

Este é um sistema de back-end desenvolvido em Go para gerenciar a alocação de salas, turmas e professores em uma instituição de ensino.

## Estrutura do Projeto

```
.
├── config/
│   └── database.go     # Configuração de conexão com o banco de dados
├── controllers/
│   └── controllers.go  # Controladores para gerenciar requisições HTTP
├── models/
│   └── models.go       # Definição dos modelos de dados
├── repositories/
│   └── repositories.go # Operações de banco de dados
├── .env                # Variáveis de ambiente
├── go.mod              # Dependências do projeto
├── main.go             # Ponto de entrada da aplicação
└── README.md           # Documentação do projeto
```

## Funcionalidades

- Gerenciamento de professores (CRUD)
- Gerenciamento de salas (CRUD)
- Gerenciamento de turmas (CRUD)
- Alocação de professores, salas e turmas (CRUD)
- Verificação de disponibilidade de salas
- Consulta de alocações por professor, sala ou turma

## Requisitos

- Go 1.16 ou superior
- PostgreSQL 12 ou superior

## Configuração

1. Clone o repositório
2. Configure o arquivo `.env` com as credenciais do banco de dados:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=class_organize
```

3. Instale as dependências:

```bash
go mod tidy
```

4. Execute a aplicação:

```bash
go run main.go
```

O servidor será iniciado na porta 8080.

## API Endpoints

### Professores

- `GET /api/professores` - Listar todos os professores
- `GET /api/professores/{id}` - Obter um professor específico
- `POST /api/professores` - Criar um novo professor
- `PUT /api/professores/{id}` - Atualizar um professor
- `DELETE /api/professores/{id}` - Remover um professor

### Salas

- `GET /api/salas` - Listar todas as salas
- `GET /api/salas/{id}` - Obter uma sala específica
- `POST /api/salas` - Criar uma nova sala
- `PUT /api/salas/{id}` - Atualizar uma sala
- `DELETE /api/salas/{id}` - Remover uma sala

### Turmas

- `GET /api/turmas` - Listar todas as turmas
- `GET /api/turmas/{id}` - Obter uma turma específica
- `POST /api/turmas` - Criar uma nova turma
- `PUT /api/turmas/{id}` - Atualizar uma turma
- `DELETE /api/turmas/{id}` - Remover uma turma

### Alocações

- `GET /api/alocacoes` - Listar todas as alocações
- `GET /api/alocacoes/{id}` - Obter uma alocação específica
- `POST /api/alocacoes` - Criar uma nova alocação
- `PUT /api/alocacoes/{id}` - Atualizar uma alocação
- `DELETE /api/alocacoes/{id}` - Remover uma alocação

### Consultas Especiais

- `GET /api/alocacoes/sala/{id}` - Listar alocações por sala
- `GET /api/alocacoes/professor/{id}` - Listar alocações por professor
- `GET /api/alocacoes/turma/{id}` - Listar alocações por turma

## Exemplos de Uso

### Criar um Professor

```bash
curl -X POST http://localhost:8080/api/professores \
  -H "Content-Type: application/json" \
  -d '{"nome":"João Silva","email":"joao@exemplo.com","formacao":"Mestrado em Computação","disciplina":"Programação Web"}'
```

### Criar uma Sala

```bash
curl -X POST http://localhost:8080/api/salas \
  -H "Content-Type: application/json" \
  -d '{"numero":"101","capacidade":40,"bloco":"A","tipo":"Laboratório"}'
```

### Criar uma Turma

```bash
curl -X POST http://localhost:8080/api/turmas \
  -H "Content-Type: application/json" \
  -d '{"nome":"Turma A","curso":"Sistemas de Informação","periodo":"Noturno","quant_alunos":35}'
```

### Criar uma Alocação

```bash
curl -X POST http://localhost:8080/api/alocacoes \
  -H "Content-Type: application/json" \
  -d '{"professor_id":1,"sala_id":1,"turma_id":1,"dia_semana":"Segunda","horario_inicio":"19:00","horario_fim":"22:30"}'
```

## Licença

Este projeto está licenciado sob a licença MIT.