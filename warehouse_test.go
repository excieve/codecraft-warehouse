package cdwarehouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerBuyCD(t *testing.T) {
	t.Run("Customer finds no CD with artist 'Foo' and title 'Bar'", func(t *testing.T) {
		warehouse := NewWarehouse()

		assert.NotNil(t, warehouse)

		cds := warehouse.Search("Foo", "Bar")

		assert.Equal(t, 0, len(cds))
	})
}
