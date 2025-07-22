# 🚦 FullCycle Stress Teste

## 🧾 Objetivo

Criar um sistema de linha de comando (CLI) em Go para realizar testes de carga em um serviço web.
O usuário deve fornecer:

* A URL do serviço
* O número total de requisições
* A quantidade de chamadas simultâneas

---

## ⚙️ Parâmetros de Entrada

| Parâmetro       | Descrição                      |
| --------------- | ------------------------------ |
| `--url`         | URL do serviço a ser testado   |
| `--requests`    | Número total de requisições    |
| `--concurrency` | Número de chamadas simultâneas |

---

## 📊 Exemplo de Relatório

```
Relatório de execução Stress-Test

Tempo decorrido                  = 53.451400 segundos
Tempo total executando requests = 531.492428 segundos
Total de requests realizadas     = 1000
Total de requests status 200     = 0

Erros de execução:
- Status: 429, Total: 1000
```

---

## ▶️ Como Executar

### Usando Go

```bash
go run cmd/root.go cmd/main.go --url=http://google.com --requests=1000 --concurrency=10
```

### Usando Docker

```bash
docker run <sua-imagem-docker> --url=http://google.com --requests=1000 --concurrency=10
```

---