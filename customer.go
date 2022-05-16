package cdwarehouse

type Customer struct {
	name      string
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
		name:      name,
		purchases: map[string]struct{}{},
	}
}
