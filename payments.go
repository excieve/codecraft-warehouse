package cdwarehouse

type Payment interface {
	IsComplete() bool
}

type FailingPayment struct{}

func (fp *FailingPayment) IsComplete() bool {
	return false
}

type SuccessfulPayment struct{}

func (sp *SuccessfulPayment) IsComplete() bool {
	return true
}
