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
	chair := new(ArtDecoChair)
	chair.SetName("Art Deco Chair")
	return chair
}

func (a *ArtDeco) CreateCoffeeTable() furnitures.ICoffeeTable {
	coffeeTable := new(ArtDecoCoffeeTable)
	coffeeTable.SetName("Art Deco Coffee Table")
	return coffeeTable
}

func (a *ArtDeco) CreateSofa() furnitures.ISofa {
	sofa := new(ArtDecoSofa)
	sofa.SetName("Art Deco Sofa")
	return sofa
}
