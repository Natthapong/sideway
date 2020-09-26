package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Go!")
	fruits := []string{"Apple", "Orange", "Plum", "Durian", "Papaya"}
	some := fruits[1:3:3]
	some = append(some, "Banana")
	fmt.Println("some: ", some)
	fmt.Println("use new slice: ", fruits)
	some = fruits[1:3]
	some = append(some, "Banana")
	fmt.Println("some: ", some)
	fmt.Println("use same slice", fruits)
}
