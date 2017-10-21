package vade

import (
	"gofair/streaming"
	"os"
	"strconv"
	"fmt"
)

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


type RecordMarketBook struct {}

func (s RecordMarketBook) Start() {
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("PublishTime,MarketId,Inplay,TotalMatched,SelectionId,LastPriceTraded,RunnerTotalMatched\n")
}

func (s RecordMarketBook) Stop() {}

func (s RecordMarketBook) ProcessMarketBook(m streaming.MarketBook) {
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	publishTimeStr := strconv.Itoa(m.PublishTime)
	publishTimeTime, _ := MsToTime(publishTimeStr)

	for _, runner := range m.Runners {
		if runner.Status == "ACTIVE" {
			s_t := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v\n",
				publishTimeTime.Format("2006-01-02 15:04:05.999999"), m.MarketId, m.Inplay, m.TotalMatched,
				runner.SelectionId, runner.LastPriceTraded, runner.TotalMatched)

			if _, err = f.WriteString(s_t); err != nil {
				panic(err)
			}
		}
	}
}
