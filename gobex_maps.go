package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//dict?

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	delete(m, "k2")
	fmt.Println("map:", m)

	//missing key or ker=0
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}
