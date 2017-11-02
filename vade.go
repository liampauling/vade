package vade

import (
	"gofair"
	"gofair/streaming"
)

type Listener struct {
	Trading		*gofair.Client
	Analytics	*Analytics
}


func NewListener(trading *gofair.Client) (*Listener) {
	i := &Listener{
		Trading: trading,
	}
	i.Analytics = new(Analytics)
	i.Analytics.Markets = make(map[string]MarketAnalytics)
	return i
}


func (l *Listener) ProcessMarket(fileList []string, strategies []Strategy) {

	for _, file := range fileList {
		outputChannel := make(chan streaming.MarketBook)
		listener := streaming.Listener{OutputChannel: outputChannel}
		listener.AddMarketStream()

		for _, strategy := range strategies {
			strategy.Start()
		}

		go l.Trading.Historical.ParseHistoricalData(
			file,
			listener,
		)

		for marketBook := range outputChannel {

			l.Analytics.ProcessMarketBook(marketBook)

			for _, strategy := range strategies {
				strategy.ProcessMarketBook(marketBook)
			}
		}

		for _, strategy := range strategies {
			strategy.Stop()
		}
	}
}
