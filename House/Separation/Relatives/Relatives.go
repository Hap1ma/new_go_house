package relatives

import (
	"fmt"
	family "new_go_house/House/Separation/Relatives/Family"
)

type Familia struct {
	FamilyMembers []family.Relatives
	Surname       string
}

func (f Familia) Print() {
	fmt.Print("Семья: ", f.Surname, "\n")
	for _, Relatives := range f.Relatives {
		Relatives.Print()
	}
}
