package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type SicCodeInfoMessage struct {
	SicCode     int64
	Description string
	RequestId   string
}

func (l *SicCodeInfoMessage) Unmarshal(values []string) {
	l.SicCode = proto.ParseInt(values[0])
	l.Description = values[1]
}

func (l *SicCodeInfoMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.SicCode = proto.ParseInt(values[1])
	l.Description = values[2]
}
