package symbol

import "github.com/kamaiu/iqfeed-go/proto"

type NaicsCodeInfoMessage struct {
	NaicsCode   int64
	Description string
	RequestId   string
}

func (l *NaicsCodeInfoMessage) Unmarshal(values []string) {
	l.NaicsCode = proto.ParseInt(values[0])
	l.Description = values[1]
}

func (l *NaicsCodeInfoMessage) UnmarshalWithRequestId(values []string) {
	l.RequestId = values[0]
	l.NaicsCode = proto.ParseInt(values[1])
	l.Description = values[2]
}
