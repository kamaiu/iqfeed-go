package level1

import (
	"github.com/kamaiu/iqfeed-go/level1/update"
	"github.com/kamaiu/iqfeed-go/proto"
	"strings"
)

func reqWatch(o []byte, symbol string) []byte {
	o = append(o, 'w')
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqTradesOnlyWatch(o []byte, symbol string) []byte {
	o = append(o, 't')
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqUnwatch(o []byte, symbol string) []byte {
	o = append(o, 'r')
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqForcedRefresh(o []byte, symbol string) []byte {
	o = append(o, 'f')
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqTimestamp(o []byte) []byte {
	o = append(o, 'T')
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqTimestamps(o []byte, on bool) []byte {
	o = append(o, "S,"...)
	if on {
		o = append(o, "TIMESTAMPSON"...)
	} else {
		o = append(o, "TIMESTAMPSOFF"...)
	}
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqRegionalWatch(o []byte, symbol string) []byte {
	o = append(o, "S,REGON,"...)
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqRegionalUnwatch(o []byte, symbol string) []byte {
	o = append(o, "S,REGOFF,"...)
	o = append(o, strings.ToUpper(symbol)...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqNews(o []byte, on bool) []byte {
	o = append(o, "S,"...)
	if on {
		o = append(o, "NEWSON"...)
	} else {
		o = append(o, "NEWSOFF"...)
	}
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqStats(o []byte) []byte {
	o = append(o, "S,REQUEST STATS"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqFundamentalFieldnames(o []byte) []byte {
	o = append(o, "S,REQUEST FUNDAMENTAL FIELDNAMES"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqUpdateFieldnames(o []byte) []byte {
	o = append(o, "S,REQUEST ALL UPDATE FIELDNAMES"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqCurrentUpdateFieldNames(o []byte) []byte {
	o = append(o, "S,REQUEST CURRENT UPDATE FIELDNAMES"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func selectUpdateFieldName(o []byte, fields []*update.Field) []byte {
	o = append(o, "S,SELECT UPDATE FIELDS"...)
	for _, f := range fields {
		o = append(o, proto.ProtocolDelimiterCharacter)
		o = append(o, f.Name...)
	}
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func setLogLevels(o []byte, levels []*proto.LogLevel) []byte {
	o = append(o, "S,SET LOG LEVELS"...)
	for _, level := range levels {
		o = append(o, proto.ProtocolDelimiterCharacter)
		o = append(o, level.String()...)
	}
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqWatchList(o []byte) []byte {
	o = append(o, "S,REQUEST WATCHES"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqUnwatchAll(o []byte) []byte {
	o = append(o, "S,UNWATCH ALL"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqServerConnect(o []byte) []byte {
	o = append(o, "S,CONNECT"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}

func reqServerDisconnect(o []byte) []byte {
	o = append(o, "S,DISCONNECT"...)
	o = append(o, proto.ProtocolTerminatingCharacters...)
	return o
}
