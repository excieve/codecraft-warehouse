package cdwarehouse

type Warehouse []Cd

type Cd struct {
	Artist string
	Title  string
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
