# 💳 Lab Supply Payment

Microsserviço desenvolvido em Go para processamento de pagamentos do sistema Lab Supply.

O serviço recebe solicitações de pagamento de forma assíncrona através do RabbitMQ, processa a transação e publica o resultado para que a API principal atualize o status do pedido.

Este projeto integra o desafio final do módulo 3 da formação em Go, demonstrando comunicação entre microsserviços, mensageria e separação de responsabilidades.

---

## 🛠️ Tecnologias

- Go 1.26.4
- PostgreSQL
- RabbitMQ
- Docker
- `net/http`
- `database/sql`
- UUID
- `github.com/lib/pq`
- `github.com/google/uuid`
- `github.com/rabbitmq/amqp091-go`

## 🏗️ Arquitetura

O serviço utiliza arquitetura em camadas, separando domínio, regras de negócio, persistência e mensageria.

```text
RabbitMQ
    │
    ▼
Consumer
    │
    ▼
Payment Service
    │
    ▼
Repository
    │
    ▼
PostgreSQL
```

### Principais responsabilidades:

- **Consumer**: recebe e interpreta os eventos de pagamento.
- **Service**: executa o caso de uso e coordena as regras de negócio.
- **Domain**: representa o pagamento, seus estados e transições.
- **Repository**: realiza a persistência dos pagamentos.
- **Messaging**: gerencia o consumo e a publicação de eventos no RabbitMQ.

## 🔄 Fluxo de pagamento

O serviço recebe eventos `PaymentRequested` publicados pela API principal através do RabbitMQ.

Após o processamento, publica um dos seguintes eventos:

- `payment.approved`
- `payment.failed`

```text
labsupply-api
      │
      │ PaymentRequested
      ▼
  RabbitMQ
      │
      ▼
labsupply-payment
      │
      ├── payment.approved
      │
      └── payment.failed
```

## 💳 Estados do pagamento

O pagamento possui três estados:

- `PENDING`: pagamento criado e ainda não processado.
- `APPROVED`: pagamento aprovado.
- `FAILED`: pagamento não processado com sucesso.

O consumidor utiliza ACK para confirmar o recebimento da mensagem após o processamento.

## 📁 Estrutura do projeto

```text
cmd/
└── api/

internal/
├── config/
├── controllers/
├── database/
├── domain/
├── events/
├── messaging/
├── repository/
├── routes/
└── service/
```

## 📦 Funcionalidades

- Recebimento de solicitações de pagamento via RabbitMQ.
- Criação e persistência de pagamentos.
- Validação do valor do pagamento.
- Controle dos estados `PENDING`, `APPROVED` e `FAILED`.
- Publicação dos resultados do processamento.
- Tratamento de pagamentos aprovados e recusados.

## ⚙️ Configuração

O serviço depende de PostgreSQL e RabbitMQ em execução.

Crie um arquivo `.env` na raiz do projeto:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=labsupply_payment

RABBITMQ_URL=amqp://guest:guest@localhost:5672/
```

Não versione o .env com credenciais reais. Utilize o .env.example como referência.

## ▶️ Execução

### 1. Inicie a infraestrutura

Certifique-se de que o PostgreSQL e o RabbitMQ estejam em execução.

### 2. Execute o serviço

Na raiz do projeto:

```bash
go run ./cmd/api
```

O serviço estará disponível na porta 8081.

### 3. Execute a API principal

Para testar o fluxo completo, execute também o labsupply-api em outro terminal:

```bash
go run ./cmd/app
```

A API principal utiliza a porta 8080.

Os dois serviços devem estar conectados ao mesmo RabbitMQ.

## 🧪 Validação

O serviço foi validado com:
```bash
go test ./...
go vet ./...
go build ./...
```

Também foram realizados testes de integração envolvendo o processamento de pagamentos, publicação dos eventos payment.approved e payment.failed, atualização dos pedidos e confirmação das mensagens através de ACK.

## 🔗 Projeto principal

Este microsserviço faz parte do projeto Lab Supply e se comunica com a API principal:

`https://github.com/Ludimila-Araujo/lab-supply-api`

## 👩‍💻 Autora

**Ludimila Araújo**

Projeto desenvolvido como desafio final do módulo 3 da formação em Go, com foco em APIs REST, arquitetura em camadas, mensageria e integração entre microsserviços.