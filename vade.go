package vade

import (
	"gofair"
	"log"
	"gofair/streaming"
)

type Listener struct {
	Trading		*gofair.Client
}


func NewListener(trading *gofair.Client) (*Listener) {
	i := &Listener{
		Trading: trading,
	}
	return i
}


func (l *Listener) ProcessMarket(eventTypeId string, marketId string, strategies []Strategy) {

	downloadFolder := DownloadZip(eventTypeId, marketId)
	fileList := FileList(downloadFolder)
	log.Println(fileList)

	for _, file := range fileList {
		outputChannel := make(chan streaming.MarketBook)
		listener := streaming.Listener{OutputChannel: outputChannel}
		listener.AddMarketStream()

		// todo strategy.start()

		go l.Trading.Historical.ParseHistoricalData(
			file,
			listener,
		)

		for marketBook := range outputChannel {
			for _, strategy := range strategies {
				strategy.ProcessMarketBook(marketBook)
			}
		}

		// todo strategy.stop()
	}
}
