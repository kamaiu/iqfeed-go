package admin

import (
	"errors"
	"github.com/kamaiu/iqfeed-go/proto"
	"strconv"
	"time"
)

type ProtocolMessage struct {
	Version string
}

type AutoConnectMessage struct {
	On bool
}

type ClientAppMessageType string

const (
	ClientAppMessageTypeRegister ClientAppMessageType = "Register"
	ClientAppMessageTypeRemove   ClientAppMessageType = "Remove"
)

type ClientAppMessage struct {
	Type ClientAppMessageType
}

type ClientType int

const (
	ClientTypeAdmin  ClientType = 0
	ClientTypeLevel1 ClientType = 1
	ClientTypeLevel2 ClientType = 2
	ClientTypeLookup ClientType = 3
)

const (
	ClientStatsDatetimeFormat = "yyyyMMdd HHmmss"
)

type ClientStatsMessage struct {
	Type            ClientType
	ClientId        int64
	ClientName      string
	StartTime       time.Time
	Symbols         int64
	RegionalSymbols int64
	KbReceived      float64
	KbSent          float64
	KbQueued        float64
}

func (c *ClientStatsMessage) Unmarshal(buf []string) error {
	if len(buf) < 10 {
		return errors.New("expected 10 fields")
	}
	var err error
	var t int64
	t, err = strconv.ParseInt(buf[1], 10, 64)
	if err != nil {
		return err
	}
	switch t {
	case 0:
		c.Type = ClientTypeAdmin
	case 1:
		c.Type = ClientTypeLevel1
	case 2:
		c.Type = ClientTypeLevel2
	case 3:
		c.Type = ClientTypeLookup
	}

	c.ClientId, err = strconv.ParseInt(buf[2], 10, 64)
	c.ClientName = string(buf[3])
	c.StartTime = proto.ParseDateTime(buf[4]) // yyyyMMdd HHmmss

	c.Symbols = proto.ParseInt(buf[5])
	c.RegionalSymbols = proto.ParseInt(buf[6])
	c.KbReceived = proto.ParseFloat(buf[6])
	c.KbSent = proto.ParseFloat(buf[7])
	c.KbQueued = proto.ParseFloat(buf[8])
	return nil
}

type LoginIdMessage struct {
	UserLoginId string
}

type LoginInfoMessageType string

const (
	LoginInfoMessageTypeSaved    LoginInfoMessageType = "Saved"
	LoginInfoMessageTypeNotSaved LoginInfoMessageType = "NotSaved"
)

type LoginInfoMessage struct {
	Type LoginInfoMessageType
}

type PasswordMessage struct {
	UserPassword string
}

type StatsStatusType string

const (
	StatsStatusTypeNotConnected StatsStatusType = "NotConnected"
	StatsStatusTypeConnected    StatsStatusType = "Connected"
)

type StatsMessage struct {
	ServerIp               string
	ServerPort             int64
	MaxSymbols             int64
	NumberOfSymbols        int64
	ClientsConnected       int64
	SecondsSinceLastUpdate int64
	Reconnections          int64
	AttemptedReconnections int64
	StartTime              time.Time
	MarketTime             time.Time
	Status                 StatsStatusType
	IQFeedVersion          string
	LoginId                string
	TotalKBsRecv           float64
	KBsPerSecRecv          float64
	AvgKBsPerSecRecv       float64
	TotalKBsSent           float64
	KBsPerSecSent          float64
	AvgKBsPerSecSent       float64
}

func (s *StatsMessage) Unmarshal(buf []string) error {
	if len(buf) < 20 {
		return errors.New("expected 20 fields")
	}

	s.ServerIp = buf[1]
	s.ServerPort = proto.ParseInt(buf[2])
	s.MaxSymbols = proto.ParseInt(buf[3])
	s.NumberOfSymbols = proto.ParseInt(buf[4])
	s.ClientsConnected = proto.ParseInt(buf[5])
	s.SecondsSinceLastUpdate = proto.ParseInt(buf[6])
	s.Reconnections = proto.ParseInt(buf[7])
	s.AttemptedReconnections = proto.ParseInt(buf[8])

	s.StartTime = proto.ParseMonthDayTime(buf[9])   // time.Parse("MMM dd H:mmtt", buf[9])
	s.MarketTime = proto.ParseMonthDayTime(buf[10]) // time.Parse("MMM dd H:mmtt", buf[10])

	switch buf[11] {
	case "Connected":
		s.Status = StatsStatusTypeConnected
	default:
		s.Status = StatsStatusTypeNotConnected
	}
	s.IQFeedVersion = buf[12]
	s.LoginId = buf[13]
	s.TotalKBsRecv = proto.ParseFloat(buf[14])
	s.KBsPerSecRecv = proto.ParseFloat(buf[15])
	s.AvgKBsPerSecRecv = proto.ParseFloat(buf[16])
	s.TotalKBsSent = proto.ParseFloat(buf[17])
	s.KBsPerSecSent = proto.ParseFloat(buf[18])
	s.AvgKBsPerSecSent = proto.ParseFloat(buf[19])
	return nil
}
