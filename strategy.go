package vade

import "fmt"
import "gofair/streaming"

type Strategy interface {
	Start()
	ProcessMarketBook(m streaming.MarketBook)
	Stop()
}


type PrintMarketBook struct {}

func (s PrintMarketBook) Start() {}

func (s PrintMarketBook) Stop() {}

func (s PrintMarketBook) ProcessMarketBook(m streaming.MarketBook) {
	fmt.Println(m)
}
