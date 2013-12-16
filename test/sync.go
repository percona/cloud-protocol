package test

import (
	"time"
)

func WaitRecv(recvChan chan interface{}) interface{} {
	select {
	case data := <-recvChan:
		return data
	case <-time.After(500 * time.Millisecond):
	}
	return nil
}
