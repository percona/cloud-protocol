package test

import (
	"os"
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

func WaitFiles(dir string) []os.FileInfo {
	for i := 0; i < 3; i++ {
		files, _ := ioutil.ReadDir(dir)
		if len(files) > 0 {
			return files
		}
		time.Sleep(100 * time.Millisecond)
	}
	files, _ := ioutil.ReadDir(dir)
	return files
}

func WaitPost(postChan chan []byte) []byte {
	select {
	case data := <-postChan:
		return data
	case <-time.After(50 * time.Millisecond):
		return nil
	}
}
