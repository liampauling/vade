package vade

import (
	"gofair"
	"gofair/streaming"
	"log"
)

type Listener struct {
	Trading		*gofair.Client
	Bucket		string
	Analytics	*Analytics
}


func NewListener(trading *gofair.Client, bucket string) (*Listener) {
	i := &Listener{
		Trading: trading,
		Bucket: bucket,
	}
	i.Analytics = new(Analytics)
	i.Analytics.Markets = make(map[string]MarketAnalytics)
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
