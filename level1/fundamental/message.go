package fundamental

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"strconv"
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
	Symbol                      string
	ExchangeID                  string
	PE                          float64
	AverageVolume               int64
	Five2WeekHigh               float64
	Five2WeekLow                float64
	CalendarYearHigh            float64
	CalendarYearLow             float64
	DividendYield               float64
	DividendAmount              float64
	DividendRate                float64
	PayDate                     time.Time
	ExdividendDate              time.Time
	ShortInterest               int64
	CurrentYearEPS              float64
	NextYearEPS                 float64
	FiveyearGrowthPercentage    float64
	FiscalYearEnd               int64
	CompanyName                 string
	RootOptionSymbol            string
	PercentHeldByInstitutions   float64
	Beta                        float64
	Leaps                       string
	CurrentAssets               float64
	CurrentLiabilities          float64
	BalanceSheetDate            time.Time
	LongtermDebt                float64
	CommonSharesOutstanding     float64
	SplitFactor1                string
	SplitFactor2                string
	MarketCenter                string
	FormatCode                  string
	Precision                   int64
	SIC                         int64
	HistoricalVolatility        float64
	SecurityType                string
	ListedMarket                string
	Five2WeekHighDate           time.Time
	Five2WeekLowDate            time.Time
	CalendarYearHighDate        time.Time
	CalendarYearLowDate         time.Time
	YearEndClose                float64
	MaturityDate                time.Time
	CouponRate                  float64
	ExpirationDate              time.Time
	StrikePrice                 float64
	NAICS                       int64
	ExchangeRoot                string
	OptionsPremiumMultiplier    float64
	OptionsMultipleDeliverables int64
}

func (t *Message) ToCSV(o []byte) []byte {
	o = append(o, 'F')
	o = append(o, ',')
	o = append(o, t.Symbol...)
	o = append(o, ',')
	o = append(o, t.ExchangeID...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.PE, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.AverageVolume, 10)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.Five2WeekHigh, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.Five2WeekLow, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CalendarYearHigh, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CalendarYearLow, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.DividendYield, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.DividendAmount, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.DividendRate, 'f', 4, 64)...)
	o = append(o, ',')
	o = proto.AppendDate(o, t.PayDate)
	o = append(o, ',')
	o = proto.AppendDate(o, t.ExdividendDate)
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.ShortInterest, 10)...)
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CurrentYearEPS, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.NextYearEPS, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.FiveyearGrowthPercentage, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.FiscalYearEnd, 10)...)
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, t.CompanyName...)
	o = append(o, ',')
	o = append(o, t.RootOptionSymbol...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.PercentHeldByInstitutions, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.Beta, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, t.Leaps...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CurrentAssets, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CurrentLiabilities, 'f', 4, 64)...)
	o = append(o, ',')
	o = proto.AppendDate(o, t.BalanceSheetDate)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.LongtermDebt, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CommonSharesOutstanding, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, t.SplitFactor1...)
	o = append(o, ',')
	o = append(o, t.SplitFactor2...)
	o = append(o, ',')
	o = append(o, ',')
	o = append(o, t.MarketCenter...)
	o = append(o, ',')
	o = append(o, t.FormatCode...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.Precision, 10)...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.SIC, 10)...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.HistoricalVolatility, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, t.SecurityType...)
	o = append(o, ',')
	o = append(o, t.ListedMarket...)
	o = append(o, ',')
	o = proto.AppendDate(o, t.Five2WeekHighDate)
	o = append(o, ',')
	o = proto.AppendDate(o, t.Five2WeekLowDate)
	o = append(o, ',')
	o = proto.AppendDate(o, t.CalendarYearHighDate)
	o = append(o, ',')
	o = proto.AppendDate(o, t.CalendarYearLowDate)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.YearEndClose, 'f', 4, 64)...)
	o = append(o, ',')
	o = proto.AppendDate(o, t.MaturityDate)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.CouponRate, 'f', 4, 64)...)
	o = append(o, ',')
	o = proto.AppendDate(o, t.ExpirationDate)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.StrikePrice, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.NAICS, 10)...)
	o = append(o, ',')
	o = append(o, t.ExchangeRoot...)
	o = append(o, ',')
	o = append(o, strconv.FormatFloat(t.OptionsPremiumMultiplier, 'f', 4, 64)...)
	o = append(o, ',')
	o = append(o, strconv.FormatInt(t.OptionsMultipleDeliverables, 10)...)
	return o
}

var (
	Symbol = &Field{
		Name: "Symbol",
		Parser: func(t *Message, value string) {
			t.Symbol = value
		},
	}
	ExchangeID = &Field{
		Name: "Exchange ID",
		Parser: func(t *Message, value string) {
			t.ExchangeID = value
		},
	}
	PE = &Field{
		Name: "PE",
		Parser: func(t *Message, value string) {
			t.PE = proto.ParseFloat(value)
		},
	}
	AverageVolume = &Field{
		Name: "Average Volume",
		Parser: func(t *Message, value string) {
			t.AverageVolume = proto.ParseInt(value)
		},
	}
	Five2WeekHigh = &Field{
		Name: "52 Week High",
		Parser: func(t *Message, value string) {
			t.Five2WeekHigh = proto.ParseFloat(value)
		},
	}
	Five2WeekLow = &Field{
		Name: "52 Week Low",
		Parser: func(t *Message, value string) {
			t.Five2WeekLow = proto.ParseFloat(value)
		},
	}
	CalendarYearHigh = &Field{
		Name: "Calendar Year High",
		Parser: func(t *Message, value string) {
			t.CalendarYearHigh = proto.ParseFloat(value)
		},
	}
	CalendarYearLow = &Field{
		Name: "Calendar Year Low",
		Parser: func(t *Message, value string) {
			t.CalendarYearLow = proto.ParseFloat(value)
		},
	}
	DividendYield = &Field{
		Name: "Dividend Yield",
		Parser: func(t *Message, value string) {
			t.DividendYield = proto.ParseFloat(value)
		},
	}
	DividendAmount = &Field{
		Name: "Dividend Amount",
		Parser: func(t *Message, value string) {
			t.DividendAmount = proto.ParseFloat(value)
		},
	}
	DividendRate = &Field{
		Name: "Dividend Rate",
		Parser: func(t *Message, value string) {
			t.DividendRate = proto.ParseFloat(value)
		},
	}
	PayDate = &Field{
		Name: "Pay Date",
		Parser: func(t *Message, value string) {
			t.PayDate = proto.ParseDate(value)
		},
	}
	ExdividendDate = &Field{
		Name: "Ex-dividend Date",
		Parser: func(t *Message, value string) {
			t.ExdividendDate = proto.ParseDate(value)
		},
	}
	ShortInterest = &Field{
		Name: "Short Interest",
		Parser: func(t *Message, value string) {
			t.ShortInterest = proto.ParseInt(value)
		},
	}
	CurrentYearEPS = &Field{
		Name: "Current Year EPS",
		Parser: func(t *Message, value string) {
			t.CurrentYearEPS = proto.ParseFloat(value)
		},
	}
	NextYearEPS = &Field{
		Name: "Next Year EPS",
		Parser: func(t *Message, value string) {
			t.NextYearEPS = proto.ParseFloat(value)
		},
	}
	FiveyearGrowthPercentage = &Field{
		Name: "Five-year Growth Percentage",
		Parser: func(t *Message, value string) {
			t.FiveyearGrowthPercentage = proto.ParseFloat(value)
		},
	}
	FiscalYearEnd = &Field{
		Name: "Fiscal Year End",
		Parser: func(t *Message, value string) {
			t.FiscalYearEnd = proto.ParseInt(value)
		},
	}
	CompanyName = &Field{
		Name: "Company Name",
		Parser: func(t *Message, value string) {
			t.CompanyName = value
		},
	}
	RootOptionSymbol = &Field{
		Name: "Root Option Symbol",
		Parser: func(t *Message, value string) {
			t.RootOptionSymbol = value
		},
	}
	PercentHeldByInstitutions = &Field{
		Name: "Percent Held By Institutions",
		Parser: func(t *Message, value string) {
			t.PercentHeldByInstitutions = proto.ParseFloat(value)
		},
	}
	Beta = &Field{
		Name: "Beta",
		Parser: func(t *Message, value string) {
			t.Beta = proto.ParseFloat(value)
		},
	}
	Leaps = &Field{
		Name: "Leaps",
		Parser: func(t *Message, value string) {
			t.Leaps = value
		},
	}
	CurrentAssets = &Field{
		Name: "Current Assets",
		Parser: func(t *Message, value string) {
			t.CurrentAssets = proto.ParseFloat(value)
		},
	}
	CurrentLiabilities = &Field{
		Name: "Current Liabilities",
		Parser: func(t *Message, value string) {
			t.CurrentLiabilities = proto.ParseFloat(value)
		},
	}
	BalanceSheetDate = &Field{
		Name: "Balance Sheet Date",
		Parser: func(t *Message, value string) {
			t.BalanceSheetDate = proto.ParseDate(value)
		},
	}
	LongtermDebt = &Field{
		Name: "Long-term Debt",
		Parser: func(t *Message, value string) {
			t.LongtermDebt = proto.ParseFloat(value)
		},
	}
	CommonSharesOutstanding = &Field{
		Name: "Common Shares Outstanding",
		Parser: func(t *Message, value string) {
			t.CommonSharesOutstanding = proto.ParseFloat(value)
		},
	}
	SplitFactor1 = &Field{
		Name: "Split Factor 1",
		Parser: func(t *Message, value string) {
			t.SplitFactor1 = value
		},
	}
	SplitFactor2 = &Field{
		Name: "Split Factor 2",
		Parser: func(t *Message, value string) {
			t.SplitFactor2 = value
		},
	}
	MarketCenter = &Field{
		Name: "Market Center",
		Parser: func(t *Message, value string) {
			t.MarketCenter = value
		},
	}
	FormatCode = &Field{
		Name: "Format Code",
		Parser: func(t *Message, value string) {
			t.FormatCode = value
		},
	}
	Precision = &Field{
		Name: "Precision",
		Parser: func(t *Message, value string) {
			t.Precision = proto.ParseInt(value)
		},
	}
	SIC = &Field{
		Name: "SIC",
		Parser: func(t *Message, value string) {
			t.SIC = proto.ParseInt(value)
		},
	}
	HistoricalVolatility = &Field{
		Name: "Historical Volatility",
		Parser: func(t *Message, value string) {
			t.HistoricalVolatility = proto.ParseFloat(value)
		},
	}
	SecurityType = &Field{
		Name: "Security Type",
		Parser: func(t *Message, value string) {
			t.SecurityType = value
		},
	}
	ListedMarket = &Field{
		Name: "Listed Market",
		Parser: func(t *Message, value string) {
			t.ListedMarket = value
		},
	}
	Five2WeekHighDate = &Field{
		Name: "52 Week High Date",
		Parser: func(t *Message, value string) {
			t.Five2WeekHighDate = proto.ParseDate(value)
		},
	}
	Five2WeekLowDate = &Field{
		Name: "52 Week Low Date",
		Parser: func(t *Message, value string) {
			t.Five2WeekLowDate = proto.ParseDate(value)
		},
	}
	CalendarYearHighDate = &Field{
		Name: "Calendar Year High Date",
		Parser: func(t *Message, value string) {
			t.CalendarYearHighDate = proto.ParseDate(value)
		},
	}
	CalendarYearLowDate = &Field{
		Name: "Calendar Year Low Date",
		Parser: func(t *Message, value string) {
			t.CalendarYearLowDate = proto.ParseDate(value)
		},
	}
	YearEndClose = &Field{
		Name: "Year End Close",
		Parser: func(t *Message, value string) {
			t.YearEndClose = proto.ParseFloat(value)
		},
	}
	MaturityDate = &Field{
		Name: "Maturity Date",
		Parser: func(t *Message, value string) {
			t.MaturityDate = proto.ParseDate(value)
		},
	}
	CouponRate = &Field{
		Name: "Coupon Rate",
		Parser: func(t *Message, value string) {
			t.CouponRate = proto.ParseFloat(value)
		},
	}
	ExpirationDate = &Field{
		Name: "Expiration Date",
		Parser: func(t *Message, value string) {
			t.ExpirationDate = proto.ParseDate(value)
		},
	}
	StrikePrice = &Field{
		Name: "Strike Price",
		Parser: func(t *Message, value string) {
			t.StrikePrice = proto.ParseFloat(value)
		},
	}
	NAICS = &Field{
		Name: "NAICS",
		Parser: func(t *Message, value string) {
			t.NAICS = proto.ParseInt(value)
		},
	}
	ExchangeRoot = &Field{
		Name: "Exchange Root",
		Parser: func(t *Message, value string) {
			t.ExchangeRoot = value
		},
	}
	OptionsPremiumMultiplier = &Field{
		Name: "Options Premium Multiplier",
		Parser: func(t *Message, value string) {
			t.OptionsPremiumMultiplier = proto.ParseFloat(value)
		},
	}
	OptionsMultipleDeliverables = &Field{
		Name: "Options Multiple Deliverables",
		Parser: func(t *Message, value string) {
			t.OptionsMultipleDeliverables = proto.ParseInt(value)
		},
	}
	byName = map[string]*Field{
		"Symbol":                        Symbol,
		"Exchange ID":                   ExchangeID,
		"PE":                            PE,
		"Average Volume":                AverageVolume,
		"52 Week High":                  Five2WeekHigh,
		"52 Week Low":                   Five2WeekLow,
		"Calendar Year High":            CalendarYearHigh,
		"Calendar Year Low":             CalendarYearLow,
		"Dividend Yield":                DividendYield,
		"Dividend Amount":               DividendAmount,
		"Dividend Rate":                 DividendRate,
		"Pay Date":                      PayDate,
		"Ex-dividend Date":              ExdividendDate,
		"Short Interest":                ShortInterest,
		"Current Year EPS":              CurrentYearEPS,
		"Next Year EPS":                 NextYearEPS,
		"Five-year Growth Percentage":   FiveyearGrowthPercentage,
		"Fiscal Year End":               FiscalYearEnd,
		"Company Name":                  CompanyName,
		"Root Option Symbol":            RootOptionSymbol,
		"Percent Held By Institutions":  PercentHeldByInstitutions,
		"Beta":                          Beta,
		"Leaps":                         Leaps,
		"Current Assets":                CurrentAssets,
		"Current Liabilities":           CurrentLiabilities,
		"Balance Sheet Date":            BalanceSheetDate,
		"Long-term Debt":                LongtermDebt,
		"Common Shares Outstanding":     CommonSharesOutstanding,
		"Split Factor 1":                SplitFactor1,
		"Split Factor 2":                SplitFactor2,
		"Market Center":                 MarketCenter,
		"Format Code":                   FormatCode,
		"Precision":                     Precision,
		"SIC":                           SIC,
		"Historical Volatility":         HistoricalVolatility,
		"Security Type":                 SecurityType,
		"Listed Market":                 ListedMarket,
		"52 Week High Date":             Five2WeekHighDate,
		"52 Week Low Date":              Five2WeekLowDate,
		"Calendar Year High Date":       CalendarYearHighDate,
		"Calendar Year Low Date":        CalendarYearLowDate,
		"Year End Close":                YearEndClose,
		"Maturity Date":                 MaturityDate,
		"Coupon Rate":                   CouponRate,
		"Expiration Date":               ExpirationDate,
		"Strike Price":                  StrikePrice,
		"NAICS":                         NAICS,
		"Exchange Root":                 ExchangeRoot,
		"Options Premium Multiplier":    OptionsPremiumMultiplier,
		"Options Multiple Deliverables": OptionsMultipleDeliverables,
	}
)

func Find(name string) *Field { return byName[name] }
