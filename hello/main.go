package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Store unordered values in maps")
	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	//delete(states, "WA")
	///fmt.Println(states)
	for k,v := range states {
		fmt.Printf("%v: %v:", k, v)
	}
	// extract the key as a slice to garanty the order
	Keys := make([]string, len(states))
	i := 0
	for k := range states {
		Keys[i] = k
		i++
	}
	fmt.Println(Keys)
	sort.Strings(Keys)
	fmt.Println(Keys)

	for i := range Keys {
		fmt.Println(states[Keys[i]])
	}

}
