package mysterybox

import (
	"fmt"
	"github.com/1makarov/binance-nft-buy/internal/domain/mysterybox"
	bapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"os"
)

type MysteryBox struct {
	BoxList map[int64]Box
}

type Box struct {
	Information *mysterybox.Information
	Box         mysterybox.Box
	Quantity    int
	Status      bool
}

func GetActiveMysteryBoxList() (*MysteryBox, error) {
	list, err := bapi.NFTMysteryBoxList()
	if err != nil {
		return nil, err
	}
	b := make(map[int64]Box)
	for _, v := range list.Data {
		if v.MappingStatus == -1 {
			b[int64(len(b)+1)] = Box{Box: mysterybox.Box{ID: v.ProductID, Name: v.Name}}
		}
	}
	return &MysteryBox{BoxList: b}, nil
}

func (b *MysteryBox) SelectBox() *Box {
	for i, v := range b.BoxList {
		fmt.Println(fmt.Sprintf("%d. %s", i, v.Box.Name))
	}
	for {
		n := getSaleNumber()
		v, ok := b.BoxList[n]
		if ok {
			fmt.Println("The sale was successfully selected")
			return &v
		}
		fmt.Println("There is no sale under this number..")
	}
}

func getSaleNumber() (number int64) {
	fmt.Print("Enter the sale number: ")
	fmt.Fscan(os.Stdin, &number)
	return number
}

func (b *Box) InitBox() error {
	if err := b.getInformation(); err != nil {
		return err
	}
	b.getQuantity()
	return nil
}

func (b *Box) getInformation() error {
	information, err := bapi.NFTMysteryBoxInfo(b.Box.ID)
	if err != nil {
		return err
	}
	b.Information = information
	return nil
}

func (b *Box) getQuantity() {
	fmt.Printf("Enter the quantity to purchase: (max %d) ", b.Information.LimitPerBuy)
	var q int
	fmt.Fscan(os.Stdin, &q)
	b.Quantity = q
}
