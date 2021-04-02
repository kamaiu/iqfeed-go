package level1

import (
	"github.com/kamaiu/iqfeed-go"
	"testing"
	"time"
)

func TestOpen(t *testing.T) {
	c, err := Open("localhost", iqfeed.Level1Port)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Watch("@EU#")
	time.Sleep(time.Hour)
}
