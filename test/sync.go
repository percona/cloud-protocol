package test

import (
	"time"
	"errors"
	"encoding/json"
)

var ETimeout = errors.New("WaitRect timeout")

func WaitRecv(recvChan chan []byte, result interface{}) error {
	select {
	case data := <-recvChan:
		return json.Unmarshal(data, result)
	case <-time.After(500 * time.Millisecond):
	}
	return ETimeout
}

func WaitRecvWithTimeout(recvChan chan []byte, result interface{}, timeout time.Duration) error {
	select {
	case data := <-recvChan:
		return json.Unmarshal(data, result)
	case <-time.After(timeout):
	}
	return ETimeout
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
