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
	chair := new(furnitures.Chair)
	chair.SetName("Victorian Chair")
	return &VictorianChair{
		Chair: *chair,
	}
}

func (v *Victorian) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(furnitures.CoffeeTable)
	coffeeTable.SetName("Victorian Coffee Table")
	return &VictorianCoffeeTable{
		CoffeeTable: *coffeeTable,
	}
}

func (v *Victorian) CreateSofa() furnitures.ISofa {
	sofa := new(furnitures.Sofa)
	sofa.SetName("Victorian Sofa")
	return &VictorianSofa{
		Sofa: *sofa,
	}
}