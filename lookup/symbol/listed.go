package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type ListedExchange string

const (
	ListedExchangeNYMEX_GBX     ListedExchange = "NYMEX_GBX"
	ListedExchangeCME           ListedExchange = "CME"
	ListedExchangeNASDAQ        ListedExchange = "NASDAQ"
	ListedExchangeEEXP          ListedExchange = "EEXP"
	ListedExchangeDTN           ListedExchange = "DTN"
	ListedExchangeENID          ListedExchange = "ENID"
	ListedExchangeNYSE_ARCA     ListedExchange = "NYSE_ARCA"
	ListedExchangeEEXE          ListedExchange = "EEXE"
	ListedExchangeEEXN          ListedExchange = "EEXN"
	ListedExchangeFTSE          ListedExchange = "FTSE"
	ListedExchangeSGX           ListedExchange = "SGX"
	ListedExchangeCME_FL        ListedExchange = "CME-FL"
	ListedExchangeNYMEX         ListedExchange = "NYMEX"
	ListedExchangeCBOT          ListedExchange = "CBOT"
	ListedExchangeCBOT_GBX      ListedExchange = "CBOT_GBX"
	ListedExchangeCME_GBX       ListedExchange = "CME_GBX"
	ListedExchangeCMEMINI       ListedExchange = "CMEMINI"
	ListedExchangeMGE_GBX       ListedExchange = "MGE_GBX"
	ListedExchangeNFX           ListedExchange = "NFX"
	ListedExchangeICEFU         ListedExchange = "ICEFU"
	ListedExchangeICEFC         ListedExchange = "ICEFC"
	ListedExchangeKCBOT_GBX     ListedExchange = "KCBOT_GBX"
	ListedExchangeICEEF         ListedExchange = "ICEEF"
	ListedExchangeCBOTMINI      ListedExchange = "CBOTMINI"
	ListedExchangeICEEC         ListedExchange = "ICEEC"
	ListedExchangeKBCB          ListedExchange = "KBCB"
	ListedExchangeMGKB          ListedExchange = "MGKB"
	ListedExchangeMGCB          ListedExchange = "MGCB"
	ListedExchangeCOMEX_GBX     ListedExchange = "COMEX_GBX"
	ListedExchangeNYMEXMINI     ListedExchange = "NYMEXMINI"
	ListedExchangeCFE           ListedExchange = "CFE"
	ListedExchangeNYSE          ListedExchange = "NYSE"
	ListedExchangePK_NYSE       ListedExchange = "PK_NYSE"
	ListedExchangeOPRA          ListedExchange = "OPRA"
	ListedExchangeDJ            ListedExchange = "DJ"
	ListedExchangeOTC           ListedExchange = "OTC"
	ListedExchangePK_GREY       ListedExchange = "PK_GREY"
	ListedExchangePK_CURRENT    ListedExchange = "PK_CURRENT"
	ListedExchangeNGSM          ListedExchange = "NGSM"
	ListedExchangePK_NOINFO     ListedExchange = "PK_NOINFO"
	ListedExchangeNYSE_AMERICAN ListedExchange = "NYSE_AMERICAN"
	ListedExchangePK_NYSE_AMEX  ListedExchange = "PK_NYSE_AMEX"
	ListedExchangeNGM           ListedExchange = "NGM"
	ListedExchangePK_LIMITED    ListedExchange = "PK_LIMITED"
	ListedExchangePK_IQXPREM    ListedExchange = "PK_IQXPREM"
	ListedExchangePK_NASDAQ     ListedExchange = "PK_NASDAQ"
	ListedExchangeNCM           ListedExchange = "NCM"
	ListedExchangePK_OTCQB      ListedExchange = "PK_OTCQB"
	ListedExchangeEEXAG         ListedExchange = "EEXAG"
	ListedExchangePK_IQXPRIME   ListedExchange = "PK_IQXPRIME"
	ListedExchangeOTCBB         ListedExchange = "OTCBB"
	ListedExchangeASXCM         ListedExchange = "ASXCM"
	ListedExchangeBMF           ListedExchange = "BMF"
	ListedExchangeBATS          ListedExchange = "BATS"
	ListedExchangeTENFORE       ListedExchange = "TENFORE"
	ListedExchangePK_QXPREM     ListedExchange = "PK_QXPREM"
	ListedExchangePK_QXPRIME    ListedExchange = "PK_QXPRIME"
	ListedExchangeCMEUR         ListedExchange = "CMEUR"
	ListedExchangePK_ARCA       ListedExchange = "PK_ARCA"
	ListedExchangeSAFEX         ListedExchange = "SAFEX"
	ListedExchangeEUREX         ListedExchange = "EUREX"
	ListedExchangeFXCM          ListedExchange = "FXCM"
	ListedExchangeBLOOMBERG     ListedExchange = "BLOOMBERG"
	ListedExchangeICEENDEX      ListedExchange = "ICEENDEX"
	ListedExchangeCBOE          ListedExchange = "CBOE"
	ListedExchangeCLEARPORT     ListedExchange = "CLEARPORT"
	ListedExchangeCVE           ListedExchange = "CVE"
	ListedExchangeTSE           ListedExchange = "TSE"
	ListedExchangeENCOM         ListedExchange = "ENCOM"
	ListedExchangeCFTC          ListedExchange = "CFTC"
	ListedExchangeKCBOT         ListedExchange = "KCBOT"
	ListedExchangeICEEA         ListedExchange = "ICEEA"
	ListedExchangeDCE           ListedExchange = "DCE"
	ListedExchangeELSPOT        ListedExchange = "ELSPOT"
	ListedExchangeN2EX          ListedExchange = "N2EX"
	ListedExchangeCANTOR        ListedExchange = "CANTOR"
	ListedExchangeCOMEX         ListedExchange = "COMEX"
	ListedExchangeMDEX          ListedExchange = "MDEX"
	ListedExchangeMCX           ListedExchange = "MCX"
	ListedExchangeMGE           ListedExchange = "MGE"
	ListedExchangeLSE           ListedExchange = "LSE"
	ListedExchangeLSEI          ListedExchange = "LSEI"
	ListedExchangePK_YL_SHEETS  ListedExchange = "PK_YL_SHEETS"
	ListedExchangeLME           ListedExchange = "LME"
	ListedExchangeLPPM          ListedExchange = "LPPM"
	ListedExchangeSGXAC         ListedExchange = "SGXAC"
	ListedExchangeRUSSELL_FL    ListedExchange = "RUSSELL-FL"
	ListedExchangeRUSSELL       ListedExchange = "RUSSELL"
	ListedExchangeEEXC          ListedExchange = "EEXC"
	ListedExchangeTOCOM         ListedExchange = "TOCOM"
	ListedExchangeUSDA          ListedExchange = "USDA"
	ListedExchangeGRNST         ListedExchange = "GRNST"
	ListedExchangeWASDE         ListedExchange = "WASDE"
	ListedExchangeZCE           ListedExchange = "ZCE"
	ListedExchangeIEX           ListedExchange = "IEX"
)

type ListedMarketMessage struct {
	ListedMarketId int64
	ShortName      string
	LongName       string
	GroupId        int64
	ShortGroupName string
	RequestId      string
}

func (l *ListedMarketMessage) Unmarshal(values []string) {
	l.ListedMarketId = proto.ParseInt(values[0])
	l.ShortName = values[1]
	l.LongName = values[2]
	l.GroupId = proto.ParseInt(values[3])
	l.ShortGroupName = values[4]
}

func (l *ListedMarketMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.ListedMarketId = proto.ParseInt(values[1])
	l.ShortName = values[2]
	l.LongName = values[3]
	l.GroupId = proto.ParseInt(values[4])
	l.ShortGroupName = values[5]
}
