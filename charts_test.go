package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCharts struct {
	mock.Mock
}

func (m *MockCharts) GetLowestPrice(artist string, title string) float64 {
	if artist == "Foo" && title == "Bar" {
		return 18.0
	}

	return 0.0
}

func (m *MockCharts) IsTop100(artist, title string) bool {
	return artist == "Foo" && title == "Bar"
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

	t.Run("A CD with artist 'Not' and title 'Top100' in not found in the charts Top100, original price is offered", func(t *testing.T) {
		cd := NewCd("Not", "Top100", 3, 20.0)
		assert.NotNil(t, cd)

		charts := new(MockCharts)

		assert.Equal(t, 20.0, cd.getFinalPrice(charts))
	})

	t.Run("A CD with artist 'Foo' and title 'Bar' is found in the charts Top100, offering a price of 17.0", func(t *testing.T) {
		cd := NewCd("Foo", "Bar", 3, 20.0)
		assert.NotNil(t, cd)

		charts := new(MockCharts)

		assert.Equal(t, 17.0, cd.getFinalPrice(charts))
	})
}
