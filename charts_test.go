package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCharts struct {
	mock.Mock
}

func (m *MockCharts) Notify(artist, title string, items int) error {
	args := m.Called(artist, title, items)
	return args.Error(0)
}

func TestCharts(t *testing.T) {
	t.Run("Customer buys a CD with artist 'Foo' and title 'Bar' and charts are notified about the purchase", func(t *testing.T) {
		cd := NewCd("Foo", "Bar", 3, 20.0)
		assert.NotNil(t, cd)

		customer := NewCustomer("tester")
		assert.NotNil(t, customer)

		payment := &SuccessfulPayment{}

		charts := new(MockCharts)

		charts.On("Notify", "Foo", "Bar", 1).Return(nil)

		assert.True(t, cd.Buy(customer, payment, charts))

		charts.AssertExpectations(t)
	})
}
