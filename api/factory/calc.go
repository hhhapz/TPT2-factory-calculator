package factory

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Raw is used to describe the amount of dust required for each tier, and cost
// in factory resource.
type Raw struct {
	Cost   int64
	Rubber int
	Hammer int
	Tier   [10]float64
}

func (r *Raw) Add(r1 *Raw) {
	r.Cost += r1.Cost
	r.Rubber += r1.Rubber
	r.Hammer += r1.Hammer
	for i, t := range r1.Tier {
		r.Tier[i] += t
	}
}

func (r *Raw) Times(n int) {
	if n == 1 {
		return
	}

	r.Cost *= int64(n)
	r.Rubber *= n
	r.Hammer *= n

	for i, t := range r.Tier {
		r.Tier[i] = t * float64(n)
	}
}

// CalculatePartPrice converts a part to the amount of dust needed for crafting
// it.
func CalculatePartPrice(p Part) (*Raw, error) {
	switch v := p.(type) {
	default:
		return nil, fmt.Errorf("unknown type: %T", v)
	case Recipe:
		return CalculateRecipePrice(v)
	case MachinePart:
		return CalculateMachinePartPrice(v)
	case ShopPart:
		return CalculateShopPartPrice(v)
	}
}

func CalculateRecipePrice(r Recipe) (*Raw, error) {
	raw := &Raw{}
	raw.Cost += r.Cost

	amount := make(map[string]int)

	for _, row := range r.Shape {
		for _, part := range row {
			if part == ' ' {
				continue
			}

			item, ok := r.Parts[string(part)]
			if !ok {
				log.Fatalf("could not find %c in parts (recipe: %s): ", part, r.Name)
			}
			amount[item]++
		}
	}

	for part, n := range amount {
		r1, err := CalculatePartPrice(Data.Parts[part])
		if err != nil {
			return nil, err
		}

		r1.Times(n)
		raw.Add(r1)
	}

	return raw, nil
}

func CalculateMachinePartPrice(m MachinePart) (*Raw, error) {
	raw := &Raw{}

	from, ok := Data.Parts[m.Source]
	if !ok {
		if !strings.HasPrefix(m.Source, "ingot") {
			return nil, fmt.Errorf("unknown machine part source: %s", m.Source)
		}

		split := strings.Split(m.Source, ".")
		tierStr := split[len(split)-1][1:]
		tier, err := strconv.Atoi(tierStr)
		if err != nil {
			return nil, fmt.Errorf("unknown tier for ingot tier: %s" + m.Source)
		}

		switch m.Produced {
		default:
			raw.Tier[tier] = 1.0 / float64(m.Produced)
		case 1:
			raw.Tier[tier] = 1
		case 2:
			raw.Tier[tier] = 0.5
		case 3:
			raw.Tier[tier] = 0.25
		}

		return raw, nil
	}

	r1, err := CalculatePartPrice(from)
	if err != nil {
		return nil, err
	}
	r1.Times(m.Produced)
	raw.Add(r1)

	return raw, err
}

func CalculateShopPartPrice(s ShopPart) (*Raw, error) {
	raw := &Raw{}
	raw.Cost = s.Cost

	switch s.Name {
	default:
		return nil, fmt.Errorf("unknown shop item: %s", s.Name)
	case "rubber":
		raw.Rubber = 1
	case "hammer":
		raw.Hammer = 1
	}

	return raw, nil
}
