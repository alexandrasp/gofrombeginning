package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

//returning multiples values

func vals(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals(3, 2)
	fmt.Println(a)
	fmt.Println(b)

	c, _ := vals(1, 4)
	fmt.Println(c)

	_, d := vals(5, 3)
	fmt.Println(d)

	c, d = vals(2, 3)
	fmt.Println(c, d)

}
