package admin

import (
	"errors"
	"fmt"
	"github.com/kamaiu/iqfeed-go/proto"
	"net"
	"strings"
	"sync"
	"unsafe"
)

type Client struct {
	buf  []string
	conn net.Conn

	saveLoginInfo bool
	autoConnect   bool
	clientStats   bool

	err error

	closed bool
	wg     sync.WaitGroup
	mu     sync.RWMutex
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
	return c, nil
}

func (c *Client) SetProtocol(version string) error {
	_, err := c.conn.Write(proto.SetProtocol(nil, version))
	return err
}

func (c *Client) SetClientName(name string) error {
	_, err := c.conn.Write(proto.SetClientName(nil, name))
	return err
}

func (c *Client) RegisterClientApp(productId, productVersion string) error {
	_, err := c.conn.Write(RegisterClientApp(nil, productId, productVersion))
	return err
}

func (c *Client) RemoveClientApp(productId, productVersion string) error {
	_, err := c.conn.Write(RemoveClientApp(nil, productId, productVersion))
	return err
}

func (c *Client) SetLoginId(userLoginId string) error {
	_, err := c.conn.Write(SetLoginId(nil, userLoginId))
	return err
}

func (c *Client) SetPassword(userLoginId string) error {
	_, err := c.conn.Write(SetPassword(nil, userLoginId))
	return err
}

func (c *Client) SetSaveLoginInfo(on bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(SetSaveLoginInfo(nil, on))
	c.saveLoginInfo = on
	return err
}

func (c *Client) SetAutoConnect(on bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(SetAutoConnect(nil, on))
	c.autoConnect = on
	return err
}

func (c *Client) ReqServerConnect() error {
	_, err := c.conn.Write(ReqServerConnect(nil))
	return err
}

func (c *Client) ReqServerDisconnect() error {
	_, err := c.conn.Write(ReqServerConnect(nil))
	return err
}

func (c *Client) SetClientStats(on bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.conn.Write(SetClientStats(nil, on))
	c.clientStats = on
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
	if msg[0] != 'S' {
		return errors.New("not supported")
	}
	buf := cl.buf[:0]
	mark := 2
	for i := 2; i < len(msg); i++ {
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
		switch buf[0] {
		case "CURRENT PROTOCOL":
			cl.processProtocolMessage(buf)
		case "REGISTER CLIENT APP COMPLETED":
			cl.processClientAppMessage(ClientAppMessageTypeRegister)
		case "REMOVE CLIENT APP COMPLETED":
			cl.processClientAppMessage(ClientAppMessageTypeRemove)
		case "CURRENT LOGINID":
			cl.processLoginIdMessage(buf)
		case "CURRENT PASSWORD":
			cl.processPasswordMessage(buf)
		case "LOGIN INFO SAVED":
			cl.processLoginInfoMessage(LoginInfoMessageTypeSaved)
		case "LOGIN INFO NOT SAVED":
			cl.processLoginInfoMessage(LoginInfoMessageTypeNotSaved)
		case "AUTOCONNECT ON":
			cl.processAutoConnectMessage(true)
		case "AUTOCONNECT OFF":
			cl.processAutoConnectMessage(false)
		case "STATS":
			cl.processStatsMessage(buf)
		case "CLIENTSTATS":
			cl.processClientStatsMessage(buf)
		}
	}
	return nil
}

func (cl *Client) processProtocolMessage(values []string) {
	msg := ProtocolMessage{Version: atIndex(values, 1)}
	_ = msg
}

func (cl *Client) processClientAppMessage(t ClientAppMessageType) {
	msg := ClientAppMessage{Type: t}
	_ = msg
}

func (cl *Client) processLoginIdMessage(values []string) {
	msg := LoginIdMessage{UserLoginId: atIndex(values, 1)}
	_ = msg
}

func (cl *Client) processPasswordMessage(values []string) {
	msg := PasswordMessage{UserPassword: atIndex(values, 1)}
	_ = msg
}

func (cl *Client) processLoginInfoMessage(t LoginInfoMessageType) {
	msg := LoginInfoMessage{Type: t}
	_ = msg
}

func (cl *Client) processAutoConnectMessage(on bool) {
	msg := AutoConnectMessage{On: on}
	_ = msg
}

func (cl *Client) processStatsMessage(values []string) {
	msg := StatsMessage{}
	msg.Unmarshal(values)
	_ = msg
}

func (cl *Client) processClientStatsMessage(values []string) {
	msg := ClientStatsMessage{}

	msg.Unmarshal(values)
	_ = msg
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
