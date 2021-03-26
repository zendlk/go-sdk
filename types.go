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
)
