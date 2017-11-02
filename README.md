# vāde
High performance betfair backtester using golang


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

	backtest := vade.NewListener(trading)

	// create strategies list
	strategies := []vade.Strategy{
		vade.RecordMarketBook{},
	}

	var f = []string{"1.135904788", "1.135904789"}
	backtest.ProcessMarket(eventTypeId, marketId, strategies)
}
```
