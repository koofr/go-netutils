package netutils

import (
	"net"
	"time"
)

func Await(addr string, step time.Duration, totalTime time.Duration) (ok bool) {
	totalSteps := int(totalTime / step)

	for i := 0; i < totalSteps; i += 1 {
		conn, err := net.Dial("tcp", addr)

		if err == nil {
			conn.Close()
			return true
		}

		time.Sleep(step)
	}

	return false
}
