# vade
High performance betfair backtester using golang

## use

```golang
import "vade"
import "gofair"

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
    vade.PrintMarketBook{},
}

var eventTypeId = "7"
var marketId = "1.135603703"
backtest.ProcessMarket(eventTypeId, marketId, strategies)

```
