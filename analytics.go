package vade

import (
	"gofair/streaming"
	"log"
	"fmt"
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

	for _, r := range m.Runners {
		if r.SelectionId == 11180319 && len(r.EX.TradedVolume) > 0 {
			fmt.Println(r.SelectionId, r.LastPriceTraded, r.EX.AvailableToBack[0], r.EX.AvailableToLay[0],
				r.EX.TradedVolume)
		}
	}

}
