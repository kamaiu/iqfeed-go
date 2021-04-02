package update

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"time"
)

func (u *Message) Unmarshal(values []string, fields []*Field) {
	l := len(values)
	if l > len(fields) {
		l = len(fields)
	}
	for i := 0; i < l; i++ {
		fields[i].Parser(u, values[i])
	}
}

type Message struct {
	IncrementalVolume           int64
	BidTick                     int64
	Strike                      float64
	MarketCenter                string
	LastTradeTime               time.Duration
	LastTradeDate               time.Time
	ExtendedTradingLast         float64
	ExpirationDate              time.Time
	RegionalVolume              int64
	NetAssetValue2              float64
	Regions                     string
	TradeMarketCenter           string
	TradeTime                   time.Duration
	SevenDayYield               float64
	Ask                         float64
	AskChange                   float64
	AskMarketCenter             int64
	AskSize                     int64
	AskTime                     time.Duration
	AvailableRegions            string
	AverageMaturity             float64
	Bid                         float64
	BidChange                   float64
	BidMarketCenter             int64
	BidSize                     int64
	BidTime                     time.Duration
	Change                      float64
	ChangeFromOpen              float64
	Close                       float64
	CloseRange1                 float64
	CloseRange2                 float64
	DaystoExpiration            int64
	DecimalPrecision            int64
	Delay                       int64
	ExchangeID                  string
	ExtendedTrade               float64
	ExtendedTradeDate           time.Time
	ExtendedTradeMarketCenter   int64
	ExtendedTradeSize           int64
	ExtendedTradeTime           time.Duration
	ExtendedTradingChange       float64
	ExtendedTradingDifference   float64
	FinancialStatusIndicator    byte
	FractionDisplayCode         string
	High                        float64
	Last                        float64
	LastDate                    time.Time
	LastMarketCenter            int64
	LastSize                    int64
	LastTime                    time.Duration
	Low                         float64
	MarketCapitalization        float64
	MarketOpen                  int64
	MessageContents             string
	MostRecentTrade             float64
	MostRecentTradeConditions   string
	MostRecentTradeDate         time.Time
	MostRecentTradeMarketCenter int64
	MostRecentTradeSize         int64
	MostRecentTradeTime         time.Duration
	NetAssetValue               float64
	NumberofTradesToday         int64
	Open                        float64
	OpenInterest                int64
	OpenRange1                  float64
	OpenRange2                  float64
	PercentChange               float64
	PercentOffAverageVolume     float64
	PreviousDayVolume           int64
	PriceEarningsRatio          float64
	Range                       float64
	RestrictedCode              string
	Settle                      float64
	SettlementDate              time.Time
	Spread                      float64
	Symbol                      string
	Tick                        int64
	TickID                      int64
	TotalVolume                 int64
	Type                        string
	Volatility                  float64
	VWAP                        float64
}

var (
	IncrementalVolume = &Field{
		Name: "Incremental Volume",
		Parser: func(u *Message, val string) {
			u.IncrementalVolume = proto.ParseInt(val)
		},
	}
	BidTick = &Field{
		Name: "Bid Tick",
		Parser: func(u *Message, val string) {
			u.BidTick = proto.ParseInt(val)
		},
	}
	Strike = &Field{
		Name: "Strike",
		Parser: func(u *Message, val string) {
			u.Strike = proto.ParseFloat(val)
		},
	}
	MarketCenter = &Field{
		Name: "Market Center",
		Parser: func(u *Message, val string) {
			u.MarketCenter = val
		},
	}
	LastTradeTime = &Field{
		Name: "Last Trade Time",
		Parser: func(u *Message, val string) {
			u.LastTradeTime = proto.ParseDuration(val)
		},
	}
	LastTradeDate = &Field{
		Name: "Last Trade Date",
		Parser: func(u *Message, val string) {
			u.LastTradeDate = proto.ParseDate(val)
		},
	}
	ExtendedTradingLast = &Field{
		Name: "Extended Trading Last",
		Parser: func(u *Message, val string) {
			u.ExtendedTradingLast = proto.ParseFloat(val)
		},
	}
	ExpirationDate = &Field{
		Name: "Expiration Date",
		Parser: func(u *Message, val string) {
			u.ExpirationDate = proto.ParseDate(val)
		},
	}
	RegionalVolume = &Field{
		Name: "Regional Volume",
		Parser: func(u *Message, val string) {
			u.RegionalVolume = proto.ParseInt(val)
		},
	}
	NetAssetValue2 = &Field{
		Name: "Net Asset Value 2",
		Parser: func(u *Message, val string) {
			u.NetAssetValue2 = proto.ParseFloat(val)
		},
	}
	Regions = &Field{
		Name: "Regions",
		Parser: func(u *Message, val string) {
			u.Regions = val
		},
	}
	TradeMarketCenter = &Field{
		Name: "Trade Market Center",
		Parser: func(u *Message, val string) {
			u.TradeMarketCenter = val
		},
	}
	TradeTime = &Field{
		Name: "Trade Time",
		Parser: func(u *Message, val string) {
			u.TradeTime = proto.ParseDuration(val)
		},
	}
	SevenDayYield = &Field{
		Name: "7 Day Yield",
		Parser: func(u *Message, val string) {
			u.SevenDayYield = proto.ParseFloat(val)
		},
	}
	Ask = &Field{
		Name: "Ask",
		Parser: func(u *Message, val string) {
			u.Ask = proto.ParseFloat(val)
		},
	}
	AskChange = &Field{
		Name: "Ask Change",
		Parser: func(u *Message, val string) {
			u.AskChange = proto.ParseFloat(val)
		},
	}
	AskMarketCenter = &Field{
		Name: "Ask Market Center",
		Parser: func(u *Message, val string) {
			u.AskMarketCenter = proto.ParseInt(val)
		},
	}
	AskSize = &Field{
		Name: "Ask Size",
		Parser: func(u *Message, val string) {
			u.AskSize = proto.ParseInt(val)
		},
	}
	AskTime = &Field{
		Name: "Ask Time",
		Parser: func(u *Message, val string) {
			u.AskTime = proto.ParseDuration(val)
		},
	}
	AvailableRegions = &Field{
		Name: "Available Regions",
		Parser: func(u *Message, val string) {
			u.AvailableRegions = val
		},
	}
	AverageMaturity = &Field{
		Name: "Average Maturity",
		Parser: func(u *Message, val string) {
			u.AverageMaturity = proto.ParseFloat(val)
		},
	}
	Bid = &Field{
		Name: "Bid",
		Parser: func(u *Message, val string) {
			u.Bid = proto.ParseFloat(val)
		},
	}
	BidChange = &Field{
		Name: "Bid Change",
		Parser: func(u *Message, val string) {
			u.BidChange = proto.ParseFloat(val)
		},
	}
	BidMarketCenter = &Field{
		Name: "Bid Market Center",
		Parser: func(u *Message, val string) {
			u.BidMarketCenter = proto.ParseInt(val)
		},
	}
	BidSize = &Field{
		Name: "Bid Size",
		Parser: func(u *Message, val string) {
			u.BidSize = proto.ParseInt(val)
		},
	}
	BidTime = &Field{
		Name: "Bid Time",
		Parser: func(u *Message, val string) {
			u.BidTime = proto.ParseDuration(val)
		},
	}
	Change = &Field{
		Name: "Change",
		Parser: func(u *Message, val string) {
			u.Change = proto.ParseFloat(val)
		},
	}
	ChangeFromOpen = &Field{
		Name: "Change From Open",
		Parser: func(u *Message, val string) {
			u.ChangeFromOpen = proto.ParseFloat(val)
		},
	}
	Close = &Field{
		Name: "Close",
		Parser: func(u *Message, val string) {
			u.Close = proto.ParseFloat(val)
		},
	}
	CloseRange1 = &Field{
		Name: "Close Range 1",
		Parser: func(u *Message, val string) {
			u.CloseRange1 = proto.ParseFloat(val)
		},
	}
	CloseRange2 = &Field{
		Name: "Close Range 2",
		Parser: func(u *Message, val string) {
			u.CloseRange2 = proto.ParseFloat(val)
		},
	}
	DaystoExpiration = &Field{
		Name: "Days to Expiration",
		Parser: func(u *Message, val string) {
			u.DaystoExpiration = proto.ParseInt(val)
		},
	}
	DecimalPrecision = &Field{
		Name: "Decimal Precision",
		Parser: func(u *Message, val string) {
			u.DecimalPrecision = proto.ParseInt(val)
		},
	}
	Delay = &Field{
		Name: "Delay",
		Parser: func(u *Message, val string) {
			u.Delay = proto.ParseInt(val)
		},
	}
	ExchangeID = &Field{
		Name: "Exchange ID",
		Parser: func(u *Message, val string) {
			u.ExchangeID = val
		},
	}
	ExtendedTrade = &Field{
		Name: "Extended Trade",
		Parser: func(u *Message, val string) {
			u.ExtendedTrade = proto.ParseFloat(val)
		},
	}
	ExtendedTradeDate = &Field{
		Name: "Extended Trade Date",
		Parser: func(u *Message, val string) {
			u.ExtendedTradeDate = proto.ParseDate(val)
		},
	}
	ExtendedTradeMarketCenter = &Field{
		Name: "Extended Trade Market Center",
		Parser: func(u *Message, val string) {
			u.ExtendedTradeMarketCenter = proto.ParseInt(val)
		},
	}
	ExtendedTradeSize = &Field{
		Name: "Extended Trade Size",
		Parser: func(u *Message, val string) {
			u.ExtendedTradeSize = proto.ParseInt(val)
		},
	}
	ExtendedTradeTime = &Field{
		Name: "Extended Trade Time",
		Parser: func(u *Message, val string) {
			u.ExtendedTradeTime = proto.ParseDuration(val)
		},
	}
	ExtendedTradingChange = &Field{
		Name: "Extended Trading Change",
		Parser: func(u *Message, val string) {
			u.ExtendedTradingChange = proto.ParseFloat(val)
		},
	}
	ExtendedTradingDifference = &Field{
		Name: "Extended Trading Difference",
		Parser: func(u *Message, val string) {
			u.ExtendedTradingDifference = proto.ParseFloat(val)
		},
	}
	FinancialStatusIndicator = &Field{
		Name: "Financial Status Indicator",
		Parser: func(u *Message, val string) {
			u.FinancialStatusIndicator = proto.CharAt(val, 0)
		},
	}
	FractionDisplayCode = &Field{
		Name: "Fraction Display Code",
		Parser: func(u *Message, val string) {
			u.FractionDisplayCode = val
		},
	}
	High = &Field{
		Name: "High",
		Parser: func(u *Message, val string) {
			u.High = proto.ParseFloat(val)
		},
	}
	Last = &Field{
		Name: "Last",
		Parser: func(u *Message, val string) {
			u.Last = proto.ParseFloat(val)
		},
	}
	LastDate = &Field{
		Name: "Last Date",
		Parser: func(u *Message, val string) {
			u.LastDate = proto.ParseDate(val)
		},
	}
	LastMarketCenter = &Field{
		Name: "Last Market Center",
		Parser: func(u *Message, val string) {
			u.LastMarketCenter = proto.ParseInt(val)
		},
	}
	LastSize = &Field{
		Name: "Last Size",
		Parser: func(u *Message, val string) {
			u.LastSize = proto.ParseInt(val)
		},
	}
	LastTime = &Field{
		Name: "Last Time",
		Parser: func(u *Message, val string) {
			u.LastTime = proto.ParseDuration(val)
		},
	}
	Low = &Field{
		Name: "Low",
		Parser: func(u *Message, val string) {
			u.Low = proto.ParseFloat(val)
		},
	}
	MarketCapitalization = &Field{
		Name: "Market Capitalization",
		Parser: func(u *Message, val string) {
			u.MarketCapitalization = proto.ParseFloat(val)
		},
	}
	MarketOpen = &Field{
		Name: "Market Open",
		Parser: func(u *Message, val string) {
			u.MarketOpen = proto.ParseInt(val)
		},
	}
	MessageContents = &Field{
		Name: "Message Contents",
		Parser: func(u *Message, val string) {
			u.MessageContents = val
		},
	}
	MostRecentTrade = &Field{
		Name: "Most Recent Trade",
		Parser: func(u *Message, val string) {
			u.MostRecentTrade = proto.ParseFloat(val)
		},
	}
	MostRecentTradeConditions = &Field{
		Name: "Most Recent Trade Conditions",
		Parser: func(u *Message, val string) {
			u.MostRecentTradeConditions = val
		},
	}
	MostRecentTradeDate = &Field{
		Name: "Most Recent Trade Date",
		Parser: func(u *Message, val string) {
			u.MostRecentTradeDate = proto.ParseDate(val)
		},
	}
	MostRecentTradeMarketCenter = &Field{
		Name: "Most Recent Trade Market Center",
		Parser: func(u *Message, val string) {
			u.MostRecentTradeMarketCenter = proto.ParseInt(val)
		},
	}
	MostRecentTradeSize = &Field{
		Name: "Most Recent Trade Size",
		Parser: func(u *Message, val string) {
			u.MostRecentTradeSize = proto.ParseInt(val)
		},
	}
	MostRecentTradeTime = &Field{
		Name: "Most Recent Trade Time",
		Parser: func(u *Message, val string) {
			u.MostRecentTradeTime = proto.ParseDuration(val)
		},
	}
	NetAssetValue = &Field{
		Name: "Net Asset Value",
		Parser: func(u *Message, val string) {
			u.NetAssetValue = proto.ParseFloat(val)
		},
	}
	NumberofTradesToday = &Field{
		Name: "Number of Trades Today",
		Parser: func(u *Message, val string) {
			u.NumberofTradesToday = proto.ParseInt(val)
		},
	}
	Open = &Field{
		Name: "Open",
		Parser: func(u *Message, val string) {
			u.Open = proto.ParseFloat(val)
		},
	}
	OpenInterest = &Field{
		Name: "Open Interest",
		Parser: func(u *Message, val string) {
			u.OpenInterest = proto.ParseInt(val)
		},
	}
	OpenRange1 = &Field{
		Name: "Open Range 1",
		Parser: func(u *Message, val string) {
			u.OpenRange1 = proto.ParseFloat(val)
		},
	}
	OpenRange2 = &Field{
		Name: "Open Range 2",
		Parser: func(u *Message, val string) {
			u.OpenRange2 = proto.ParseFloat(val)
		},
	}
	PercentChange = &Field{
		Name: "Percent Change",
		Parser: func(u *Message, val string) {
			u.PercentChange = proto.ParseFloat(val)
		},
	}
	PercentOffAverageVolume = &Field{
		Name: "Percent Off Average Volume",
		Parser: func(u *Message, val string) {
			u.PercentOffAverageVolume = proto.ParseFloat(val)
		},
	}
	PreviousDayVolume = &Field{
		Name: "Previous Day Volume",
		Parser: func(u *Message, val string) {
			u.PreviousDayVolume = proto.ParseInt(val)
		},
	}
	PriceEarningsRatio = &Field{
		Name: "Price-Earnings Ratio",
		Parser: func(u *Message, val string) {
			u.PriceEarningsRatio = proto.ParseFloat(val)
		},
	}
	Range = &Field{
		Name: "Range",
		Parser: func(u *Message, val string) {
			u.Range = proto.ParseFloat(val)
		},
	}
	RestrictedCode = &Field{
		Name: "Restricted Code",
		Parser: func(u *Message, val string) {
			u.RestrictedCode = val
		},
	}
	Settle = &Field{
		Name: "Settle",
		Parser: func(u *Message, val string) {
			u.Settle = proto.ParseFloat(val)
		},
	}
	SettlementDate = &Field{
		Name: "Settlement Date",
		Parser: func(u *Message, val string) {
			u.SettlementDate = proto.ParseDate(val)
		},
	}
	Spread = &Field{
		Name: "Spread",
		Parser: func(u *Message, val string) {
			u.Spread = proto.ParseFloat(val)
		},
	}
	Symbol = &Field{
		Name: "Symbol",
		Parser: func(u *Message, val string) {
			u.Symbol = val
		},
	}
	Tick = &Field{
		Name: "Tick",
		Parser: func(u *Message, val string) {
			u.Tick = proto.ParseInt(val)
		},
	}
	TickID = &Field{
		Name: "TickID",
		Parser: func(u *Message, val string) {
			u.TickID = proto.ParseInt(val)
		},
	}
	TotalVolume = &Field{
		Name: "Total Volume",
		Parser: func(u *Message, val string) {
			u.TotalVolume = proto.ParseInt(val)
		},
	}
	Type = &Field{
		Name: "Type",
		Parser: func(u *Message, val string) {
			u.Type = val
		},
	}
	Volatility = &Field{
		Name: "Volatility",
		Parser: func(u *Message, val string) {
			u.Volatility = proto.ParseFloat(val)
		},
	}
	VWAP = &Field{
		Name: "VWAP",
		Parser: func(u *Message, val string) {
			u.VWAP = proto.ParseFloat(val)
		},
	}
	byName = map[string]*Field{"Incremental Volume": IncrementalVolume,
		"Bid Tick":                        BidTick,
		"Strike":                          Strike,
		"Market Center":                   MarketCenter,
		"Last Trade Time":                 LastTradeTime,
		"Last Trade Date":                 LastTradeDate,
		"Extended Trading Last":           ExtendedTradingLast,
		"Expiration Date":                 ExpirationDate,
		"Regional Volume":                 RegionalVolume,
		"Net Asset Value 2":               NetAssetValue2,
		"Regions":                         Regions,
		"Trade Market Center":             TradeMarketCenter,
		"Trade Time":                      TradeTime,
		"7 Day Yield":                     SevenDayYield,
		"Ask":                             Ask,
		"Ask Change":                      AskChange,
		"Ask Market Center":               AskMarketCenter,
		"Ask Size":                        AskSize,
		"Ask Time":                        AskTime,
		"Available Regions":               AvailableRegions,
		"Average Maturity":                AverageMaturity,
		"Bid":                             Bid,
		"Bid Change":                      BidChange,
		"Bid Market Center":               BidMarketCenter,
		"Bid Size":                        BidSize,
		"Bid Time":                        BidTime,
		"Change":                          Change,
		"Change From Open":                ChangeFromOpen,
		"Close":                           Close,
		"Close Range 1":                   CloseRange1,
		"Close Range 2":                   CloseRange2,
		"Days to Expiration":              DaystoExpiration,
		"Decimal Precision":               DecimalPrecision,
		"Delay":                           Delay,
		"Exchange ID":                     ExchangeID,
		"Extended Trade":                  ExtendedTrade,
		"Extended Trade Date":             ExtendedTradeDate,
		"Extended Trade Market Center":    ExtendedTradeMarketCenter,
		"Extended Trade Size":             ExtendedTradeSize,
		"Extended Trade Time":             ExtendedTradeTime,
		"Extended Trading Change":         ExtendedTradingChange,
		"Extended Trading Difference":     ExtendedTradingDifference,
		"Financial Status Indicator":      FinancialStatusIndicator,
		"Fraction Display Code":           FractionDisplayCode,
		"High":                            High,
		"Last":                            Last,
		"Last Date":                       LastDate,
		"Last Market Center":              LastMarketCenter,
		"Last Size":                       LastSize,
		"Last Time":                       LastTime,
		"Low":                             Low,
		"Market Capitalization":           MarketCapitalization,
		"Market Open":                     MarketOpen,
		"Message Contents":                MessageContents,
		"Most Recent Trade":               MostRecentTrade,
		"Most Recent Trade Conditions":    MostRecentTradeConditions,
		"Most Recent Trade Date":          MostRecentTradeDate,
		"Most Recent Trade Market Center": MostRecentTradeMarketCenter,
		"Most Recent Trade Size":          MostRecentTradeSize,
		"Most Recent Trade Time":          MostRecentTradeTime,
		"Net Asset Value":                 NetAssetValue,
		"Number of Trades Today":          NumberofTradesToday,
		"Open":                            Open,
		"Open Interest":                   OpenInterest,
		"Open Range 1":                    OpenRange1,
		"Open Range 2":                    OpenRange2,
		"Percent Change":                  PercentChange,
		"Percent Off Average Volume":      PercentOffAverageVolume,
		"Previous Day Volume":             PreviousDayVolume,
		"Price-Earnings Ratio":            PriceEarningsRatio,
		"Range":                           Range,
		"Restricted Code":                 RestrictedCode,
		"Settle":                          Settle,
		"Settlement Date":                 SettlementDate,
		"Spread":                          Spread,
		"Symbol":                          Symbol,
		"Tick":                            Tick,
		"TickID":                          TickID,
		"Total Volume":                    TotalVolume,
		"Type":                            Type,
		"Volatility":                      Volatility,
		"VWAP":                            VWAP,
	}
)

func Find(name string) *Field { return byName[name] }
