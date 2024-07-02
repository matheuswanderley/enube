# Documentação do Projeto

## Requisitos

- [Docker](https://www.docker.com/get-started) instalado
- [Docker Compose](https://docs.docker.com/compose/install/) instalado
- [Go](https://golang.org/doc/install) 1.20. instalado (para execução sem Docker)

## Passos para Executar o Projeto

### 1. Clonar o Repositório

Clone o repositório para sua máquina local:

```sh
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio
```

### 2. Instalar Dependências

Certifique-se de ter o Go instalado e execute o comando para instalar as dependências do projeto:

```sh
go mod tidy
```

### 3. Executar a Aplicação

Para executar a aplicação, utilize o comando:
```sh
go run main.go
```

### 4. Acessar o Swaggers

Com a aplicação rodando, o Swagger estará disponível em:
```
http://localhost:8080/swagger/index.html
```
### 5. Executar os Testes

Para executar os testes unitários, utilize o comando:
```sh
go test ./...
```

### 6. Executar com docker

Para executar a aplicação, utilize o comando:
```sh
docker compose up --build
```