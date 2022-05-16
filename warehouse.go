package cdwarehouse

type Warehouse []Cd

type Cd struct {
	Artist string
	Title  string
	stock  int
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
