package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTop100Title struct {
	artist      string
	title       string
	lowestPrice float64
}

type MockCharts struct {
	mock.Mock

	top100 []MockTop100Title
}

func NewMockCharts(top100 []MockTop100Title) *MockCharts {
	return &MockCharts{top100: top100}
}

func (m *MockCharts) GetLowestPrice(artist string, title string) float64 {
	for _, cd := range m.top100 {
		if cd.artist == artist && cd.title == title {
			return cd.lowestPrice
		}
	}

	return 0.0
}

func (m *MockCharts) IsTop100(artist, title string) bool {
	for _, cd := range m.top100 {
		if cd.artist == artist && cd.title == title {
			return true
		}
	}

	return false
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

		payment := new(MockPayment)
		payment.On("IsComplete", 20.0).Return(true)

		charts := NewMockCharts([]MockTop100Title{})
		charts.On("Notify", "Foo", "Bar", 1).Return(nil)

		assert.True(t, cd.Buy(customer, payment, charts))

		payment.AssertExpectations(t)
		charts.AssertExpectations(t)
	})

	t.Run("A CD with artist 'Not' and title 'Top100' in not found in the charts Top100, original price is offered", func(t *testing.T) {
		cd := NewCd("Not", "Top100", 3, 20.0)
		assert.NotNil(t, cd)

		top100 := []MockTop100Title{
			{"Foo", "Bar", 18.0},
		}
		charts := NewMockCharts(top100)

		assert.Equal(t, 20.0, cd.getFinalPrice(charts))
	})

	t.Run("A CD with artist 'Foo' and title 'Baz' is found in the charts Top100, offering a price of 17.0", func(t *testing.T) {
		cd := NewCd("Foo", "Baz", 3, 20.0)
		assert.NotNil(t, cd)

		top100 := []MockTop100Title{
			{"Foo", "Bar", 18.0},
			{"Foo", "Baz", 19.0},
		}
		charts := NewMockCharts(top100)

		assert.Equal(t, 18.0, cd.getFinalPrice(charts))
	})
}
