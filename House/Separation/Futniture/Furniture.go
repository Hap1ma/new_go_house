package furniture

import (
	"fmt"
	units "new_go_house/House/Separation/Units"
)

type Furniture struct {
	Type          string
	Length        units.Centimeter
	Width         units.Centimeter
	Color         string
	ComfortRating int
}

func (f Furniture) Print() {
	fmt.Print("Тип мебели: ", f.Type, "\nДлина: ", f.Length, "\nШирина: ", f.Width, "\nЦвет: ", f.Color, "\nРейтинг комфорта: ", f.ComfortRating, "\n")
}
