package zend

func NewClient(token string, sender string) *Client {
	return &Client {
		Token: token,
		Sender: sender,
	}
}
