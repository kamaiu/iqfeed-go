package level1

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"time"
)

type UpdateSummaryMessage struct {
	Symbol                      string
	MostRecentTrade             float64
	MostRecentTradeSize         int64
	MostRecentTradeTime         time.Duration
	MostRecentTradeMarketCenter int64
	TotalVolume                 int64
	Bid                         float64
	BidSize                     int64
	Ask                         float64
	AskSize                     int64
	Open                        float64
	High                        float64
	Low                         float64
	Close                       float64
	MessageContents             string
	MostRecentTradeConditions   string
}

func (t *UpdateSummaryMessage) Unmarshal(values []string) {
	t.Symbol = proto.Value(values, 1)
	t.MostRecentTrade = proto.ParseFloat(proto.Value(values, 2))
	t.MostRecentTradeSize = proto.ParseInt(proto.Value(values, 3))
	t.MostRecentTradeTime = proto.ParseDuration(proto.Value(values, 4))
	t.MostRecentTradeMarketCenter = proto.ParseInt(proto.Value(values, 5))
	t.TotalVolume = proto.ParseInt(proto.Value(values, 6))
	t.Bid = proto.ParseFloat(proto.Value(values, 7))
	t.BidSize = proto.ParseInt(proto.Value(values, 8))
	t.Ask = proto.ParseFloat(proto.Value(values, 9))
	t.AskSize = proto.ParseInt(proto.Value(values, 10))
	t.Open = proto.ParseFloat(proto.Value(values, 11))
	t.High = proto.ParseFloat(proto.Value(values, 12))
	t.Low = proto.ParseFloat(proto.Value(values, 13))
	t.Close = proto.ParseFloat(proto.Value(values, 14))
	t.MessageContents = proto.Value(values, 15)
	t.MostRecentTradeConditions = proto.Value(values, 16)
}
