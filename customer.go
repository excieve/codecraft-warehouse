package cdwarehouse

type Customer struct {
	Name      string
	purchases map[string]struct{}
}

func (c *Customer) AddPurchaseID(id string) {
	if id != "" {
		c.purchases[id] = struct{}{}
	}
}

func (c *Customer) HasPurchased(id string) bool {
	_, found := c.purchases[id]
	return found
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name:      name,
		purchases: map[string]struct{}{},
	}
}
