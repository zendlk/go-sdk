package zend

import (
	"fmt"
)

func (c *Client) Message(Recipient string, Message string) (error) {

	fmt.Println("called to Message")

	return nil
}
