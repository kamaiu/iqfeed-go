package level1

import (
	"github.com/kamaiu/iqfeed-go/proto"
	"time"
)

type NewsMessage struct {
	DistributorCode string
	StoryId         string
	SymbolList      string
	Timestamp       time.Time
	Headline        string
}

func (t *NewsMessage) Unmarshal(values []string) {
	t.DistributorCode = proto.Value(values, 1)
	t.StoryId = proto.Value(values, 2)
	t.SymbolList = proto.Value(values, 3)
	t.Timestamp = proto.ParseDateTime(proto.Value(values, 4))
	t.Headline = proto.Value(values, 5)
}
