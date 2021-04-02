package level1

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"time"
)

type RegionalUpdateMessage struct {
	Symbol              string
	Exchange            string
	RegionalBid         float64
	RegionalBidSize     int64
	RegionalBidTime     time.Duration
	RegionalAsk         float64
	RegionalAskSize     int64
	RegionalAskTime     time.Duration
	FractionDisplayCode int64
	DecimalPrecision    int64
	MarketCenter        int64
}

func (t *RegionalUpdateMessage) Unmarshal(values []string) {
	t.Symbol = proto.Value(values, 1)
	t.Exchange = proto.Value(values, 2)
	t.RegionalBid = proto.ParseFloat(proto.Value(values, 3))
	t.RegionalBidSize = proto.ParseInt(proto.Value(values, 4))
	t.RegionalBidTime = proto.ParseDurationSecs(proto.Value(values, 5))
	t.RegionalAsk = proto.ParseFloat(proto.Value(values, 6))
	t.RegionalAskSize = proto.ParseInt(proto.Value(values, 7))
	t.RegionalAskTime = proto.ParseDurationSecs(proto.Value(values, 8))
	t.FractionDisplayCode = proto.ParseInt(proto.Value(values, 9))
	t.DecimalPrecision = proto.ParseInt(proto.Value(values, 10))
	t.MarketCenter = proto.ParseInt(proto.Value(values, 11))
}
