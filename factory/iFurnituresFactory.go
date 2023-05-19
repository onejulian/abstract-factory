package factory

import (
	"furniture-style/furnitures"
	"furniture-style/styles"
)

type IFurnituresFactory interface {
	CreateChair() furnitures.IChair
	CreateCoffeeTable() furnitures.ICoffeeTable
	CreateSofa() furnitures.ISofa
}

func GetFurnituresFactory(style string) IFurnituresFactory {
	switch style {
	case "artDeco":
		return new(styles.ArtDeco)
	case "modern":
		return new(styles.Modern)
	case "victorian":
		return new(styles.Victorian)
	default:
		return nil
	}
}
