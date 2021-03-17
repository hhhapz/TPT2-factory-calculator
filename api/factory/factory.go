package factory

import (
	"encoding/json"
	"fmt"
)

const (
	TypeRecipe  = "recipe"
	TypeMachine = "machine"
	TypeShop    = "shop"
)

// Factory is the data structure that contains all of the information for the
// factory calculator.
type Factory struct {
	Parts map[string]Part
}

// Part is an interface that masks either a MachinePart, ShopPart, or a Recipe.
type Part interface {
	ItemType() string
	ItemName() string
	ItemDisplayName() string
	ImageURL() string
	CraftCost() int64
}

type rawPart struct {
	Type string
}

func (d *Factory) UnmarshalJSON(b []byte) error {
	d.Parts = make(map[string]Part)

	var partMap struct {
		Parts map[string]json.RawMessage
	}
	err := json.Unmarshal(b, &partMap)
	if err != nil {
		return err
	}

	for name, part := range partMap.Parts {
		raw := rawPart{}
		err := json.Unmarshal(part, &raw)
		if err != nil {
			return fmt.Errorf("error while parsing %s: %v", name, err)
		}

		switch raw.Type {
		default:
			return fmt.Errorf("unknown part type for %s: '%s'", name, raw.Type)
		case TypeRecipe:
			var recipe Recipe
			err := json.Unmarshal(part, &recipe)
			if err != nil {
				return err
			}
			Data.Parts[name] = recipe
		case TypeMachine:
			var machinePart MachinePart
			err := json.Unmarshal(part, &machinePart)
			if err != nil {
				return err
			}
			Data.Parts[name] = machinePart
		case TypeShop:
			var shopPart ShopPart
			err := json.Unmarshal(part, &shopPart)
			if err != nil {
				return err
			}
			Data.Parts[name] = shopPart
		}
	}

	return nil
}

// Recipe holds data for a Part that is crafted in Shape, with each rune's part
// defined in Parts.
type Recipe struct {
	Type        string
	Name        string
	DisplayName string
	Category    string
	URLPath     string
	Cost        int64
	Shape       []string
	Parts       map[string]string
}

// MachinePart is a factory part that is created when used inside Machine.
type MachinePart struct {
	Type        string
	Name        string
	DisplayName string
	URLPath     string
	Source      string
	Machine     string
	Produced    int
}

// ShopPart is a part that can be purchased in the game store.
type ShopPart struct {
	Type        string
	Name        string
	DisplayName string
	URLPath     string
	Cost        int64
}

func (r Recipe) ItemType() string        { return r.Type }
func (r Recipe) ItemName() string        { return r.Name }
func (r Recipe) ItemDisplayName() string { return r.DisplayName }
func (r Recipe) ImageURL() string        { return r.URLPath }
func (r Recipe) CraftCost() int64        { return r.Cost }

func (m MachinePart) ItemType() string        { return m.Type }
func (m MachinePart) ItemName() string        { return m.Name }
func (m MachinePart) ItemDisplayName() string { return m.DisplayName }
func (m MachinePart) ImageURL() string        { return m.URLPath }
func (m MachinePart) CraftCost() int64        { return 0 }

func (s ShopPart) ItemType() string        { return s.Type }
func (s ShopPart) ItemName() string        { return s.Name }
func (s ShopPart) ItemDisplayName() string { return s.DisplayName }
func (s ShopPart) ImageURL() string        { return s.URLPath }
func (s ShopPart) CraftCost() int64        { return s.Cost }
