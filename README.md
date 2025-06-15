# GO-LOGGER

Um microserviço centralizador de logs, construído em Go, projetado para receber, armazenar e gerenciar logs de múltiplas aplicações via uma API REST simples.

O serviço utiliza Fiber para alta performance na recepção de requisições HTTP, GORM para interação com o banco de dados e PostgreSQL para um armazenamento de dados.


## Como rodar?

### Pré-requisitos:
- [Go](https://go.dev/doc/install) (versão 1.23+)
- [Docker](https://docs.docker.com/engine/install/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### 1. Clone o repositório:
```bash
git clone https://github.com/LXSCA7/go-logger
cd go-logger
```

### 2. Configure as Variáveis de Ambiente

1.  Copie o arquivo de exemplo:
    ```bash
    cp .env.example .env
    ```
2.  Abra o arquivo `.env` e preencha com os valores desejados.

### 3. Configure as aplicações autorizadas:
1. Copie o arquivo de exemplo:
    ```bash
    cp apps.json.example apps.json
    ```
2. Abra o arquivo `apps.json` e preencha com as aplicações desejadas.

### 4. Rode o Projeto

Você pode rodar o projeto de duas formas:

#### Com Docker (Recomendado)

Este método sobe a aplicação e o banco de dados PostgreSQL em containers, já conectados.

```bash
docker compose up --build -d
```

A aplicação estará disponível em `http://localhost:<API_PORT>`.

#### Localmente

Este método exige que você tenha um banco de dados PostgreSQL rodando e acessível pela sua máquina.

```bash
# 1. Certifique-se que o banco de dados está rodando e as variáveis no .env apontam para ele.

# 2. Instale as dependências
go mod tidy

# 3. Rode a aplicação
go run main.go
```
