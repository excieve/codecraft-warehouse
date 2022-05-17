package cdwarehouse

type Payment interface {
	IsComplete(amount float64) bool
}

type FailingPayment struct{}

func (fp *FailingPayment) IsComplete(float64) bool {
	return false
}

type SuccessfulPayment struct{}

func (sp *SuccessfulPayment) IsComplete(float64) bool {
	return true
}
