package furnitures

type ICoffeeTable interface {
	SetName(name string)
	GetName() string
}

type CoffeeTable struct {
	name string
}

func (c *CoffeeTable) SetName(name string) {
	c.name = name
}

func (c *CoffeeTable) GetName() string {
	return c.name
}