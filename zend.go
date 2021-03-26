package zend

import (
	"fmt"
)

type Client struct {
	Token			string
	Sender			string
}

func NewClient(Token string, Sender string) *Client {
	return &Client {
		Token: Token,
		Sender: Sender,
	}
}

func (c *Client) Message(Recipient string, Message string) (error) {

	fmt.Println("called to Message")

	return nil
}
