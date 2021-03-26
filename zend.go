package zend

type Client struct {
	Token			string
	Sender			string
}

func NewClient(tkn string, sndr string) *Client {
	return &Client {
		Token: tkn,
		Sender: sndr,
	}
}
