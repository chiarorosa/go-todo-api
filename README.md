# go-todo-api

API REST simples para gerenciamento de tarefas (to-do) escrita em Go, utilizando o framework Gin. Os dados são mantidos em memória, ideal para estudos ou prototipação. Todo o desenvolvimento segue o [método de aprendizado MCA](https://github.com/chiarorosa/mca-method/blob/main/golang/go_beginner.guideline.md), guiando a evolução do projeto passo a passo.

## Visão Geral

- Exposição de endpoints REST para listar e criar tarefas.
- Repositório em memória com incremento automático de IDs.
- Estrutura modular (camadas de handler e repositório) pronta para evoluir para banco de dados ou outras integrações.

## Tecnologias

- Go 1.25+
- Gin Web Framework

## Como Executar

1. **Pré-requisitos**: instale o Go 1.25 ou superior.
2. **Instalação**:

   ```bash
   git clone https://github.com/chiarorosa/go-todo-api.git
   cd go-todo-api
   go mod download
   ```

3. **Execução**:

   ```bash
   go run ./cmd/server
   ```

   O servidor inicia em `http://localhost:8080`.

## Endpoints

- **GET `/tasks`**
  Retorna todas as tarefas cadastradas.

  ```bash
  curl http://localhost:8080/tasks
  ```

  ```json
  [
    {
      "id": 1,
      "title": "Revisar pull request",
      "description": "Verificar comentários e aprovar",
      "done": false
    }
  ]
  ```

- **POST `/tasks`**
  Cria uma nova tarefa. Envie um JSON com `title`, `description` e `done`.

  ```bash
  curl -X POST http://localhost:8080/tasks \
    -H "Content-Type: application/json" \
    -d '{
      "title": "Estudar Go",
      "description": "Praticar goroutines",
      "done": false
    }'
  ```

  ```json
  {
    "id": 2,
    "title": "Estudar Go",
    "description": "Praticar goroutines",
    "done": false
  }
  ```
