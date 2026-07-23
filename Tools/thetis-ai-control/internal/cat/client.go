// Package cat implements a client for Thetis's CAT-over-TCP control channel
// (Project Files/Source/Console/CAT/TCPIPcatServer.cs), a Kenwood-style ASCII
// protocol: "<CMD><params>;" requests, ";"-terminated replies.
package cat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Client is a connection to a Thetis CAT-over-TCP server.
type Client struct {
	conn    net.Conn
	r       *bufio.Reader
	mu      sync.Mutex
	timeout time.Duration
}

// Dial connects to a Thetis CAT server at addr (host:port, e.g. "192.168.1.50:13013").
func Dial(addr string, timeout time.Duration) (*Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, fmt.Errorf("cat: dial %s: %w", addr, err)
	}
	return newClient(conn, timeout), nil
}

func newClient(conn net.Conn, timeout time.Duration) *Client {
	return &Client{conn: conn, r: bufio.NewReader(conn), timeout: timeout}
}

// Close closes the underlying connection.
func (c *Client) Close() error {
	return c.conn.Close()
}

// Send writes "<raw>;" without waiting for a reply. Use for set commands that
// Thetis does not echo back (most sets on this server are fire-and-forget).
func (c *Client) Send(raw string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.write(raw)
}

// Do writes "<raw>;" and returns the next ";"-terminated reply (without the
// trailing ";"). Use for query commands that always produce a reply.
func (c *Client) Do(raw string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if err := c.write(raw); err != nil {
		return "", err
	}
	return c.readReply()
}

func (c *Client) write(raw string) error {
	if c.timeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.timeout))
	}
	if _, err := c.conn.Write([]byte(raw + ";")); err != nil {
		return fmt.Errorf("cat: write %q: %w", raw, err)
	}
	return nil
}

func (c *Client) readReply() (string, error) {
	if c.timeout > 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.timeout))
	}
	s, err := c.r.ReadString(';')
	if err != nil {
		return "", fmt.Errorf("cat: read: %w", err)
	}
	return strings.TrimSuffix(s, ";"), nil
}

// Set sends a "<code><params>;" set command and does not wait for a reply.
func (c *Client) Set(code, params string) error {
	return c.Send(code + params)
}

// Query sends "<code>;" and returns the payload with the code prefix stripped.
func (c *Client) Query(code string) (string, error) {
	reply, err := c.Do(code)
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(reply, code) {
		return "", fmt.Errorf("cat: query %q: unexpected reply %q", code, reply)
	}
	return strings.TrimPrefix(reply, code), nil
}
