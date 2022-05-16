package cdwarehouse

type Warehouse []string

func (w *Warehouse) Search(artist string, title string) []string {
	return []string{}
}

func NewWarehouse() *Warehouse {
	return &Warehouse{}
}
