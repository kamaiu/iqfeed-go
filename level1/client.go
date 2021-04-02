package level1

import (
	"errors"
	"fmt"
	"github.com/kamaiu/iqfeed-go/level1/fundamental"
	"github.com/kamaiu/iqfeed-go/level1/update"
	"github.com/kamaiu/iqfeed-go/proto"
	"net"
	"os"
	"strings"
	"sync"
	"unsafe"
)

type Client struct {
	buf  []string
	conn net.Conn

	err error

	dynamic           bool
	levels            []*proto.LogLevel
	fields            []*update.Field
	fundamentalFields []*fundamental.Field
	updateFields      []*update.Field
	currentFields     []*update.Field
	closed            bool
	disconnect        bool

	wg sync.WaitGroup
	mu sync.RWMutex
}

func Open(ip string, port int) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return nil, err
	}
	c := &Client{
		buf:  make([]string, 0, 24),
		conn: conn,
	}
	c.wg.Add(1)
	go c.run()

	_ = c.ServerConnect()
	_ = c.Timestamps(false)
	_ = c.FundamentalFieldnames()
	_ = c.UpdateFieldnames()
	_ = c.CurrentUpdateFieldNames()

	return c, nil
}

type Symbol struct {
	Name string
}

func (c *Client) Watch(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqWatch(nil, symbol))
	return err
}

func (c *Client) TradesOnlyWatch(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqTradesOnlyWatch(nil, symbol))
	return err
}

func (c *Client) Unwatch(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqUnwatch(nil, symbol))
	return err
}

func (c *Client) ForcedRefresh(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqForcedRefresh(nil, symbol))
	return err
}

func (c *Client) Timestamp() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqTimestamp(nil))
	return err
}

func (c *Client) Timestamps(on bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqTimestamps(nil, on))
	return err
}

func (c *Client) RegionalWatch(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqRegionalWatch(nil, symbol))
	return err
}

func (c *Client) RegionalUnwatch(symbol string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqRegionalUnwatch(nil, symbol))
	return err
}

func (c *Client) News(on bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqNews(nil, on))
	return err
}

func (c *Client) Stats() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqStats(nil))
	return err
}

func (c *Client) FundamentalFieldnames() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqFundamentalFieldnames(nil))
	return err
}

func (c *Client) UpdateFieldnames() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqUpdateFieldnames(nil))
	return err
}

func (c *Client) CurrentUpdateFieldNames() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqCurrentUpdateFieldNames(nil))
	return err
}

func (c *Client) SelectUpdateFieldName(fields []*update.Field) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.fields = fields
	_, err := c.conn.Write(selectUpdateFieldName(nil, fields))
	return err
}

func (c *Client) SetLogLevels(levels []*proto.LogLevel) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.levels = levels
	_, err := c.conn.Write(setLogLevels(nil, levels))
	return err
}

func (c *Client) WatchList() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqWatchList(nil))
	return err
}

func (c *Client) UnwatchAll() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(reqUnwatchAll(nil))
	return err
}

func (c *Client) ServerConnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.disconnect = false
	_, err := c.conn.Write(reqServerConnect(nil))
	return err
}

func (c *Client) ServerDisconnect() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.disconnect {
		return os.ErrClosed
	}
	c.disconnect = true
	_, err := c.conn.Write(reqServerDisconnect(nil))
	return err
}

func (cl *Client) run() {
	defer func() {
		cl.mu.Lock()
		cl.closed = true
		cl.mu.Unlock()
		cl.wg.Done()
	}()

	const (
		bufferSize    = 8192
		maxBufferSize = bufferSize * 512
	)
	var (
		err   error
		b     = make([]byte, bufferSize)
		n     = 0
		mark  = 0
		count = 0
	)

LOOP:
	for {
		n, err = cl.conn.Read(b[mark:])
		if err != nil {
			break LOOP
		}
		if n < 1 {
			continue
		}

		i := mark
		n = n + mark
		mark = 0
		count = 0
		for ; i < n; i++ {
			c := b[i]
			switch c {
			case '\n':
				str := btos(b[mark:i])
				str = strings.TrimSpace(str)
				_ = cl.process(str)
				count++
				mark = i + 1
			}
		}

		if mark < n {
			if count <= 1 {
				if cap(b)+bufferSize > maxBufferSize {
					err = errors.New("max buffer size reached")
					break LOOP
				}
				next := make([]byte, cap(b)+bufferSize)
				copy(next, b[mark:])
				b = next
			} else {
				if mark > 0 {
					copy(b[0:], b[mark:])
				}
			}
		} else {
			mark = 0
		}
	}

	if err != nil {
		cl.mu.Lock()
		cl.err = err
		cl.mu.Unlock()
	}
}

func (cl *Client) process(msg string) error {
	if len(msg) == 0 {
		return nil
	}
	buf := cl.buf[:0]
	mark := 0
	for i := 0; i < len(msg); i++ {
		c := msg[i]
		switch c {
		case proto.ProtocolDelimiterCharacter:
			buf = append(buf, msg[mark:i])
			mark = i + 1
		}
	}

	if mark < len(msg) {
		buf = append(buf, msg[mark:])
	}
	cl.buf = buf

	if len(buf) > 0 {
		switch msg[0] {
		case 'F': // A fundamental message
			cl.processFundamentalMessage(buf)
		case 'P': // A summary message
			cl.processSummaryMessage(buf)
		case 'Q': // An update message
			cl.processUpdateMessage(buf)
		case 'R': // A regional update message
			cl.processRegionalUpdateMessage(buf)
		case 'N': // A news headline message
			cl.processNewsMessage(buf)
		case 'S': // A system message
			cl.processSystemMessage(buf)
		case 'T': // A timestamp message
			cl.processTimestampMessage(buf)
		case 'n': // Symbol not found message
			cl.processSymbolNotFoundMessage(buf)
		case 'E': // An error message
			cl.processErrorMessage(buf)
		}
	}
	return nil
}

func (c *Client) processFundamentalMessage(values []string) {
	c.mu.RLock()
	fields := c.fundamentalFields
	c.mu.RUnlock()

	m := fundamental.Message{}
	m.Unmarshal(values[1:], fields)
	fmt.Println(m)
}

func (c *Client) processSummaryMessage(values []string) {
	c.mu.RLock()
	fields := c.currentFields
	c.mu.RUnlock()

	if len(fields) > 0 {
		m := &update.Message{}
		m.Unmarshal(values[1:], fields)
		fmt.Println("Summary -> bid:", m.Bid, " ask:", m.Ask)
	} else {
		m := &UpdateSummaryMessage{}
		m.Unmarshal(values)
		fmt.Println(m)
	}
}

func (c *Client) processUpdateMessage(values []string) {
	c.mu.RLock()
	fields := c.updateFields
	c.mu.RUnlock()
	if len(fields) > 0 {
		m := &update.Message{}
		m.Unmarshal(values[1:], fields)
		fmt.Println("Update -> bid:", m.Bid, " ask:", m.Ask)
	} else {
		m := &UpdateSummaryMessage{}
		m.Unmarshal(values)
	}
}

func (c *Client) processRegionalUpdateMessage(values []string) {
	m := &RegionalUpdateMessage{}
	m.Unmarshal(values)
}

func (c *Client) processNewsMessage(values []string) {
	m := &NewsMessage{}
	m.Unmarshal(values)
}

func (c *Client) processSystemMessage(values []string) {
	switch proto.Value(values, 1) {
	case "FUNDAMENTAL FIELDNAMES":
		fields := make([]*fundamental.Field, 0, len(values)-2)
		for i := 2; i < len(values); i++ {
			f := fundamental.Find(values[i])
			if f == nil {
				if values[i] != "(Reserved)" {
					fmt.Println(values[i])
				}
				f = &fundamental.Field{
					Name: values[i],
					Parser: func(u *fundamental.Message, val string) {
						// Ignore
					},
				}
			}
			fields = append(fields, f)
		}

		c.mu.Lock()
		c.fundamentalFields = fields
		c.mu.Unlock()

	case "UPDATE FIELDNAMES", "CURRENT UPDATE FIELDNAMES":
		fields := make([]*update.Field, 0, len(values)-2)
		for i := 2; i < len(values); i++ {
			f := update.Find(values[i])
			if f == nil {
				if values[i] != "(Reserved)" {
					fmt.Println(values[i])
				}
				f = &update.Field{
					Name: values[i],
					Parser: func(u *update.Message, val string) {
						// Ignore
					},
				}
			}
			fields = append(fields, f)
		}

		switch proto.Value(values, 1) {
		case "UPDATE FIELDNAMES":
			c.mu.Lock()
			c.updateFields = fields
			c.mu.Unlock()
		case "CURRENT UPDATE FIELDNAMES":
			c.mu.Lock()
			c.currentFields = fields
			c.mu.Unlock()
		}

	default:
		m := &proto.SystemMessage{}
		m.Unmarshal(values)
	}

}

func (c *Client) processTimestampMessage(values []string) {
	m := &proto.TimestampMessage{}
	m.Unmarshal(values)
}

func (c *Client) processSymbolNotFoundMessage(values []string) {
	m := &proto.SymbolNotFoundMessage{}
	m.Unmarshal(values)
}

func (c *Client) processErrorMessage(values []string) {
	m := &proto.ErrorMessage{}
	m.Unmarshal(values)
}

func btos(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func atIndex(values []string, i uint) string {
	if len(values) == 0 {
		return ""
	}
	if i < uint(len(values)) {
		return values[i]
	}
	return ""
}
