package netutils

import (
	"net"
)

func UnusedPort() (port int, err error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return
	}

	defer l.Close()

	port = l.Addr().(*net.TCPAddr).Port

	return
}
