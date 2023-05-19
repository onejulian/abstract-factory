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
	chair := new(ModernChair)
	chair.SetName("Modern Chair")
	return chair
}

func (m *Modern) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(ModernCoffeeTable)
	coffeeTable.SetName("Modern Coffee Table")
	return coffeeTable
}

func (m *Modern) CreateSofa() furnitures.ISofa {
	sofa := new(ModernSofa)
	sofa.SetName("Modern Sofa")
	return sofa
}