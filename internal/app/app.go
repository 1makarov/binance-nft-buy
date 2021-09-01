package app

import (
	"fmt"
	"github.com/1makarov/binance-nft-buy/internal/pkg/account"
	"github.com/1makarov/binance-nft-buy/internal/pkg/mysterybox"
	bapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"log"
	"time"
)

func App(account *account.Account, box *mysterybox.Box) {
	defer fmt.Scanf("\n")

	body, err := bapi.MarshalMysteryBoxBuy(box.Box.ID, box.Quantity)
	if err != nil {
		log.Fatalf("error marshal buy box: %s\n", err.Error())
	}
	req := account.Auth.NFTMysteryBoxGenerateRequest(body)

	log.Println("Waiting started successfully")
	wait(box.Information.StartTime)
	log.Println("Start buy")

	go func() {
		for {
			if !box.Status {
				resp, err := account.Auth.NFTMysteryBoxBuy(req)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println(string(resp.Body()))
				return
			} else {
				return
			}
		}
	}()

	time.Sleep(6 * time.Second)
	box.Status = true
	time.Sleep(1 * time.Second)

	fmt.Println("Purchases are completed")
}

func wait(s int64) {
	t := time.Unix(s, 0).UTC().Add(-3 * time.Second).Unix()
	for {
		if time.Now().UTC().Unix() >= t {
			return
		}
	}
}
