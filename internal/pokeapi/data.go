package pokeapi

import (
	"fmt"
)

func ParseStruct(pokemon *Pokemon, id *string)  {
	fmt.Printf("Name: %s\n", *id)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	
	
	stat := pokemon.Stats
	fmt.Println("Stats:")
	for j := 0; j < len(stat); j++ {
		fmt.Printf("\t-%s: %v\n", stat[j].Stat.Name, stat[j].BaseStat)
	}

	types := pokemon.Types
	fmt.Println("Types:")
	for i := 0; i < len(types); i++ {
		fmt.Printf("\t- %s\n", types[i].Type.Name)
	}
}	