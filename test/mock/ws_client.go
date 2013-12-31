package mock

import (
	"log"
	"code.google.com/p/go.net/websocket"
)

type WebsocketClient struct {
	sendChan chan interface{}
	recvChan chan []byte
	// --
	SendErr error
	RecvErr error
	// --
	conn *websocket.Conn
}

func NewWebsocketClient(sendChan chan interface{}, recvChan chan []byte) *WebsocketClient {
	c := &WebsocketClient{
		sendChan: sendChan,
		recvChan: recvChan,
	}
	return c
}

func (c *WebsocketClient) Connect() error {
	return nil
}

func (c *WebsocketClient) Disconnect() error {
	return nil
}

func (c *WebsocketClient) run() {
	go func() {
		for data := range c.sendChan {
			err := websocket.JSON.Send(c.conn, data)
			if err != nil {
				log.Printf("ERROR: websocket.JSON.Send: %s\n", err)
				c.SendErr = err
				break
			}
		}
	}()

	go func() {
		for {
			var data []byte
			err := websocket.Message.Receive(c.conn, &data)
			if err != nil {
				log.Printf("ERROR: websocket.Message.Receive: %s\n", err)
				c.RecvErr = err
				break
			}
			c.recvChan <- data
		}
	}()
}
