package zend

func NewClient(tkn string, sndr string) *Client {
	return &Client {
		Token: tkn,
		Sender: sndr,
	}
}
