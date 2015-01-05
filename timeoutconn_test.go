package netutils

import (
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func startServer(t *testing.T) (net.Listener, int) {
	port, err := UnusedPort()
	if err != nil {
		t.Error(err)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			t.Error(err)
		}

		_, err = io.ReadAtLeast(conn, make([]byte, 10*1024*1024), 10*1024*1024)
		if err != nil {
			t.Error(err)
		}

		_, err = conn.Write([]byte{1})
		if err != nil {
			t.Error(err)
		}
	}()

	return ln, port
}

func TestConn(t *testing.T) {
	server, port := startServer(t)
	defer server.Close()

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	b := make([]byte, 10*1024*1024)

	_, err = conn.Write(b)
	if err != nil {
		t.Error(err)
	}

	_, err = conn.Read([]byte{1})
	if err != nil {
		t.Error(err)
	}

	written := make(chan bool)

	go func() {
		_, err = conn.Write(b)
		written <- true
	}()

	select {
	case <-written:
		t.Error("Data should not be written")
	case <-time.After(3 * time.Second):
	}
}

func TestTimeoutConn(t *testing.T) {
	server, port := startServer(t)
	defer server.Close()

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		t.Error(err)
	}

	conn = NewTimeoutConn(conn, 1*time.Second, 2*time.Second)

	b := make([]byte, 10*1024*1024)

	_, err = conn.Write(b)
	if err != nil {
		t.Error(err)
	}

	_, err = conn.Read([]byte{1})
	if err != nil {
		t.Error(err)
	}

	timeouted := make(chan bool)

	go func() {
		_, err = conn.Write(b)
		timeouted <- strings.Contains(err.Error(), "i/o timeout")
	}()

	select {
	case to := <-timeouted:
		if !to {
			t.Error("Write return timeout error")
		}
	case <-time.After(4 * time.Second):
		t.Error("Write should be interrupted")
	}
}
