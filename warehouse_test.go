package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerBuyCD(t *testing.T) {
	t.Run("Customer finds no CD with artist 'Foo' and title 'Bar'", func(t *testing.T) {
		warehouse := NewWarehouse([]Cd{})

		assert.NotNil(t, warehouse)

		cds := warehouse.Search("Foo", "Bar")

		assert.Equal(t, 0, len(cds))
	})

	t.Run("Customer finds one CD with artist 'Foo' and title 'Bar' that's not in stock", func(t *testing.T) {
		warehouse := NewWarehouse([]Cd{{"Foo", "Bar", 0}})

		assert.NotNil(t, warehouse)

		cds := warehouse.Search("Foo", "Bar")

		assert.Len(t, cds, 1)

		foundCd := cds[0]
		assert.Equal(t, "Foo", foundCd.Artist)
		assert.Equal(t, "Bar", foundCd.Title)

		assert.False(t, foundCd.InStock())
	})

	t.Run("Customer finds one CD with artist 'Foo' and title 'Bar' that's in stock", func(t *testing.T) {
		warehouse := NewWarehouse([]Cd{{"Foo", "Bar", 1}})

		assert.NotNil(t, warehouse)

		cds := warehouse.Search("Foo", "Bar")

		foundCd := cds[0]
		assert.Equal(t, "Foo", foundCd.Artist)
		assert.Equal(t, "Bar", foundCd.Title)

		assert.True(t, foundCd.InStock())

		t.Run("and fails to buy it", func(t *testing.T) {
			payment := &FailingPayment{}

			assert.False(t, foundCd.Buy(payment))
			assert.True(t, foundCd.InStock())
		})

		t.Run("and buys it", func(t *testing.T) {
			payment := &SuccessfulPayment{}

			assert.True(t, foundCd.Buy(payment))
			assert.False(t, foundCd.InStock())
		})
	})
}
