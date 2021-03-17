package factory

// Breakdown is a representation of the direct children, dust tier costs,
// and machine parts needed.
type Breakdown struct {
	Source       Part
	Children     map[Part]int
	Raw          *Raw
	MachineParts map[Part]int
}

func BreakdownPart(p Part) (Breakdown, error) {
	// TO DO: Write this function lol
	return Breakdown{}, nil
}
