package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type SecurityType string

const (
	SecurityTypeSTRIP            SecurityType = "STRIP"
	SecurityTypeFUTURE           SecurityType = "FUTURE"
	SecurityTypeCOMBINED_FUTURE  SecurityType = "COMBINED_FUTURE"
	SecurityTypeINDEX            SecurityType = "INDEX"
	SecurityTypeCOMBINED_FOPTION SecurityType = "COMBINED_FOPTION"
	SecurityTypeFOPTION          SecurityType = "FOPTION"
	SecurityTypeSTRATSPREAD      SecurityType = "STRATSPREAD"
	SecurityTypeSPREAD           SecurityType = "SPREAD"
	SecurityTypeICSPREAD         SecurityType = "ICSPREAD"
	SecurityTypeEQUITY           SecurityType = "EQUITY"
	SecurityTypeIEOPTION         SecurityType = "IEOPTION"
	SecurityTypeBONDS            SecurityType = "BONDS"
	SecurityTypeSPOT             SecurityType = "SPOT"
	SecurityTypeMUTUAL           SecurityType = "MUTUAL"
	SecurityTypeMONEY            SecurityType = "MONEY"
	SecurityTypeFOREX            SecurityType = "FOREX"
	SecurityTypeMKTSTATS         SecurityType = "MKTSTATS"
	SecurityTypePRECMTL          SecurityType = "PRECMTL"
	SecurityTypeMKTRPT           SecurityType = "MKTRPT"
	SecurityTypeTREASURIES       SecurityType = "TREASURIES"
	SecurityTypeFORWARD          SecurityType = "FORWARD"
	SecurityTypeSWAPS            SecurityType = "SWAPS"
)

type SecurityTypeMessage struct {
	SecurityTypeId int64
	ShortName      SecurityType
	LongName       string
	RequestId      string
}

func (l *SecurityTypeMessage) Unmarshal(values []string) {
	l.SecurityTypeId = proto.ParseInt(values[0])
	l.ShortName = SecurityType(values[1])
	l.LongName = values[2]
}

func (l *SecurityTypeMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.SecurityTypeId = proto.ParseInt(values[1])
	l.ShortName = SecurityType(values[2])
	l.LongName = values[3]
}
