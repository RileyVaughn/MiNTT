package main

import (
	"fmt"

	num "github.com/RileyVaughn/MiNTT/numbers"
)

func main() {

	// p1 := poly.Polynomial{1, 2, 3}
	// p2 := poly.Polynomial{4, 5, 0, 6}

	// p1 = p1.PolyMult(p2)
	// fmt.Println(p1)

	z1 := num.Intq{Z: 6, Q: 9}
	z2 := num.Intq{Z: 4, Q: 9}

	fmt.Print(z2.Add(z1))
	fmt.Print(z2.Sub(z1))
	fmt.Print(z2.Mult(z1))

}
