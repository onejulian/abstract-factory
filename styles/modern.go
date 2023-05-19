package styles

import "furniture-style/furnitures"

type Modern struct {
}

type ModernChair struct {
	furnitures.Chair
}

type ModernCoffeeTable struct {
	furnitures.CoffeeTable
}

type ModernSofa struct {
	furnitures.Sofa
}

func (m *Modern) CreateChair() furnitures.IChair {
	chair := new(furnitures.Chair)
	chair.SetName("Modern Chair")
	return &ModernChair{
		Chair: *chair,
	}
}

func (m *Modern) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(furnitures.CoffeeTable)
	coffeeTable.SetName("Modern Coffee Table")
	return &ModernCoffeeTable{
		CoffeeTable: *coffeeTable,
	}
}

func (m *Modern) CreateSofa() furnitures.ISofa {
	sofa := new(furnitures.Sofa)
	sofa.SetName("Modern Sofa")
	return &ModernSofa{
		Sofa: *sofa,
	}
}