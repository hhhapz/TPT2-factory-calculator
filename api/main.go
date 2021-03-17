package main

import (
	"fmt"
	"log"

	"github.com/hhhapz/tpt2-factory-calculator/api/factory"
)

func main() {
	err := factory.Parse("./recipes.json")
	if err != nil {
		log.Fatal(err)
	}

	part := factory.Data.Parts["producer.exoticgems.t1"]
	raw, err := factory.CalculatePartPrice(part)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cost for exotic producer:")
	fmt.Printf("Total rubber: %d\n", raw.Rubber)
	fmt.Printf("Total hammers: %d\n", raw.Hammer)
	for i, n := range raw.Tier {
		fmt.Printf("Tier %2d dust: %.2f\n", i+1, n)
	}
}
