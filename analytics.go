package vade

import (
	"gofair/streaming"
	"log"
)

type Analytics struct {
	Markets map[string]MarketAnalytics
}

func (a *Analytics) ProcessMarketBook(m streaming.MarketBook) {
	if marketAnalytics, ok := a.Markets[m.MarketId]; ok {
		marketAnalytics.process(m)
	} else {
		marketAnalytics := createMarketAnalytics(m)
		a.Markets[m.MarketId] = *marketAnalytics
		log.Println("Created new analytics cache", m.MarketId)
	}
}

func createMarketAnalytics(m streaming.MarketBook) *MarketAnalytics {
	ma := &MarketAnalytics{
		MarketBook: m,
	}
	return ma
}

type MarketAnalytics struct {
	MarketBook streaming.MarketBook
}

func (ma *MarketAnalytics) process(m streaming.MarketBook) {

}
