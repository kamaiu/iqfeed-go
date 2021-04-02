package iqfeed

const (
	Hostname       = "localhost"
	Level1Port     = 5009
	LookupPort     = 9100
	Level2Port     = 9200
	AdminPort      = 9300
	DerivativePort = 9400
	LoginPort      = 60020
)

type Connect struct {
	username            string
	password            string
	productId           string
	productVersion      string
	connectionTimeoutMs int
	retry               int
}

//func DoConnect(
//	username string,
//	password string,
//	productId string,
//	productVersion string,
//	connectionTimeoutMs int,
//	retry int,
//) (*Connect, error) {
//
//}
