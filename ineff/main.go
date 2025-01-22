package main

import util "github.com/RileyVaughn/MiNTT/ineff/util"

func main() {

	nums := util.GenRandInt2d(50, 256, 10000)
	util.WriteIntCSV("./polynom/polynomials.csv", nums)

	// nums := util.ReadIntCSV("test.csv")
	// fmt.Println(nums)

}
