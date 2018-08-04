package main

import (
	"context"
	"html/template"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ThirdViewFormattingFuncMap() template.FuncMap {
	return template.FuncMap{
		"formatOddOrEven": formatOddOrEven,
	}
}

func formatOddOrEven(number string) string {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash(number)
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	return string(tx.Data())
}
