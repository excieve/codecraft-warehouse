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

	t.Run("Customer finds one CD with artist 'Foo' and title 'Bar'", func(t *testing.T) {
		warehouse := NewWarehouse([]Cd{{"Foo", "Bar"}})

		assert.NotNil(t, warehouse)

		cds := warehouse.Search("Foo", "Bar")

		assert.Len(t, cds, 1)

		foundCd := cds[0]
		assert.Equal(t, "Foo", foundCd.Artist)
		assert.Equal(t, "Bar", foundCd.Title)
	})
}
