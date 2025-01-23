package hash

import (
	"math/rand"

	c "github.com/RileyVaughn/MiNTT/ineff/constant"
	p "github.com/RileyVaughn/MiNTT/ineff/polynom"
)

func TempKeyGen() [][]p.Polynom {

	key := make([][]p.Polynom, c.M)

	for i := 0; i < c.M; i++ {
		dim := make([]p.Polynom, c.D)
		for j := 0; j < c.D; j++ {
			poly := make(p.Polynom, c.N)
			for k := 0; k < c.N; k++ {
				poly[k] = rand.Intn(c.Q)
			}
			dim[j] = poly
		}
		key[i] = dim
	}

	return key
}
