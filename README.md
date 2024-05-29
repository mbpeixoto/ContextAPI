
## Como testar
1. Abra o terminal e na pasta Server rodar server.go: `go run server.go`
2. Abra nova janela do terminal e na pasta Client rodar client.go: `go run client.go`

## Banco de dados
1. Criei o banco de dados sqlite3 cotacao.db em /Server
`sqlite3 cotacao.db`

2. Adicionei a tabela
```
CREATE TABLE IF NOT EXISTS cotacao (
    code TEXT NOT NULL,
    codein TEXT NOT NULL,
    name TEXT NOT NULL,
    high TEXT NOT NULL,
    low TEXT NOT NULL,
    varBid TEXT NOT NULL,
    pctChange TEXT NOT NULL,
    bid TEXT NOT NULL,
    ask TEXT NOT NULL,
    timestamp TEXT NOT NULL,
    create_date TEXT NOT NULL
);
```

## Sobre o projeto
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.