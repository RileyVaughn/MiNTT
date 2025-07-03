package hash

import (
	"math/rand"
	"strconv"

	c "github.com/RileyVaughn/MiNTT/ineff/constant"
	p "github.com/RileyVaughn/MiNTT/ineff/polynom"
)

//Seed rng with MiNNT
func TempKeyGen() [c.M][c.D]p.Polynom {
	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))
	var key [c.M][c.D]p.Polynom

	for i := 0; i < c.M; i++ {
		var dim [c.D]p.Polynom
		for j := 0; j < c.D; j++ {
			var poly p.Polynom
			for k := 0; k < c.N; k++ {
				poly[k] = rand.Intn(c.Q)
			}
			dim[j] = poly
		}
		key[i] = dim
	}

	return key
}
