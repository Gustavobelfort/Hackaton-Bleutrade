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

func formatOddOrEven(number int) string {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash("0xb1ed364e4333aae1da4a901d5231244ba6a35f9421d4607f7cb90d60bf45578a")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	return string(tx.Data())
}
