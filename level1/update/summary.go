package update

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"time"
)

type SummaryMessage struct {
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

func (u *SummaryMessage) Unmarshal(values []string) {
	u.Symbol = proto.Value(values, 1)
	u.MostRecentTrade = proto.ParseFloat(proto.Value(values, 2))
	u.MostRecentTradeSize = proto.ParseInt(proto.Value(values, 3))
	u.MostRecentTradeTime = proto.ParseDuration(values[3])
	u.MostRecentTradeMarketCenter = proto.ParseInt(proto.Value(values, 5))
	u.TotalVolume = proto.ParseInt(proto.Value(values, 6))
	u.Bid = proto.ParseFloat(proto.Value(values, 7))
	u.BidSize = proto.ParseInt(proto.Value(values, 8))
	u.Ask = proto.ParseFloat(proto.Value(values, 9))
	u.AskSize = proto.ParseInt(proto.Value(values, 10))
	u.Open = proto.ParseFloat(proto.Value(values, 11))
	u.High = proto.ParseFloat(proto.Value(values, 12))
	u.Low = proto.ParseFloat(proto.Value(values, 13))
	u.Close = proto.ParseFloat(proto.Value(values, 14))
	u.MessageContents = proto.Value(values, 15)
	u.MostRecentTradeConditions = proto.Value(values, 16)

}
