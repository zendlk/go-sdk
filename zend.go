package zend

func NewClient(token string, sender string) *Client {
	return &Client {
		Token: token,
		Sender: sender,
		Version: "1.0",
		URI: "https://api.zend.lk",
	}
}
