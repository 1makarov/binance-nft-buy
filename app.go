package main

import (
	"github.com/1makarov/binance-nft-buy/internal/app"
	acc "github.com/1makarov/binance-nft-buy/internal/domain/account"
	"github.com/1makarov/binance-nft-buy/internal/pkg/account"
	"github.com/1makarov/binance-nft-buy/internal/pkg/mysterybox"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := initEnv(); err != nil {
		log.Println(err)
		return
	}

	a, err := account.InitAccount(acc.Setting{
		Proxy: os.Getenv("PROXY"),
		BAuth: &acc.BAuth{Cookie: os.Getenv("COOKIE"), Csrf: os.Getenv("CSRFTOKEN")},
	})

	if err != nil {
		log.Println(err)
		return
	}

	if err = a.HandleAccount(); err != nil {
		log.Println(err)
		return
	}

	boxList, err := mysterybox.GetActiveMysteryBoxList()
	if err != nil {
		log.Println(err)
		return
	}

	box, err := boxList.SelectBox()
	if err != nil {
		log.Println(err)
		return
	}

	if err = box.InitBox(); err != nil {
		log.Println(err)
		return
	}

	app.App(a, box)
}

func initEnv() error {
	return godotenv.Load()
}
