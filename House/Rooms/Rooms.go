package rooms

import (
	"fmt"
	"math/rand"
)

type Room struct {
	Name   string
	Length float32
	Width  float32
	Height float32
	Square float32
}

func (r *Room) calculateSquare() {
	r.Square = r.Length * r.Width
}

func CreateRoom(name string) Room {
	rm := Room{Name: name,
		Length: rand.Float32()*(10-2) + 2,
		Width:  rand.Float32()*(7-2) + 2,
		Height: 2.8,
	}
	rm.calculateSquare()

	fmt.Println(rm)
	return rm
}

type Kitchen struct {
	Room      Room
	Furniture furniture.KitchenCreate
}

type Bedroom struct {
	Room      Room
	Furniture furniture.BedroomCreate
}

type Bathroom struct {
	Room      Room
	Furniture furniture.BathroomCreate
}

func CreateKitchen(name string) Kitchen {
	room := CreateRoom(name)
	kitchen := Kitchen{Room: room, Furniture: furniture.Kitchen()}
	return kitchen
}

func CreateBedroom(name string) Bedroom {
	room := CreateRoom(name)
	bedroom := Bedroom{Room: room, Furniture: furniture.Bedroom()}
	return bedroom
}

func CreateBathroom(name string) Bathroom {
	room := CreateRoom(name)
	bathroom := Bathroom{Room: room, Furniture: furniture.Bathroom()}
	return bathroom
}
