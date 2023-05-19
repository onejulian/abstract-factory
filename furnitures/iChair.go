package furnitures

type IChair interface {
	SetName(name string)
	GetName() string
}

type Chair struct {
	name string
}

func (c *Chair) SetName(name string) {
	c.name = name
}

func (c *Chair) GetName() string {
	return c.name
}
