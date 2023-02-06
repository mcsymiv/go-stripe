package card

type Card struct {
	Secret   string
	Key      string
	Currency string
}

type Transaction struct {
	// Id
	// holds transaction status id
	Id int

	// Amount
	// represents transaction amount, avoids float point
	Amount int

	// Currency
	// holds info on transaction currency in USD like format
	Currency string

	// Last4
	// holds card last four digits
	Last4 string

	// ReturnCode
	// holds transaction bank return code from Stripe API
	ReturnCode string
}
