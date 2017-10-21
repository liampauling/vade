package vade

import (
	"gofair"
	"log"
	"gofair/streaming"
)

type Listener struct {
	Trading		*gofair.Client
}


func NewInstance(trading *gofair.Client) (*Listener, error) {
	i := &Listener{
		Trading: trading,
	}
	return i, nil
}


func (l *Listener) ProcessMarket(eventTypeId string, marketId string) {

	downloadFolder := DownloadZip(eventTypeId, marketId)
	fileList := FileList(downloadFolder)
	log.Println(fileList)

	for _, file := range fileList {
		outputChannel := make(chan streaming.MarketBook)
		listener := streaming.Listener{OutputChannel: outputChannel}
		listener.AddMarketStream()

		go l.Trading.Historical.ParseHistoricalData(
			file,
			listener,
		)

		for i := range outputChannel {
			log.Println(i)
		}
	}
}
