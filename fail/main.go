package main

import (
	"github.com/RileyVaughn/MiNTT/numbers"
	"github.com/RileyVaughn/MiNTT/polynom"
)

func main() {

	// nums := util.GenRandInt2d(50, 20)
	// util.WriteIntCSV("./polynom/polynomials.csv", nums)

	// nums := util.ReadIntCSV("test.csv")
	// fmt.Println(nums)

	// var polyTest []poly.Polynom = []poly.Polynom{
	// 	[]num.Intq{Intq{Z: 0, Q: 9}, Intq{Z: 5, Q: 9}, Intq{Z: 7, Q: 9}, Intq{Z: 6, Q: 9}},
	// 	[]Intq{Intq{Z: 9, Q: 9}, Intq{Z: 5, Q: 9}, Intq{Z: 7, Q: 9}, Intq{Z: 6, Q: 9}},
	// 	[]Intq{Intq{Z: 0, Q: 9}, Intq{Z: 0, Q: 9}, Intq{Z: 0, Q: 9}, Intq{Z: 0, Q: 9}},
	// }

	var poly polynom.Polynom = polynom.Polynom{numbers.Intq{Z: 0, Q: 9}, numbers.Intq{Z: 5, Q: 9}, numbers.Intq{Z: 7, Q: 9}, numbers.Intq{Z: 6, Q: 9}}

}
