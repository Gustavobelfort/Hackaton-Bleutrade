package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math"
	"math/big"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/text/language"
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

	blockNumber := big.NewInt(6007492)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144

	for _, tx := range block.Transactions() {

		if strings.ContainsAny(string(tx.Data()), "兽") {

			// i := len(string(tx.Data())) / 5000

			translated := make([]string, 0)

			splited := make([]string, 0)

			original := strings.Split(string(tx.Data()), "\n")
			fmt.Println(len(original))

			for i := 0; i < len(original); i += 30 {
				max := math.Min(float64(i+31), float64(len(original)))
				splited = append(splited, strings.Join(original[i:int(max)], "\n"))

				//original = original[:i+copy(original[i:], original[i+1:])]

			}
			fmt.Println(len(splited))
			//fmt.Println(splited)

			// fmt.Println(len(original))

			// for j := 0; j < i; j++ {
			// 	original = append(original, string(tx.Data())[j*5000:(j+1)*5000])
			// }

			// original = append(original, string(tx.Data())[i*5000:])

			ctx := context.Background()

			// Creates a client.
			client, err := translate.NewClient(ctx)
			if err != nil {
				log.Fatalf("Failed to create client: %v", err)
			}
			// Sets the target language.
			target, err := language.Parse("pt")
			if err != nil {
				log.Fatalf("Failed to parse target language: %v", err)
			}

			if err != nil {
				log.Fatalf("Failed to translate text: %v", err)

			}

			// tt, _ := client.Translate(ctx, original[0:1], target, nil)

			// for index, k := range original {
			// 	client.Translate
			// 	translated[index], _ = client.Translate(ctx, []string{string(k)}, target, nil)
			// }

			fmt.Println(len(splited))

			for l := 0; l < len(splited); l = l + 1 {

				translation, _ := client.Translate(ctx, []string{splited[l]}, target, nil)
				translated = append(translated, translation[0].Text)
			}

			fmt.Println(translated)
			// translations, err := client.Translate(ctx, []string{original[0]}, target, nil)
			// if err != nil {
			// 	log.Fatalf("Failed to translate text: %v", err)
			// }
			fmt.Println(len(translated))
			return strings.Join(translated, "\n")

		}
	}

	return ""
}
