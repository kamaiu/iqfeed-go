package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type TradeConditionMessage struct {
	TradeConditionId int64
	ShortName        string
	LongName         string
	RequestId        string
}

func (l *TradeConditionMessage) Unmarshal(values []string) {
	l.TradeConditionId = proto.ParseInt(values[0])
	l.ShortName = values[1]
	l.LongName = values[2]
}

func (l *TradeConditionMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.TradeConditionId = proto.ParseInt(values[1])
	l.ShortName = values[2]
	l.LongName = values[3]
}
