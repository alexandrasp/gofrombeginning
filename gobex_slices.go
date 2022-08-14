package main

import "fmt"

func main() {

	s := make([]string, 3)
	//empty
	fmt.Println("empty:", s)

	//you can set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	//tunned array
	s = append(s, "d")
	fmt.Println("set:", s)
	s = append(s, "e", "f")

	//auto copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//slices
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl5:", l)

	l = s[2:]
	fmt.Println("sl2:", l)

}
