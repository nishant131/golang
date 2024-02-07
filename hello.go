package main

import (
	"fmt"
	"math"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(math.Pi)
	var i = "2"
	const c = 'a'
	j := 3
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("type: %T\n", c)
	fmt.Println(string(c))
	fmt.Println(quote.Go())
}
