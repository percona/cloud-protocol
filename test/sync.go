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

func WaitConsume(msgChan chan string) []string {
	var buf []string
	var haveData bool = true
	for haveData {
		select {
		case msg := <-msgChan:
			buf = append(buf, msg)
		case <-time.After(100 * time.Millisecond):
			haveData = false
		}
	}
	return buf
}
