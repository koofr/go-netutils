//
// Inspiration: https://gist.github.com/idada/9144886
//
package netutils

import (
	"net"
	"time"
)

type TimeoutConn struct {
	conn         net.Conn
	readTimeout  time.Duration
	writeTimeout time.Duration
}

func NewTimeoutConn(conn net.Conn, readTimeout time.Duration, writeTimeout time.Duration) *TimeoutConn {
	return &TimeoutConn{
		conn:         conn,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

func (c *TimeoutConn) Read(b []byte) (n int, err error) {
	if c.readTimeout > 0 {
		c.SetReadDeadline(time.Now().Add(c.readTimeout))
	}
	return c.conn.Read(b)
}

func (c *TimeoutConn) Write(b []byte) (n int, err error) {
	if c.writeTimeout > 0 {
		c.SetWriteDeadline(time.Now().Add(c.writeTimeout))
	}
	return c.conn.Write(b)
}

func (c *TimeoutConn) Close() error {
	return c.conn.Close()
}

func (c *TimeoutConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *TimeoutConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *TimeoutConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *TimeoutConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *TimeoutConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
