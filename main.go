package main

import (
	"fmt"

	poly "github.com/RileyVaughn/MiNTT/polynomial"
)

func main() {

	p1 := poly.Polynomial{1, 2, 3}
	p2 := poly.Polynomial{4, 5, 0, 6}

	p1 = p1.PolyMult(p2)
	fmt.Println(p1)

}
