package zend


type (
	Client struct {
		Token			string
		Sender			string
		Version			string
		URI				string
	}

	Message struct {
		ID				uint64
	}

	OtpTime struct {
		Created			uint64
		Expire			uint64
	}

	Otp struct {
		ID				uint64
		Time			OtpTime
	}
)
