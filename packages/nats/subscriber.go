package nats

import "github.com/nats-io/nats.go"

type Handler func(data []byte)

func (c *Client) Subscribe(subject string, handler Handler) error {
	_, err := c.conn.Subscribe(subject, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	return err
}
