package zend

type Client struct {
	Token   string
	Sender  string
	Version string
	URI     string

	Otp *OtpService
}

func NewClient(token string, sender string) *Client {

	c := &Client{
		Token:   token,
		Sender:  sender,
		Version: "1.0",
		URI:     "https://api.zend.lk",
	}

	c.Otp = &OtpService{client: c}

	return c
}
