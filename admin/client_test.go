package admin

import (
	"github.com/kamaiu/iqfeed-go"
	"github.com/kamaiu/iqfeed-go/proto"
	"testing"
	"time"
)

func TestStatsMessage_Unmarshal(t *testing.T) {
	msg := StatsMessage{}
	err := msg.Unmarshal(proto.Split("S,STATS,66.112.156.227,60003,500,1,8,0,0,0,Apr 01 6:48PM,Apr 01 08:49PM,Connected,6.2.0.15,497214,3.24,0.04,0.05,50.27,1.34,0.72,", nil))
	if err != nil {
		t.Fatal(err)
	}

	/*
		"S,AUTOCONNECT ON,"
		"S,CURRENT PROTOCOL,6.0,"
		"S,CURRENT LOGINID,3244,"
		"S,CURRENT PASSWORD,(3*****2),"
		"S,STATS,66.112.156.227,60003,500,1,8,0,0,0,Apr 01 6:48PM,Apr 01 08:54PM,Connected,6.2.0.15,497214,10.27,0.04,0.03,366.45,1.14,0.96,"
	*/
}

func TestOpen(t *testing.T) {
	c, err := Open("localhost", iqfeed.AdminPort)
	if err != nil {
		t.Fatal(err)
	}
	_ = c
	c.SetAutoConnect(true)
	c.SetProtocol(proto.ProtocolVersion)
	c.SetLoginId("")
	c.SetPassword("")
	c.SetClientStats(true)
	time.Sleep(time.Hour)
}
