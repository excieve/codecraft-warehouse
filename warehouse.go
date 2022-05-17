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
	reviews []CdReview
}

type CdReview struct {
	Rating       int
	Comment      string
	CustomerName string
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

func (c *Cd) AddReview(customer *Customer, rating int, comment string) bool {
	if rating < 1 || rating > 10 {
		return false
	}

	if !customer.HasPurchased(c.id) {
		return false
	}

	review := CdReview{
		Rating:       rating,
		Comment:      comment,
		CustomerName: customer.Name,
	}

	c.reviews = append(c.reviews, review)

	return true
}

func NewCd(artist string, title string, stock int) *Cd {
	return &Cd{
		Artist:  artist,
		Title:   title,
		id:      uuid.NewString(),
		stock:   stock,
		reviews: []CdReview{},
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
