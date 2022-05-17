package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerBuyCD(t *testing.T) {
	t.Run("Customer finds no CD with artist 'Foo' and title 'Bar'", func(t *testing.T) {
		warehouse := NewWarehouse([]Cd{})

		assert.NotNil(t, warehouse)

		foundCd := warehouse.Search("Foo", "Bar")

		assert.Nil(t, foundCd)
	})

	t.Run("Customer finds one CD with artist 'Foo' and title 'Bar' that's not in stock", func(t *testing.T) {
		fooCd := NewCd("Foo", "Bar", 0)
		warehouse := NewWarehouse([]Cd{*fooCd})

		assert.NotNil(t, warehouse)

		foundCd := warehouse.Search("Foo", "Bar")
		assert.NotNil(t, foundCd)

		assert.Equal(t, "Foo", foundCd.Artist)
		assert.Equal(t, "Bar", foundCd.Title)

		assert.False(t, foundCd.InStock())

		t.Run("and fails to buy it", func(t *testing.T) {
			customer := NewCustomer("tester")
			payment := &SuccessfulPayment{}

			assert.False(t, foundCd.Buy(customer, payment))
			assert.False(t, foundCd.InStock())
		})
	})

	t.Run("Customer finds one CD with artist 'Foo' and title 'Bar' that's in stock", func(t *testing.T) {
		fooCd := NewCd("Foo", "Bar", 1)
		warehouse := NewWarehouse([]Cd{*fooCd})

		assert.NotNil(t, warehouse)

		foundCd := warehouse.Search("Foo", "Bar")
		assert.NotNil(t, foundCd)

		assert.Equal(t, "Foo", foundCd.Artist)
		assert.Equal(t, "Bar", foundCd.Title)

		assert.True(t, foundCd.InStock())

		t.Run("and fails to buy it", func(t *testing.T) {
			customer := NewCustomer("tester")
			payment := &FailingPayment{}

			assert.False(t, foundCd.Buy(customer, payment))
			assert.True(t, foundCd.InStock())

			t.Run("failing to leave a review", func(t *testing.T) {
				assert.False(t, foundCd.Reviews.AddReview(customer, 1, "awful"))
			})
		})

		t.Run("and buys it", func(t *testing.T) {
			customer := NewCustomer("tester")
			payment := &SuccessfulPayment{}

			assert.True(t, foundCd.Buy(customer, payment))
			assert.False(t, foundCd.InStock())

			t.Run("failing to leave a review", func(t *testing.T) {
				assert.False(t, foundCd.Reviews.AddReview(customer, 15, "terrific"))
			})

			t.Run("leaving a review successfully", func(t *testing.T) {
				assert.True(t, foundCd.Reviews.AddReview(customer, 2, "not great"))

				assert.Len(t, foundCd.Reviews.reviews, 1)

				assert.Equal(t, 2, foundCd.Reviews.reviews[0].Rating)
				assert.Equal(t, "not great", foundCd.Reviews.reviews[0].Comment)
				assert.Equal(t, customer.Name, foundCd.Reviews.reviews[0].CustomerName)
			})

			t.Run("label adds an item to the stock", func(t *testing.T) {
				assert.Equal(t, 2, foundCd.AddStock(2))
				assert.True(t, foundCd.InStock())
			})
		})
	})
}
