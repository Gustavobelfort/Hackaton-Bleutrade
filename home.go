package main

import (
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/text/language"
)

func HomeFormattingFuncMap() template.FuncMap {

	return template.FuncMap{
		"HomeformatOddOrEven": HomeformatOddOrEven,
	}
}

func HomeformatOddOrEven() string {
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	/*
		blockNumber := big.NewInt(6007492)
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}

			count, err := client.TransactionCount(context.Background(), block.Hash())
			if err != nil {
				log.Fatal(err)
			}*/
	file, err := os.Open("hashNews.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(data))
	newsExp := make([]string, 0)
	news := strings.Split(string(data), "\n")
	currentNews := ""
	fmt.Println(len(news))

	fmt.Println("oi")
	fmt.Println(news[indexText])
	fmt.Println(len("0xb1ed364e4333aae1da4a901d5231244ba6a35f9421d4607f7cb90d60bf45578a"))
	if len(news[indexText]) == 67 {
		currentNews = (news[indexText])[:len(news[indexText])-1]
	} else {
		currentNews = news[indexText]
	}

	txHash := common.HexToHash(currentNews)

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	translated := make([]string, 0)

	splited := make([]string, 0)

	original := strings.Split(string(tx.Data()), "\n")

	fmt.Println(len(original))

	for i := 0; i < len(original); i += 30 {
		max := math.Min(float64(i+31), float64(len(original)))
		splited = append(splited, strings.Join(original[i:int(max)], "\n"))

	}
	fmt.Println(len(splited))

	ctx := context.Background()

	// Creates a client.
	client2, err := translate.NewClient(ctx)
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

		translation, _ := client2.Translate(ctx, []string{splited[l]}, target, nil)
		translated = append(translated, translation[0].Text)
	}

	fmt.Println(translated)

	fmt.Println(len(translated))
	newsExp = append(newsExp, (strings.Join(translated, "\n")))

	result := ""
	for k := 0; k < len(newsExp); k++ {
		result += newsExp[k]
	}
	indexText += 1
	if indexText >= 2 {
		indexText = 0
	}
	return result

	/*
		for indexText := 0; indexText < len(news); indexText += 1 {
			fmt.Println("oi")
			fmt.Println(news[indexText])
			fmt.Println(len("0xb1ed364e4333aae1da4a901d5231244ba6a35f9421d4607f7cb90d60bf45578a"))
			if len(news[indexText]) == 67 {
				currentNews = (news[indexText])[:len(news[indexText])-1]
			} else {
				currentNews = news[indexText]
			}

			txHash := common.HexToHash(currentNews)

			tx, _, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				log.Fatal(err)
			}

			translated := make([]string, 0)

			splited := make([]string, 0)

			original := strings.Split(string(tx.Data()), "\n")

			fmt.Println(len(original))

			for i := 0; i < len(original); i += 30 {
				max := math.Min(float64(i+31), float64(len(original)))
				splited = append(splited, strings.Join(original[i:int(max)], "\n"))

			}
			fmt.Println(len(splited))

			ctx := context.Background()

			// Creates a client.
			client2, err := translate.NewClient(ctx)
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

				translation, _ := client2.Translate(ctx, []string{splited[l]}, target, nil)
				translated = append(translated, translation[0].Text)
			}

			fmt.Println(translated)

			fmt.Println(len(translated))
			newsExp = append(newsExp, (strings.Join(translated, "\n")))
		}
		result := ""
		for k := 0; k < len(newsExp); k++ {
			result += newsExp[k]
		}*/

}
