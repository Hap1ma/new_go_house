package house

import (
	"fmt"
	devices "new_go_house/House/Separation/Devices"
	furniture "new_go_house/House/Separation/Futniture"
	relatives "new_go_house/House/Separation/Relatives"
	family "new_go_house/House/Separation/Relatives/Family"
	units "new_go_house/House/Separation/Units"
)

type House struct {
	Family        []relatives.Family
	Devices       []devices.Devices
	Furniture     []furniture.Furniture
	Rooms         int
	Square        units.SquareMeter
	CeilingHeight units.Centimeter
}

func (h House) Print() {
	fmt.Print("Комнаты: ", h.Rooms, "\nПлощадь: ", h.Square, "\nВысота потолков: ", h.CeilingHeight, "\n")
	for _, families := range h.Family {
		families.Print()
	}
	for _, tmpDevices := range h.Devices {
		tmpDevices.Print()
	}
	for _, tmpFurniture := range h.Furniture {
		tmpFurniture.Print()
	}
}
func Make() House {
	Table := furniture.Furniture{
		Type:          "Круглый стол",
		Length:        140,
		Width:         50,
		Color:         "White",
		ComfortRating: 1,
	}
	Сhair := furniture.Furniture{
		Type:          "Стул",
		Length:        30,
		Width:         30,
		Color:         "Зеленый",
		ComfortRating: 1,
	}
	Bed := furniture.Furniture{
		Type:          "Кровать",
		Length:        200,
		Width:         180,
		Color:         "Серый",
		ComfortRating: 9999,
	}
	Wardrobe := furniture.Furniture{
		Type:          "Шкаф",
		Length:        400,
		Width:         80,
		Color:         "Белый",
		ComfortRating: 2,
	}
	Sofa := furniture.Furniture{
		Type:          "Диван",
		Length:        30,
		Width:         10,
		Color:         "Белый",
		ComfortRating: 1000,
	}

	Me := family.FamilyMembers{
		Sex:          true,
		Age:          0,
		Name:         "Я",
		FamilyStatus: false,
		Children:     10,
	}
	FroggySonya := family.FamilyMembers{
		Sex:          true,
		Age:          0,
		Name:         "Жабий Соня",
		FamilyStatus: true,
		Children:     10,
	}
	FroggyDima := family.FamilyMembers{
		Sex:          true,
		Age:          1,
		Name:         "Жабий Дима",
		FamilyStatus: true,
		Children:     10,
	}
	Capybara := family.FamilyMembers{
		Sex:          false,
		Age:          0,
		Name:         "Капибара Дима(или Соня)",
		FamilyStatus: false,
		Children:     0,
	}
	BillyHerrington := family.FamilyMembers{
		Sex:          false,
		Age:          777,
		Name:         "Дакимакура",
		FamilyStatus: false,
		Children:     0,
	}
	Shark := family.FamilyMembers{
		Sex:          true,
		Age:          1,
		Name:         "Петя",
		FamilyStatus: false,
		Children:     0,
	}
	Dinosaur := family.FamilyMembers{
		Sex:          true,
		Age:          1,
		Name:         "Игорян",
		FamilyStatus: false,
		Children:     0,
	}

	MacBook := devices.Devices{
		Type:         "Ноутбук",
		Length:       31,
		Width:        22,
		Color:        "Золотистый",
		VoiceControl: true,
	}
	iPhone := devices.Devices{
		Type:         "Телефон",
		Length:       17,
		Width:        8,
		Color:        "Черный",
		VoiceControl: true,
	}
	Playstation := devices.Devices{
		Type:         "Приставка",
		Length:       30,
		Width:        33,
		Color:        "Черный",
		VoiceControl: false,
	}
	Watch := devices.Devices{
		Type:         "Настенные часы(большие, фигурные)",
		Length:       60,
		Width:        1,
		Color:        "Черный",
		VoiceControl: false,
	}
	Speaker := devices.Devices{
		Type:         "Колонка",
		Length:       255,
		Width:        196,
		Color:        "Белый",
		VoiceControl: true,
	}
	house := House{
		Family: []relatives.Family{relatives.Family{
			FamilyMembers: []family.FamilyMembers{Me, FroggyDima, FroggySonya, Capybara, BillyHerrington, Shark, Dinosaur},
			Surname:       "моя хата",
		},
		},
		Devices:       []devices.Devices{MacBook, iPhone, Playstation, Watch, Speaker},
		Furniture:     []furniture.Furniture{Table, Chair, Wardrobe, Bed, Sofa},
		Rooms:         1,
		Square:        35,
		CeilingHeight: 300,
	}
	return house
}
