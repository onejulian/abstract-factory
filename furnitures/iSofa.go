package furnitures

type ISofa interface {
	SetName(name string)
	GetName() string
}

type Sofa struct {
	name string
}

func (s *Sofa) SetName(name string) {
	s.name = name
}

func (s *Sofa) GetName() string {
	return s.name
}