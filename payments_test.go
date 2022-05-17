package cdwarehouse

import "github.com/stretchr/testify/mock"

type MockPayment struct {
	mock.Mock
}

func (m *MockPayment) IsComplete(amount float64) bool {
	args := m.Called(amount)
	return args.Bool(0)
}
