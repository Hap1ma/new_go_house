package relatives

import (
	"fmt"
	"new_go_house/House/Relatives/Family"
)

type Family struct {
	FamilyMembers []family_members.FamilyMembers
	Surname       string
}

func (f Family) Print() {
	fmt.Print("Семья: ", f.Surname, "\n")
	for _, familyMembers := range f.FamilyMembers {
		familyMembers.Print()
	}
}