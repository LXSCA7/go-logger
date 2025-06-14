# GO-LOGGER

Aplicação feita para centralizar os logs das minhas APIs.


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

| Variável | Descrição | Exemplo |
| :--- | :--- | :--- |
| `APP_ENV` | Define o ambiente (`development` ou `production`). Em `production`, o .env não é carregado. | `development` |
| `API_PORT` | Porta em que a API irá rodar. | `3000` |
| `DB_HOST` | Host do banco de dados. Use `localhost` para local, `postgres` para Docker. | `postgres` |
| `DB_PORT` | Porta do banco de dados. | `5432` |
| `DB_USER` | Usuário de conexão com o banco. | `postgres` |
| `DB_PASS` | Senha de conexão com o banco. | `yourSuperSecretPassword` |
| `DB_NAME` | Nome do banco de dados a ser utilizado. | `go_logger_db` |

### 3. Rode o Projeto

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
