package nats

func (c *Client) Publish(subject string, data []byte) error {
	return c.conn.Publish(subject, data)
}
