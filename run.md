# Como Executar o Projeto Item Detail ML

## Pré-requisitos

- [Go 1.18+](https://golang.org/dl/) instalado na máquina
- (Opcional) [Make](https://gnuwin32.sourceforge.net/packages/make.htm) para facilitar comandos no Windows/Linux
- (Opcional) [swag](https://github.com/swaggo/swag) para gerar a documentação Swagger

---

## Passos para execução

### 1. Instale as dependências

Abra o terminal na raiz do projeto e execute:

```sh
go mod tidy
```

---

### 2. Execute a aplicação

Se estiver usando o Makefile:

```sh
make run
```

Ou diretamente com Go:

```sh
go run ./cmd/main.go
```

---

### 3. Acesse a API

A aplicação estará disponível em:

```
http://localhost:8080
```

Exemplo de endpoint:

```
GET http://localhost:8080/products/{id}
```

---

### 4. Visualize a documentação Swagger

Após rodar a aplicação, acesse:

```
http://localhost:8080/swagger/index.html
```

Se ainda não gerou a documentação, execute:

```sh
swag init --generalInfo cmd/main.go
```

---

### 5. Executando os testes

- Todos os testes:
  ```sh
  make test
  ```
- Cobertura de testes:
  ```sh
  make cover
  ```
- Relatório de cobertura em HTML:
  ```sh
  make cover-html
  ```
