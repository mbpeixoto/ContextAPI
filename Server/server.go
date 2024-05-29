package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" //(requires gcc and CGO_ENABLED=1 )
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func retornaCotacao(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(2 * time.Second)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	requisicao, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao criar requisição %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resposta, err := http.DefaultClient.Do(requisicao)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao realizar requisição %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resposta.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao ler resposta da API de cotação %s", err), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var aux Cotacao
	json.Unmarshal(body, &aux)

	db, err := sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Printf("Erro ao abrir o banco de dados: %v\n", err)
	}
	defer db.Close()

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err = db.ExecContext(ctx, "INSERT INTO cotacao (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		aux.USDBRL.Code, aux.USDBRL.Codein, aux.USDBRL.Name, aux.USDBRL.High, aux.USDBRL.Low, aux.USDBRL.VarBid, aux.USDBRL.PctChange, aux.USDBRL.Bid, aux.USDBRL.Ask, aux.USDBRL.Timestamp, aux.USDBRL.CreateDate)
	if err != nil {
		log.Printf("Erro ao inserir dados no banco de dados: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(aux)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", retornaCotacao)
	fmt.Println("Serviço rodando na porta 8080...")
	http.ListenAndServe(":8080", mux)
}
