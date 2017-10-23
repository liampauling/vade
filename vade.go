package vade

import (
	"gofair"
	"gofair/streaming"
	"log"
)

type Listener struct {
	Trading		*gofair.Client
	Bucket	string
}


func NewListener(trading *gofair.Client, bucket string) (*Listener) {
	i := &Listener{
		Trading: trading,
		Bucket: bucket,
	}
	return i
}


func (l *Listener) ProcessMarket(eventTypeId string, marketId string, strategies []Strategy) {

	downloadFolder := DownloadZip(l.Bucket, eventTypeId, marketId)
	fileList := FileList(downloadFolder)
	log.Println(fileList)

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
			for _, strategy := range strategies {
				strategy.ProcessMarketBook(marketBook)
			}
		}

		for _, strategy := range strategies {
			strategy.Stop()
		}
	}
}
