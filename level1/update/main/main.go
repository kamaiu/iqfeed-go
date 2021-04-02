package main

import (
	"fmt"
	"github.com/kamaiu/iqfeed-go/proto"
	"strings"
)

const unknownFields = `
Incremental Volume
Bid Tick
Last Trade Time
Strike
Market Center
Last Trade Date
(Reserved)
Extended Trading Last
Expiration Date
Regional Volume
Net Asset Value 2
Regions
Trade Market Center
Trade Time
`

func main() {
	d := `
	[FieldsetDescription("Incremental Volume", typeof(int))]
	IncrementalVolume

	[FieldsetDescription("Bid Tick", typeof(int))]
	BidTick

	[FieldsetDescription("Strike", typeof(double))]
	Strike

	[FieldsetDescription("Market Center", typeof(string))]
	MarketCenter

	[FieldsetDescription("Last Trade Time", typeof(TimeSpan))]
	LastTradeTime

	[FieldsetDescription("Last Trade Date", typeof(DateTime))]
	LastTradeDate

	[FieldsetDescription("Extended Trading Last", typeof(double))]
	ExtendedTradingLast

	[FieldsetDescription("Expiration Date", typeof(DateTime))]
	ExpirationDate

	[FieldsetDescription("Regional Volume", typeof(int))]
	RegionalVolume

	[FieldsetDescription("Net Asset Value 2", typeof(double))]
	NetAssetValue2

	[FieldsetDescription("Regions", typeof(string))]
	Regions

	[FieldsetDescription("Trade Market Center", typeof(string))]
	TradeMarketCenter

	[FieldsetDescription("Trade Time", typeof(TimeSpan))]
	TradeTime

	[FieldsetDescription("7 Day Yield", typeof(double))]
    SevenDayYield,

   [FieldsetDescription("Ask", typeof(double))]
   Ask,

   [FieldsetDescription("Ask Change", typeof(double))]
   AskChange,

   [FieldsetDescription("Ask Market Center", typeof(int))]
   AskMarketCenter,

   [FieldsetDescription("Ask Size", typeof(int))]
   AskSize,

   [FieldsetDescription("Ask Time", typeof(TimeSpan))]
   AskTime,

   [FieldsetDescription("Available Regions", typeof(string))]
   AvailableRegions,

   [FieldsetDescription("Average Maturity", typeof(double))]
   AverageMaturity,

   [FieldsetDescription("Bid", typeof(double))]
   Bid,

   [FieldsetDescription("Bid Change", typeof(double))]
   BidChange,

   [FieldsetDescription("Bid Market Center", typeof(int))]
   BidMarketCenter,

   [FieldsetDescription("Bid Size", typeof(int))]
   BidSize,

   [FieldsetDescription("Bid Time", typeof(TimeSpan))]
   BidTime,

   [FieldsetDescription("Change", typeof(double))]
   Change,

   [FieldsetDescription("Change From Open", typeof(double))]
   ChangeFromOpen,

   [FieldsetDescription("Close", typeof(double))]
   Close,

   [FieldsetDescription("Close Range 1", typeof(double))]
   CloseRange1,

   [FieldsetDescription("Close Range 2", typeof(double))]
   CloseRange2,

   [FieldsetDescription("Days to Expiration", typeof(int))]
   DaysToExpiration,

   [FieldsetDescription("Decimal Precision", typeof(int))]
   DecimalPrecision,

   [FieldsetDescription("Delay", typeof(int))]
   Delay,

   [FieldsetDescription("Exchange ID", typeof(string))]
   ExchangeID,

   [FieldsetDescription("Extended Trade", typeof(double))]
   ExtendedTrade,

   [FieldsetDescription("Extended Trade Date", typeof(DateTime))]
   ExtendedTradeDate,

   [FieldsetDescription("Extended Trade Market Center", typeof(int))]
   ExtendedTradeMarketCenter,

   [FieldsetDescription("Extended Trade Size", typeof(int))]
   ExtendedTradeSize,

   [FieldsetDescription("Extended Trade Time", typeof(TimeSpan))]
   ExtendedTradeTime,

   [FieldsetDescription("Extended Trading Change", typeof(double))]
   ExtendedTradingChange,

   [FieldsetDescription("Extended Trading Difference", typeof(double))]
   ExtendedTradingDifference,

   [FieldsetDescription("Financial Status Indicator", typeof(char))]
   FinancialStatusIndicator,

   [FieldsetDescription("Fraction Display Code", typeof(string))]
   FractionDisplayCode,

   [FieldsetDescription("High", typeof(double))]
   High,

   [FieldsetDescription("Last", typeof(double))]
   Last,

   [FieldsetDescription("Last Date", typeof(DateTime))]
   LastDate,

   [FieldsetDescription("Last Market Center", typeof(int))]
   LastMarketCenter,

   [FieldsetDescription("Last Size", typeof(int))]
   LastSize,

   [FieldsetDescription("Last Time", typeof(TimeSpan))]
   LastTime,

   [FieldsetDescription("Low", typeof(double))]
   Low,

   [FieldsetDescription("Market Capitalization", typeof(double))]
   MarketCapitalization,

   [FieldsetDescription("Market Open", typeof(int))]
   MarketOpen,

   [FieldsetDescription("Message Contents", typeof(string))]
   MessageContents,

   [FieldsetDescription("Most Recent Trade", typeof(double))]
   MostRecentTrade,

   [FieldsetDescription("Most Recent Trade Conditions", typeof(string))]
   MostRecentTradeConditions,

   [FieldsetDescription("Most Recent Trade Date", typeof(DateTime))]
   MostRecentTradeDate,

   [FieldsetDescription("Most Recent Trade Market Center", typeof(int))]
   MostRecentTradeMarketCenter,

   [FieldsetDescription("Most Recent Trade Size", typeof(int))]
   MostRecentTradeSize,

   [FieldsetDescription("Most Recent Trade Time", typeof(TimeSpan))]
   MostRecentTradeTime,

   [FieldsetDescription("Net Asset Value", typeof(double))]
   NetAssetValue,

   [FieldsetDescription("Number of Trades Today", typeof(int))]
   NumberOfTradesToday,

   [FieldsetDescription("Open", typeof(double))]
   Open,

   [FieldsetDescription("Open Interest", typeof(int))]
   OpenInterest,

   [FieldsetDescription("Open Range 1", typeof(double))]
   OpenRange1,

   [FieldsetDescription("Open Range 2", typeof(double))]
   OpenRange2,

   [FieldsetDescription("Percent Change", typeof(double))]
   PercentChange,

   [FieldsetDescription("Percent Off Average Volume", typeof(double))]
   PercentOffAverageVolume,

   [FieldsetDescription("Previous Day Volume", typeof(int))]
   PreviousDayVolume,

   [FieldsetDescription("Price-Earnings Ratio", typeof(double))]
   PriceEarningsRatio,

   [FieldsetDescription("Range", typeof(double))]
   Range,

   [FieldsetDescription("Restricted Code", typeof(string))]
   RestrictedCode,

   [FieldsetDescription("Settle", typeof(double))]
   Settle,

   [FieldsetDescription("Settlement Date", typeof(DateTime))]
   SettlementDate,

   [FieldsetDescription("Spread", typeof(double))]
   Spread,

   [FieldsetDescription("Symbol", typeof(string))]
   Symbol,

   [FieldsetDescription("Tick", typeof(int))]
   Tick,

   [FieldsetDescription("TickID", typeof(int))]
   TickID,

   [FieldsetDescription("Total Volume", typeof(int))]
   TotalVolume,

   [FieldsetDescription("Type", typeof(string))]
   Type,

   [FieldsetDescription("Volatility", typeof(double))]
   Volatility,

   [FieldsetDescription("VWAP", typeof(double))]
   VWAP,`

	lines := strings.Split(d, "\n")

	type Field struct {
		GoName string
		Name   string
		Code   byte
	}

	fields := make([]*Field, 0, 64)
	fieldsMap := make(map[string]*Field)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line[0] != '[' {
			continue
		}

		index := strings.Index(line, "\"")
		if index < 0 {
			continue
		}

		line = line[index+1:]
		index = strings.Index(line, "\"")

		name := strings.TrimSpace(line[0:index])
		line = line[index+1:]
		index = strings.Index(line, "(")

		line = line[index+1:]
		index = strings.Index(line, ")")
		kind := line[0:index]

		kindCode := byte('U')
		switch kind {
		case "double":
			kindCode = 'F'
		case "int":
			kindCode = 'I'
		case "string":
			kindCode = 'S'
		case "DateTime":
			kindCode = 'T'
		case "TimeSpan":
			kindCode = 'D'
		case "char":
			kindCode = 'C'

		default:
			panic("unknown kind: " + kind)
		}

		field := &Field{
			GoName: toGoName(name),
			Name:   name,
			Code:   kindCode,
		}
		fields = append(fields, field)

		fieldsMap[name] = field
	}

	sortedNames := strings.Split("Symbol,Most Recent Trade,Most Recent Trade Size,Most Recent Trade Time,Most Recent Trade Market Center,Total Volume,Bid,Bid Size,Ask,Ask Size,Open,High,Low,Close,Message Contents,Most Recent Trade Conditions", ",")

	sortedFields := make([]*Field, 0, len(sortedNames))
	for _, name := range sortedNames {
		field := fieldsMap[name]
		if field == nil {
			panic("not found: " + name)
		}
		sortedFields = append(sortedFields, field)
	}

	b := &strings.Builder{}

	_, _ = b.WriteString("type Message struct {\n")
	for _, field := range fields {
		b.WriteString("    " + field.GoName)
		switch field.Code {
		case 'F':
			_, _ = b.WriteString(" float64\n")
		case 'I':
			_, _ = b.WriteString(" int64\n")
		case 'S':
			_, _ = b.WriteString(" string\n")
		case 'T':
			_, _ = b.WriteString(" time.Time\n")
		case 'D':
			_, _ = b.WriteString(" time.Duration\n")
		case 'C':
			_, _ = b.WriteString(" byte\n")
		}
	}
	_, _ = b.WriteString("}\n\n")
	b.WriteString("var (\n")
	for _, field := range fields {
		b.WriteString("    " + field.GoName + " = &Field{\n")
		b.WriteString("        Name: \"")
		b.WriteString(field.Name)
		b.WriteString("\",\n")
		b.WriteString("        Parser: func (u *Message, val string) {\n")
		b.WriteString("            u." + field.GoName + " = ")
		switch field.Code {
		case 'F':
			_, _ = b.WriteString(" proto.ParseFloat(val)\n")
		case 'I':
			_, _ = b.WriteString(" proto.ParseInt(val)\n")
		case 'S':
			_, _ = b.WriteString(" val\n")
		case 'T':
			_, _ = b.WriteString(" proto.ParseDate(val)\n")
		case 'D':
			_, _ = b.WriteString(" proto.ParseDuration(val)\n")
		case 'C':
			_, _ = b.WriteString(" proto.CharAt(val, 0)\n")
		}
		b.WriteString("    },\n}\n")
	}

	_, _ = b.WriteString("    byName = map[string]*Field{")
	for _, field := range fields {
		b.WriteString(fmt.Sprintf("        \"%s\": %s,\n", field.Name, field.GoName))
	}
	_, _ = b.WriteString("    }\n\n")
	b.WriteString(")\n")

	b.WriteString("func Find(name string) *Field { return byName[name] }\n\n")

	_, _ = b.WriteString("type UpdateSummaryMessage struct {\n")
	for _, field := range sortedFields {
		b.WriteString("    " + field.GoName)
		switch field.Code {
		case 'F':
			_, _ = b.WriteString(" float64\n")
		case 'I':
			_, _ = b.WriteString(" int64\n")
		case 'S':
			_, _ = b.WriteString(" string\n")
		case 'T':
			_, _ = b.WriteString(" time.Time\n")
		case 'D':
			_, _ = b.WriteString(" time.Duration\n")
		case 'C':
			_, _ = b.WriteString(" byte\n")
		}
	}
	_, _ = b.WriteString("}\n")

	b.WriteString("func (u *UpdateSummaryMessage) Unmarshal(values []string) {\n")
	for i, field := range sortedFields {
		index := i + 1
		b.WriteString("    u." + field.GoName + " = ")
		switch field.Code {
		case 'F':
			_, _ = b.WriteString(fmt.Sprintf("proto.ParseFloat(proto.Value(values, %d))\n", index))
		case 'I':
			_, _ = b.WriteString(fmt.Sprintf("proto.ParseInt(proto.Value(values, %d))\n", index))
		case 'S':
			_, _ = b.WriteString(fmt.Sprintf("proto.Value(values, %d)\n", index))
		case 'T':
			_, _ = b.WriteString(fmt.Sprintf(" proto.ParseDate(proto.Value(values, %d))\n", i))
		case 'D':
			_, _ = b.WriteString(fmt.Sprintf(" proto.ParseDuration(values[%d])\n", i))
		case 'C':
			_, _ = b.WriteString(fmt.Sprintf(" charAt(proto.Value(values, %d), 0)\n", i))
		}
	}
	b.WriteString("}\n")

	fmt.Println(b.String())
}

func toGoName(name string) string {
	if len(name) == 0 {
		return ""
	}
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ReplaceAll(name, "-", "")
	if proto.IsNumeral(name[0]) {
		return proto.NumeralToWord(name[0]) + name[1:]
	}
	return name
}
