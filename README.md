# Item Detail ML

Projeto em Go para detalhamento de itens, simulação de produtos e meios de pagamento, seguindo princípios de Clean Architecture.

## Estrutura do Projeto

```
.
├── cmd/                    # Ponto de entrada da aplicação
├── internal/
│   ├── domain/             # Entidades de domínio
│   ├── factory/            # Factories para responses
│   ├── handler/            # Handlers HTTP (controllers)
│   ├── repository/         # Repositórios (acesso a dados)
│   ├── usecase/            # Casos de uso (regras de negócio)
├── mock_data/              # Dados mockados em JSON
├── test/                   # Testes unitários e arquivos de mock
├── Makefile                # Comandos utilitários
└── go.mod
```

## Como rodar

1. **Instale as dependências:**

   ```sh
   go mod tidy
   ```

2. **Execute a aplicação:**

   ```sh
   make run
   ```

3. **Acesse a API:**
   - Exemplo: `GET http://localhost:8080/products/{id}`

## Testes

- **Rodar todos os testes:**

  ```sh
  make test
  ```

- **Cobertura de testes:**

  ```sh
  make cover
  ```

- **Relatório de cobertura em HTML:**
  ```sh
  make cover-html
  ```

## Visualizando a documentação Swagger

Após rodar a aplicação, acesse a documentação interativa da API em:

```
http://localhost:8080/swagger/index.html
```

Você verá todos os endpoints, parâmetros e exemplos de resposta gerados automaticamente a partir dos comentários do código.

Se ainda não gerou a documentação, execute:

```sh
swag init --generalInfo cmd/main.go
```

Isso criará a pasta `docs/` necessária para o Swagger funcionar.

## Mocks

Mocks de interfaces são gerados com [mockery](https://github.com/vektra/mockery):

```sh
mockery --name=NomeDaInterface --output=internal/mocks
```

## Observabilidade

O projeto utiliza [logrus](https://github.com/sirupsen/logrus) para logging estruturado.

## Organização dos testes

- Testes unitários estão em `test/` e usam [testify](https://github.com/stretchr/testify) e mocks do mockery.
- Mocks de dados estão em arquivos separados para fácil reutilização.

## Contribuição

Sugestões e melhorias são bem vindas!

---

**Desenvolvido por:**  
Priscila Albertini da Silva
