package cdwarehouse

type Warehouse []Cd

type Cd struct {
	Artist  string
	Title   string
	stock   int
	reviews []CdReview
}

type CdReview struct {
	Rating  int
	Comment string
}

func (c *Cd) InStock() bool {
	return c.stock > 0
}

func (c *Cd) Buy(payment Payment) bool {
	if c.stock < 1 {
		return false
	}

	if payment.IsComplete() {
		c.stock--
		return true
	}

	return false
}

func (c *Cd) AddReview(rating int, comment string) bool {
	if rating < 1 || rating > 10 {
		return false
	}

	review := CdReview{
		Rating:  rating,
		Comment: comment,
	}

	c.reviews = append(c.reviews, review)

	return true
}

func NewCd(artist string, title string, stock int) *Cd {
	return &Cd{
		Artist:  artist,
		Title:   title,
		stock:   stock,
		reviews: []CdReview{},
	}
}

func (w *Warehouse) Search(artist string, title string) []Cd {
	var cds = make([]Cd, 0)

	for _, cd := range *w {
		if cd.Title == title && cd.Artist == artist {
			cds = append(cds, cd)
		}
	}

	return cds
}

func NewWarehouse(cds []Cd) *Warehouse {
	warehouse := Warehouse(cds)
	return &warehouse
}
