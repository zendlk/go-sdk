package zend

import "net/http"

func (c *Config) New() *Client {
	return Client
}
