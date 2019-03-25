package helpers

// a helper id generator, should not be used in production it's just an explanation

var accountId = 0

func GetNextAccountId() int {
	accountId += 1
	return accountId
}

var accountEventId = 0

func GetNextAccountEventId() int {
	accountEventId += 1
	return accountEventId
}