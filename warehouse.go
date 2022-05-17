package cdwarehouse

import (
	"github.com/google/uuid"
)

type Warehouse []Cd

type Cd struct {
	Artist string
	Title  string

	id      string
	stock   int

	Reviews *CdReviews
}

func (c *Cd) InStock() bool {
	return c.stock > 0
}

func (c *Cd) Buy(customer *Customer, payment Payment) bool {
	if c.stock < 1 {
		return false
	}

	if payment.IsComplete() {
		c.stock--

		customer.AddPurchaseID(c.id)

		return true
	}

	return false
}

func (c *Cd) AddStock(items int) int {
	c.stock += items
	return c.stock
}

func NewCd(artist string, title string, stock int) *Cd {
	id := uuid.NewString()

	return &Cd{
		Artist:  artist,
		Title:   title,
		id:      id,
		stock:   stock,
		Reviews: NewCdReviews(id),
	}
}

func (w Warehouse) Search(artist string, title string) *Cd {
	for _, cd := range w {
		if cd.Title == title && cd.Artist == artist {
			return &cd
		}
	}

	return nil
}

func NewWarehouse(cds []Cd) Warehouse {
	return cds
}
