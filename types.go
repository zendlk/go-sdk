package zend

type (
	Message struct {
		ID uint64
	}

	Times struct {
		Created uint64
		Updated uint64
		Deleted uint64
		Expires uint64
	}

	Otp struct {
		ID   uint64
		Time Times
	}
)
