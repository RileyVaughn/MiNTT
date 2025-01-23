package main

import (
	"fmt"

	h "github.com/RileyVaughn/MiNTT/ineff/hash"
)

func main() {

	// nums := util.GenRandInt2d(50, 256, 10000)
	// util.WriteIntCSV("./polynom/polynomials.csv", nums)

	// nums := util.ReadIntCSV("test.csv")
	// fmt.Println(nums)

	fmt.Println(h.MiNTT("A paragraph is a group of sentences that present a single idea or topic. Paragraphs are a way to organize written work and help readers understand the author's ideas."))

}
