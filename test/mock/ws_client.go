package mock

import (
	"log"
	"code.google.com/p/go.net/websocket"
)

type WebsocketClient struct {
	sendChan chan interface{}
	recvChan chan interface{}
	// --
	SendErr error
	RecvErr error
	// --
	conn *websocket.Conn
}

func NewWebsocketClient(sendChan chan interface{}, recvChan chan interface{}) *WebsocketClient {
	c := &WebsocketClient{
		sendChan: sendChan,
		recvChan: recvChan,
	}
	return c
}

func (c *WebsocketClient) Connect(url string, origin string) error {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		return err
	}
	c.conn = conn
	c.run()
	return nil
}

func (c *WebsocketClient) Close() {
	c.conn.Close()
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
			var data interface{}
			err := websocket.JSON.Receive(c.conn, &data)
			if err != nil {
				log.Printf("ERROR: websocket.JSON.Receive: %s\n", err)
				c.RecvErr = err
				break
			}
			c.recvChan <-data
		}
	}()
}
