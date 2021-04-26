package zend

type Client struct {
	Token   string
	Sender  string
	Version string
	URI     string

	Otp     *OtpService
	Message *MessageService
}

func NewClient(token string, sender string) *Client {

	c := &Client{
		Token:   token,
		Sender:  sender,
		Version: "1.0",
		URI:     "https://api.zend.lk",
	}

	c.Otp = &OtpService{client: c}
	c.Message = &MessageService{client: c}

	return c
}
