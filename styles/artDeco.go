package styles

import "furniture-style/furnitures"

type ArtDeco struct {
}

type ArtDecoChair struct {
	furnitures.Chair
}

type ArtDecoCoffeeTable struct {
	furnitures.CoffeeTable
}

type ArtDecoSofa struct {
	furnitures.Sofa
}

func (a *ArtDeco) CreateChair() furnitures.IChair {
	chair := new(furnitures.Chair)
	chair.SetName("Art Deco Chair")
	return &ArtDecoChair{
		Chair: *chair,
	}
}

func (a *ArtDeco) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(furnitures.CoffeeTable)
	coffeeTable.SetName("Art Deco Coffee Table")
	return &ArtDecoCoffeeTable{
		CoffeeTable: *coffeeTable,
	}
}

func (a *ArtDeco) CreateSofa() furnitures.ISofa {
	sofa := new(furnitures.Sofa)
	sofa.SetName("Art Deco Sofa")
	return &ArtDecoSofa{
		Sofa: *sofa,
	}
}
