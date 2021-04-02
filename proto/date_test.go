package proto

import (
	"fmt"
	"testing"
)

func TestParseDateTime(t *testing.T) {
	d := ParseDateTime("20210115 010876")
	fmt.Println(d)

	dur := ParseDuration("02:04:23.343524")
	fmt.Println(dur)
}
