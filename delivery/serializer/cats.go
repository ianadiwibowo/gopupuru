package serializer

import "github.com/ianadiwibowo/scout-regiment/model"

// SerializedSimpleCat ...
type SerializedSimpleCat struct {
	ID        uint64 `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Naughty   bool   `json:"naughty"`
	Dexterity int    `json:"dexterity"`
}

// SimpleCat ...
func (s *Serializer) SimpleCat(cat model.Cat) *SerializedSimpleCat {
	return &SerializedSimpleCat{
		ID:        cat.ID,
		Name:      cat.Name,
		Color:     cat.Color.ToString(),
		Naughty:   cat.Naughty,
		Dexterity: cat.Dexterity,
	}
}
