package furniture

import (
	"fmt"
	"new_go_house/House/Units"
)

type Furniture struct {
	Type          string
	Length        types.Centimeter
	Width         types.Centimeter
	Color         string
	ComfortRating int
}

func (f Furniture) Print() {
	fmt.Print("Тип мебели: ", f.Type, "\nДлина: ", f.Length, "\nШирина: ", f.Width, "\nЦвет: ", f.Color, "\nРейтинг комфорта: ", f.ComfortRating, "\n")
}