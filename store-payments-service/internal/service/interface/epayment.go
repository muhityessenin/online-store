package interfaces

type EPaymentService interface {
	MakePayment(amount float64) (string, error)
}
