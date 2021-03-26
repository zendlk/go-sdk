package zend

import (
	"fmt"
)

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
