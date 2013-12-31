package mock

type HttpClient struct {
	PostChan chan []byte
}

func (c *HttpClient) Post(url string, data []byte) error {
	c.PostChan <-data
	return nil
}
