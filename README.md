# vāde
High performance betfair backtester using golang

Assumes AWS credentials are in bash profile and there is a bucket in the following format containing streaming data:

    'bucket'/marketdata/streaming/'event_type_id'/'market_id'.zip

## use

```golang
import "vade"
import "gofair"

func main() {
	config := &gofair.Config{
		CertFile: "/gocerts/client-2048.crt",
		KeyFile: "/gocerts/client-2048.key",
	}

	trading, err := gofair.NewClient(config)
	if err != nil {
		panic(err)
	}

	backtest := vade.NewListener(trading, "flumine")

	// create strategies list
	strategies := []vade.Strategy{
		vade.RecordMarketBook{},
	}

	var eventTypeId = "7"
	var marketId = "1.135904788"
	backtest.ProcessMarket(eventTypeId, marketId, strategies)
}
```
