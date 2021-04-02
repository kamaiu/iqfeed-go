package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type MarketSymbol struct {
	Symbol       string
	Description  string
	Exchange     Exchange
	ListedMarket ListedExchange
	SecurityType SecurityType
	Sic          string
	FrontMonth   string
	Naics        string
}

func (m *MarketSymbol) Unmarshal(values []string) {
	if len(values) < 8 {
		return
	}
	m.Symbol = values[0]
	m.Description = values[1]
	m.Exchange = (Exchange)(values[2])
	m.ListedMarket = (ListedExchange)(values[3])
	m.SecurityType = (SecurityType)(values[4])
	m.Sic = values[5]
	m.FrontMonth = values[6]
	m.Naics = values[7]
}

type SymbolByFilterMessage struct {
	Symbol         string
	ListedMarketId int64
	SecurityTypeId int64
	Description    string
	RequestId      string
}

func (l *SymbolByFilterMessage) Unmarshal(values []string) {
	l.Symbol = values[0]
	l.ListedMarketId = proto.ParseInt(values[1])
	l.SecurityTypeId = proto.ParseInt(values[2])
	l.Description = values[3]
}

func (l *SymbolByFilterMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.Symbol = values[1]
	l.ListedMarketId = proto.ParseInt(values[2])
	l.SecurityTypeId = proto.ParseInt(values[3])
	l.Description = values[4]
}

type SymbolByNaicsCodeMessage struct {
	NaicsCode      int64
	Symbol         string
	ListedMarketId int64
	SecurityTypeId int64
	Description    string
	RequestId      string
}

func (l *SymbolByNaicsCodeMessage) Unmarshal(values []string) {
	l.NaicsCode = proto.ParseInt(values[0])
	l.Symbol = values[1]
	l.ListedMarketId = proto.ParseInt(values[2])
	l.SecurityTypeId = proto.ParseInt(values[3])
	l.Description = values[4]
}

func (l *SymbolByNaicsCodeMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.NaicsCode = proto.ParseInt(values[1])
	l.Symbol = values[2]
	l.ListedMarketId = proto.ParseInt(values[3])
	l.SecurityTypeId = proto.ParseInt(values[4])
	l.Description = values[5]
}

type SymbolBySicCodeMessage struct {
	SicCode        int64
	Symbol         string
	ListedMarketId int64
	SecurityTypeId int64
	Description    string
	RequestId      string
}

func (l *SymbolBySicCodeMessage) Unmarshal(values []string) {
	l.SicCode = proto.ParseInt(values[0])
	l.Symbol = values[1]
	l.ListedMarketId = proto.ParseInt(values[2])
	l.SecurityTypeId = proto.ParseInt(values[3])
	l.Description = values[4]
}

func (l *SymbolBySicCodeMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.SicCode = proto.ParseInt(values[1])
	l.Symbol = values[2]
	l.ListedMarketId = proto.ParseInt(values[3])
	l.SecurityTypeId = proto.ParseInt(values[4])
	l.Description = values[5]
}
