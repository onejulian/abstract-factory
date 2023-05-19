package styles

import "furniture-style/furnitures"

type Victorian struct {
}

type VictorianChair struct {
	furnitures.Chair
}

type VictorianCoffeeTable struct {
	furnitures.CoffeeTable
}

type VictorianSofa struct {
	furnitures.Sofa
}

func (v *Victorian) CreateChair() furnitures.IChair {
	chair := new(VictorianChair)
	chair.SetName("Victorian Chair")
	return chair
}

func (v *Victorian) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(VictorianCoffeeTable)
	coffeeTable.SetName("Victorian Coffee Table")
	return coffeeTable
}

func (v *Victorian) CreateSofa() furnitures.ISofa {
	sofa := new(VictorianSofa)
	sofa.SetName("Victorian Sofa")
	return sofa
}