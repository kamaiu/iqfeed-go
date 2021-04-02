package main

import (
	"fmt"
	"github.com/kamaiu/iqfeed-go/proto"
	"strings"
)

type FieldType string

const (
	FieldTypeUnknown  FieldType = ""
	FieldTypeString   FieldType = "S"
	FieldTypeFloat    FieldType = "F"
	FieldTypeInt      FieldType = "I"
	FieldTypeDateTime FieldType = "D"
	FieldTypeTimeSpan FieldType = "T"
)

func (f FieldType) GoTypeName() string {
	switch f {
	case FieldTypeString:
		return "string"
	case FieldTypeFloat:
		return "float64"
	case FieldTypeInt:
		return "int64"
	case FieldTypeDateTime:
		return "time.Time"
	case FieldTypeTimeSpan:
		return "time.Duration"
	}
	return ""
}

type Field struct {
	GoName    string
	Name      string
	Type      FieldType
	Nullable  bool
	Precision int
	Format    string
}

func main() {
	generateFundamentalMessage()
	//generateRegionalUpdateMessage()
	//generateUpdateSummaryMessage()
	//generateNewsMessage()
}

func generateFundamentalMessage() {
	fields := parse(`
string  	Symbol
string		Exchange ID
double		PE
int?		Average Volume
double?		52 Week High
double?  	52 Week Low
double?  	Calendar Year High
double?  	Calendar Year Low
double?  	Dividend Yield
double?  	Dividend Amount
double?  	Dividend Rate
DateTime?  	Pay Date
DateTime?  	Ex-dividend Date
(Reserved)
(Reserved)
(Reserved)
int			Short Interest
(Reserved)
double?		Current Year EPS
double?		Next Year EPS
double?	  	Five-year Growth Percentage
int		 	Fiscal Year End
(Reserved)
string  	Company Name
string	  	Root Option Symbol
double?	  	Percent Held By Institutions
double?	  	Beta
string  	Leaps
double?  	Current Assets
double?  	Current Liabilities
DateTime?	Balance Sheet Date
double?  	Long-term Debt
double?  	Common Shares Outstanding
(Reserved)
string		Split Factor 1
string	  	Split Factor 2
(Reserved)
string	  	Market Center
string  	Format Code
int?  		Precision
int?  		SIC
double?  	Historical Volatility
string  	Security Type
string  	Listed Market
DateTime  	52 Week High Date
DateTime?  	52 Week Low Date
DateTime?  	Calendar Year High Date
DateTime?  	Calendar Year Low Date
double?  	Year End Close
DateTime?  	Maturity Date
double?  	Coupon Rate
DateTime?  	Expiration Date
double?  	Strike Price
int?  		NAICS
string  	Exchange Root
double? 	Options Premium Multiplier
int?	 	Options Multiple Deliverables`,
		"MM/dd/yyyy", "MM/dd/yyyy")

	name := "Message"
	fmt.Println(generateStruct(name, fields))
	fmt.Println()
	fmt.Println(generateUnmarshal(name, 1, fields))
	fmt.Println()
	fmt.Println(generateToCSV(name, "F", fields))
	fmt.Println()
	fmt.Println(generateFields(name, fields))
}

func generateFields(structName string, fields []*Field) string {
	b := &strings.Builder{}
	b.WriteString("var (\n")
	for _, field := range fields {
		if len(field.Name) == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("    %s = &Field{\n", field.GoName))
		b.WriteString(fmt.Sprintf("        Name: \"%s\",\n", field.Name))
		b.WriteString(fmt.Sprintf("        Parser: func(t *%s, value string) {\n", structName))
		switch field.Type {
		case FieldTypeString:
			b.WriteString(fmt.Sprintf("            t.%s = value\n", field.GoName))
		case FieldTypeInt:
			b.WriteString(fmt.Sprintf("            t.%s = proto.ParseInt(value)\n", field.GoName))
		case FieldTypeFloat:
			b.WriteString(fmt.Sprintf("            t.%s = proto.ParseFloat(value)\n", field.GoName))
		case FieldTypeDateTime, FieldTypeTimeSpan:
			b.WriteString(fmt.Sprintf("            t.%s = %s(value)\n", field.GoName, dateParserFunc(field.Format)))
		}
		b.WriteString("        },\n")
		b.WriteString("    }\n")
	}
	b.WriteString("    byName = map[string]*Field{\n")
	for _, field := range fields {
		if len(field.Name) == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("    \"%s\": %s,\n", field.Name, field.GoName))
	}
	b.WriteString("    }\n")
	b.WriteString(")\n\n")

	b.WriteString("func Find(name string) *Field { return byName[name] }\n")
	return b.String()
}

func generateRegionalUpdateMessage() {
	fields := parse(`
		string 		Symbol
		string 		Exchange
        double 		Regional Bid
        int 		Regional Bid Size
        TimeSpan 	Regional Bid Time
        double 		Regional Ask
        int 		Regional Ask Size
        TimeSpan 	Regional Ask Time
        int 		Fraction Display Code
        int 		Decimal Precision
        int 		Market Center`, "MM/dd/yyyy", "HH:mm:ss")

	fmt.Println(generateStruct("RegionalUpdateMessage", fields))
	fmt.Println()
	fmt.Println(generateUnmarshal("RegionalUpdateMessage", 1, fields))
	fmt.Println()
}

func generateUpdateSummaryMessage() {
	fields := parse(`
		string 		Symbol
		double 		MostRecentTrade
        int 		MostRecentTradeSize
        TimeSpan 	MostRecentTradeTime
        int 		MostRecentTradeMarketCenter
        int 		TotalVolume
        double 		Bid
        int 		BidSize
        double 		Ask
        int 		AskSize
        double 		Open
        double 		High
        double 		Low
        double 		Close
        string 		MessageContents
        string 		MostRecentTradeConditions`, "MM/dd/yyyy", "HH:mm:ss.ffffff")

	fmt.Println(generateStruct("UpdateSummaryMessage", fields))
	fmt.Println()
	fmt.Println(generateUnmarshal("UpdateSummaryMessage", 1, fields))
	fmt.Println()
}

func generateNewsMessage() {
	fields := parse(`
		string 		DistributorCode
        string 		StoryId
        string 		SymbolList
        DateTime 	Timestamp
        string 		Headline`, "yyyyMMdd HHmmss", "HH:mm:ss.ffffff")

	fmt.Println(generateStruct("NewsMessage", fields))
	fmt.Println()
	fmt.Println(generateUnmarshal("NewsMessage", 1, fields))
	fmt.Println()
}

func dateFormatFunc(format string) string {
	switch format {
	case "MM/dd/yyyy":
		return "proto.AppendDate"
	}
	return "NO_FUNC"
}

func dateParserFunc(format string) string {
	switch format {
	case "yyyyMMdd HHmmss":
		return "proto.ParseDateTime"
	case "MM/dd/yyyy":
		return "proto.ParseDate"
	case "HH:mm:ss.ffffff":
		return "proto.ParseDuration"
	case "HH:mm:ss":
		return "proto.ParseDurationSecs"
	}
	return "NO_FUNC"
}

func generateStruct(name string, fields []*Field) string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("type %s struct {\n", name))

	for _, field := range fields {
		if len(field.Name) == 0 {
			continue
		}
		b.WriteString(fmt.Sprintf("    %s %s\n", field.GoName, field.Type.GoTypeName()))
	}

	b.WriteString("}\n")
	return b.String()
}

func generateUnmarshal(structName string, offset int, fields []*Field) string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("func (t *%s) Unmarshal(values []string) {\n", structName))
	//b.WriteString(fmt.Sprintf("    if len(values) < %d { return }\n", len(fields)+offset))

	for i, field := range fields {
		index := i + offset
		switch field.Type {
		case FieldTypeString:
			b.WriteString(fmt.Sprintf("    t.%s = proto.Value(values, %d)\n", field.GoName, index))
		case FieldTypeInt:
			b.WriteString(fmt.Sprintf("    t.%s = proto.ParseInt(proto.Value(values, %d))\n", field.GoName, index))
		case FieldTypeFloat:
			b.WriteString(fmt.Sprintf("    t.%s = proto.ParseFloat(proto.Value(values, %d))\n", field.GoName, index))
		case FieldTypeDateTime, FieldTypeTimeSpan:
			b.WriteString(fmt.Sprintf("    t.%s = %s(proto.Value(values, %d))\n", field.GoName, dateParserFunc(field.Format), index))
		}
	}
	b.WriteString("}\n")
	return b.String()
}

func generateToCSV(structName, prefix string, fields []*Field) string {
	b := &strings.Builder{}
	b.WriteString(fmt.Sprintf("func (t *%s) ToCSV(o []byte) []byte {\n", structName))
	if len(prefix) == 1 {
		b.WriteString(fmt.Sprintf("    o = append(o, '%s')\n", prefix[0:1]))
	} else {
		b.WriteString(fmt.Sprintf("    o = append(o, \"%s\"...)\n", prefix))
	}
	for _, field := range fields {
		b.WriteString("    o = append(o, ',')\n")

		switch field.Type {
		case FieldTypeString:
			b.WriteString(fmt.Sprintf("    o = append(o, t.%s...)\n", field.GoName))
		case FieldTypeInt:
			b.WriteString(fmt.Sprintf("    o = append(o, strconv.FormatInt(t.%s, 10)...)\n", field.GoName))
		case FieldTypeFloat:
			b.WriteString(fmt.Sprintf("    o = append(o, strconv.FormatFloat(t.%s, 'f', %d, 64)...)\n", field.GoName, field.Precision))
		case FieldTypeDateTime, FieldTypeTimeSpan:
			b.WriteString(fmt.Sprintf("    o = %s(o, t.%s)\n", dateFormatFunc(field.Format), field.GoName))
		default:
		}
	}
	b.WriteString("    return o\n")
	b.WriteString("}\n")
	return b.String()
}

func indexOfWhitespace(s string) int {
	for i, c := range s {
		switch c {
		case '\r', '\t', ' ':
			return i
		}
	}
	return -2
}

func parse(s string, dateFormat, durationFormat string) []*Field {
	fields := make([]*Field, 0, 64)
	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.Index(line, "(Reserved)") > -1 {
			fields = append(fields, &Field{
				Name:      "",
				Type:      "",
				Nullable:  false,
				Precision: 0,
				Format:    "",
			})
		}
		index := indexOfWhitespace(line)
		if index < 0 {
			continue
		}
		typeString := strings.TrimSpace(line[0:index])
		name := strings.TrimSpace(line[index+1:])
		t := FieldTypeUnknown

		nullable := false
		precision := 0
		format := ""
		switch typeString {
		case "double":
			t = FieldTypeFloat
			precision = 4
		case "double?":
			t = FieldTypeFloat
			nullable = true
			precision = 4
		case "DateTime":
			t = FieldTypeDateTime
			format = dateFormat
		case "DateTime?":
			t = FieldTypeDateTime
			format = dateFormat
			nullable = true
		case "TimeSpan":
			t = FieldTypeTimeSpan
			format = durationFormat
			nullable = true
		case "TimeSpan?":
			t = FieldTypeTimeSpan
			format = durationFormat
			nullable = true
		case "int":
			t = FieldTypeInt
		case "int?":
			t = FieldTypeInt
			nullable = true
		case "string":
			t = FieldTypeString
		case "string?":
			t = FieldTypeString
			nullable = true
		}
		field := &Field{
			GoName:    toGoName(name),
			Name:      name,
			Type:      t,
			Nullable:  nullable,
			Precision: precision,
			Format:    format,
		}
		fields = append(fields, field)
	}
	return fields
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

/*
string  	Symbol
string		Exchange ID
double		PE
int?		Average Volume
double?		52 Week High
double?  	52 Week Low
double?  	Calendar Year High
double?  	Calendar Year Low
double?  	Dividend Yield
double?  	Dividend Amount
double?  	Dividend Rate
DateTime?  	Pay Date
DateTime?  	Ex-dividend Date
(Reserved)
(Reserved)
(Reserved)
int			Short Interest
(Reserved)
double?		Current Year EPS
double?		Next Year EPS
double?	  	Five-year Growth Percentage
int		 	Fiscal Year End
(Reserved)
string  	Company Name
string	  	Root Option Symbol
double?	  	Percent Held By Institutions
double?	  	Beta
string  	Leaps
double?  	Current Assets
double?  	Current Liabilities
DateTime?	Balance Sheet Date
double?  	Long-term Debt
double?  	Common Shares Outstanding
(Reserved)
string		Split Factor 1
string	  	Split Factor 2
(Reserved)
string	  	Market Center
string  	Format Code
int?  		Precision
int?  		SIC
double?  	Historical Volatility
string  	Security Type
string  	Listed Market
DateTime  	52 Week High Date
DateTime?  	52 Week Low Date
DateTime?  	Calendar Year High Date
DateTime?  	Calendar Year Low Date
double?  	Year End Close
DateTime?  	Maturity Date
double?  	Coupon Rate
DateTime?  	Expiration Date
double?  	Strike Price
int?  		NAICS
string  	Exchange Root
double? 	Options Premium Multiplier
int?	 	Options Multiple Deliverables
*/

/*
`
			string Symbol
         string 	Symbol 					Symbol
         string 	ExchangeId 				Exchange ID
         double? 	PE 						PE
         int? 		AverageVolume 			Average Volume
         double? 	FiftyTwoWeekHigh		52 Week High
         double? 	FiftyTwoWeekLow			52 Week Low
         double? 	CalendarYearHigh { get; private set; }                   // 6
         double? 	CalendarYearLow // 7
         double? DividendYield   // 8
         double? DividendAmount  // 9
         double? DividendRate    // 10
         DateTime? PayDate        // 11
         DateTime? ExDividendDate // 12
        // (Reserved)                                               // 13
        // (Reserved)                                               // 14
        // (Reserved)                                               // 15
         int? ShortInterest       // 16
        // (Reserved)                                               // 17
         double? CurrentYearEarningsPerShare { get; private set; }        // 18
         double? NextYearEarningsPerShare { get; private set; }           // 19
         double? FiveYearGrowthPercentage { get; private set; }           // 20
         int? FiscalYearEnd       // 21
        // (Reserved)                                               // 22
         string CompanyName       // 23
         string RootOptionSymbol  // 24
         double? PercentHeldByInstitutions { get; private set; }          // 25
         double? Beta            // 26
         string Leaps             // 27
         double? CurrentAssets   // 28
         double? CurrentLiabilities { get; private set; }                 // 29
         DateTime? BalanceSheetDate { get; private set; }                  // 30
         double? LongTermDebt    // 31
         double? CommonSharesOutstanding { get; private set; }            // 32
        // (Reserved)                                               // 33
         string SplitFactor1      // 34
         string SplitFactor2      // 35
        // (Reserved)                                               // 36
        // (Reserved)                                               // 37
         string FormatCode        // 38
         int? Precision           // 39
         int? SIC                 // 40
         double? HistoricalVolatility { get; private set; }               // 41
         string SecurityType      // 42
         string ListedMarket      // 43
         DateTime? FiftyTwoWeekHighDate { get; private set; }              // 44
         DateTime? FiftyTwoWeekLowDate { get; private set; }               // 45
         DateTime? CalendarYearHighDate { get; private set; }              // 46
         DateTime? CalendarYearLowDate { get; private set; }               // 47
         double? YearEndClose    // 48
         DateTime? MaturityDate   // 49
         double? CouponRate      // 50
         DateTime? ExpirationDate // 51
         double? StrikePrice     // 52
         int? NAICS               // 53
         string ExchangeRoot      // 54
         double? OptionsPremiumMultiplier { get; private set; }           // 55
         int? OptionsMultipleDeliverables { get; private set; }            // 56`
*/
