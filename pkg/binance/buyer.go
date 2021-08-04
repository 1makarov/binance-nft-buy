package binance

import (
	"fmt"
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"log"
	"os"
	"time"
)

func (c *Client) Start() {
	defer fmt.Scanf("\n")
	// generate httpclient
	c.api.GenerateHttpClient()
	//get user id (check cookie)
	user, err := c.api.User()
	if err != nil {
		log.Fatalln(err)
	}
	//notification for cookie
	notification(user.Data.Email)
	//get nft info
	nftInfo, err := c.api.NFTInfo()
	if err != nil {
		log.Fatalln(err)
	}
	// handle nft
	box := handleNFTInfo(nftInfo)
	//generate bytes from json
	boxjson := binanceapi.MarshalBoxBuy(binanceapi.MysteryBox{
		ProductId: box.Productid,
		Amount:    getQuantity(),
	})
	// generate request to buy box
	req := c.api.GenerateRequest(binanceapi.URLBuy, boxjson)
	// wait buy time
	waitBuyTime(box.StartTime, getDelay())
	// buy box
	if err = c.api.MysteryBoxBuy(req); err != nil {
		log.Fatalln(err)
	}
}

func handleNFTInfo(b *binanceapi.NftInfoResponse) binanceapi.Box {
	for i, box := range b.Boxes {
		if box.MappingStatus == -1 {
			fmt.Println(fmt.Sprintf("%d. %s", i+1, box.Name))
		}
	}
	return b.Boxes[getSaleNumber()].Box
}

func getQuantity() (quantity int) {
	fmt.Print("Quantity to buy?\t")
	fmt.Fscan(os.Stdin, &quantity)
	return
}

func getDelay() (delay int64) {
	fmt.Print("How many seconds delay?\t")
	fmt.Fscan(os.Stdin, &delay)
	return
}

func getSaleNumber() (number int) {
	fmt.Print("Enter the sale number:\t")
	fmt.Fscan(os.Stdin, &number)
	return
}

func notification(email string) {
	log.Println(fmt.Sprintf("You have entered working cookies, your email: %s", email))
}

func waitBuyTime(starttime, delay int64) {
	for {
		if time.Now().UTC().Unix() >= starttime+(delay) {
			return
		}
	}
}
