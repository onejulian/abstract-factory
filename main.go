package main

import (
	"fmt"
	"furniture-style/factory"
	"furniture-style/furnitures"
)

func main() {
	artDecoFactory := factory.GetFurnituresFactory("artDeco")
	modernFactory := factory.GetFurnituresFactory("modern")
	victorianFactory := factory.GetFurnituresFactory("victorian")

	artDecoChair := artDecoFactory.CreateChair()
	artDecoCoffeeTable := artDecoFactory.CreateCoffeeTable()
	artDecoSofa := artDecoFactory.CreateSofa()

	modernChair := modernFactory.CreateChair()
	modernCoffeeTable := modernFactory.CreateCoffeeTable()
	modernSofa := modernFactory.CreateSofa()

	victorianChair := victorianFactory.CreateChair()
	victorianCoffeeTable := victorianFactory.CreateCoffeeTable()
	victorianSofa := victorianFactory.CreateSofa()

	showChairDetails(artDecoChair)
	showCoffeeTableDetails(artDecoCoffeeTable)
	showSofaDetails(artDecoSofa)

	showChairDetails(modernChair)
	showCoffeeTableDetails(modernCoffeeTable)
	showSofaDetails(modernSofa)

	showChairDetails(victorianChair)
	showCoffeeTableDetails(victorianCoffeeTable)
	showSofaDetails(victorianSofa)
}

func showChairDetails(chair furnitures.IChair) {
	fmt.Printf("Chair name: %s\n\n", chair.GetName())
}

func showCoffeeTableDetails(coffeeTable furnitures.ICoffeeTable) {
	fmt.Printf("Coffee table name: %s\n\n", coffeeTable.GetName())
}

func showSofaDetails(sofa furnitures.ISofa) {
	fmt.Printf("Sofa name: %s\n\n", sofa.GetName())
}
