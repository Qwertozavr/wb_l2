package pattern

//package main

import "fmt"

type Apartment struct {
	Bedrooms  int
	Bathrooms int
	Hall      int
}

func (a *Apartment) Show() {
	fmt.Printf("\n--------------\nBed Rooms: %d\nBath Rooms: %d\nHalls: %d\n--------------", a.Bedrooms, a.Bathrooms, a.Hall)
}

type Builder interface {
	buildBedrooms()
	buildBathrooms()
	buildHall()
	MakeApart() Apartment
}

// Однокомнатная квартира
type OneRoomed struct {
	Apartment
}

func (o *OneRoomed) buildBedrooms() {
	o.Bedrooms = 1
}

func (o *OneRoomed) buildBathrooms() {
	o.Bathrooms = 1
}

func (o *OneRoomed) buildHall() {
	o.Hall = 1
}

func (o *OneRoomed) MakeApart() Apartment {
	return Apartment{
		Bedrooms:  o.Bedrooms,
		Bathrooms: o.Bathrooms,
		Hall:      o.Hall,
	}
}

// Двухкомнатная квартира
type TwoRoomed struct {
	Apartment
}

func (t *TwoRoomed) buildBedrooms() {
	t.Bedrooms = 2
}

func (t *TwoRoomed) buildBathrooms() {
	t.Bathrooms = 2
}

func (t *TwoRoomed) buildHall() {
	t.Hall = 1
}

func (t *TwoRoomed) MakeApart() Apartment {
	return Apartment{
		Bedrooms:  t.Bedrooms,
		Bathrooms: t.Bathrooms,
		Hall:      t.Hall,
	}
}

// Выбор шаблона для постройки
func MakeBuilder(b string) Builder {
	switch b {
	case "OneRoom":
		return &OneRoomed{}
	case "TwoRooms":
		return &TwoRoomed{}
	default:
		panic("invalid builder")
	}
}

type ConcreteBuilder struct {
	Builder Builder
}

func NewConc(builder Builder) *ConcreteBuilder {
	return &ConcreteBuilder{
		Builder: builder,
	}
}

func (c *ConcreteBuilder) ChangeBuilder(builder Builder) {
	c.Builder = builder
}

func (c *ConcreteBuilder) Build() Apartment {
	c.Builder.buildBedrooms()
	c.Builder.buildBathrooms()
	c.Builder.buildHall()
	return c.Builder.MakeApart()
}

// func main() {
// 	oneRoomBuilder := MakeBuilder("OneRoom")
// 	twoRoomBuilder := MakeBuilder("TwoRooms")

// 	conc := NewConc(oneRoomBuilder)

// 	apartment := conc.Build()
// 	apartment.Show()

// 	conc.ChangeBuilder(twoRoomBuilder)
// 	apartment = conc.Build()
// 	apartment.Show()
// }
