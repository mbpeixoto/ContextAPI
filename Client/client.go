package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	requisicao, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	resposta, err := http.DefaultClient.Do(requisicao)

	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resposta.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	var aux Cotacao
	json.Unmarshal(body, &aux)

	file, err := os.Create("./cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dolar := aux.USDBRL.Bid

	cotacaoJSON, err := json.Marshal(dolar)
	if err != nil {
		panic(err)
	}

	_, err = file.Write([]byte(fmt.Sprintf("Dólar: %s\n", cotacaoJSON)))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Dólar: %s\n", dolar)

}
