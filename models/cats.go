package models

import "time"

// Cat model represents table `cats`
type Cat struct {
	ID        uint64
	Name      string
	Color     CatColor
	Naughty   bool
	Dexterity int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CatColor is the enumeration of Cat.Color, part of Cat model
type CatColor uint8

// CatColor possible values
const (
	CatColorWhite CatColor = iota
	CatColorBlack
	CatColorOrange
	CatColorGreen
)

// ToString returns CatColor string value
func (c CatColor) ToString() string {
	values := map[CatColor]string{
		CatColorWhite:  "white",
		CatColorBlack:  "black",
		CatColorOrange: "orange",
		CatColorGreen:  "green",
	}
	return values[c]
}
