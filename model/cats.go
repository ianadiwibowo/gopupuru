package model

import "time"

// Cat model represents table `cats`
type Cat struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Color     CatColor  `json:"color"`
	Naughty   bool      `json:"naughty"`
	Dexterity int       `json:"dexterity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
