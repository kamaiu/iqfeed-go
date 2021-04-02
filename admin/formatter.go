package admin

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"strings"
)

func RegisterClientApp(b []byte, productId, productVersion string) []byte {
	b = append(b, "S,REGISTER CLIENT APP,"...)
	b = append(b, strings.ToUpper(productId)...)
	b = append(b, ',')
	b = append(b, strings.ToUpper(productVersion)...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func RemoveClientApp(b []byte, productId, productVersion string) []byte {
	b = append(b, "S,REMOVE CLIENT APP,"...)
	b = append(b, strings.ToUpper(productId)...)
	b = append(b, ',')
	b = append(b, strings.ToUpper(productVersion)...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func SetLoginId(b []byte, userLoginId string) []byte {
	b = append(b, "S,SET LOGINID,"...)
	b = append(b, strings.ToUpper(userLoginId)...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func SetPassword(b []byte, password string) []byte {
	b = append(b, "S,SET PASSWORD,"...)
	b = append(b, strings.ToUpper(password)...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func SetSaveLoginInfo(b []byte, on bool) []byte {
	b = append(b, "S,SET SAVE LOGIN INFO,"...)
	if on {
		b = append(b, "On"...)
	} else {
		b = append(b, "Off"...)
	}
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func SetAutoConnect(b []byte, on bool) []byte {
	b = append(b, "S,SET AUTOCONNECT,"...)
	if on {
		b = append(b, "On"...)
	} else {
		b = append(b, "Off"...)
	}
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func ReqServerConnect(b []byte) []byte {
	b = append(b, "S,CONNECT"...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func ReqServerDisconnect(b []byte) []byte {
	b = append(b, "S,DISCONNECT"...)
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}

func SetClientStats(b []byte, on bool) []byte {
	b = append(b, "S,CLIENTSTATS "...)
	if on {
		b = append(b, "ON"...)
	} else {
		b = append(b, "OFF"...)
	}
	b = append(b, proto.ProtocolTerminatingCharacters...)
	return b
}
