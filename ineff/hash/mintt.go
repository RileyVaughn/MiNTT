package hash

import (
	c "github.com/RileyVaughn/MiNTT/ineff/constant"
	p "github.com/RileyVaughn/MiNTT/ineff/polynom"
)

func MiNTT(input []int64) {

	//poly_input = p.Polynom(input)

	//convert to binary
	//convert TempKeyGen()

}

// Little endian
// buts must be stores as bytes
func byteToBits(val byte) [8]byte {

	var bits [8]bytes

	for i := 0; i < 8; i++ {
		bits[i] = val % 2
		val = val >> 1
	}
	return bits
}

//Takes N/8 bytes and returns a polynomial with binary coefficients
func Ndiv8BytesToBitPoly(input []byte) p.Polynom {

	var poly p.Polynom

	for i := 0; i < c.N/8; i++ {
		bytes := byteToBits(input[i])
		for j := 0; j < 8; j++ {
			poly[i+j] = int(bytes[j])
		}
	}
	return poly
}

//512 byte input. 
//assume all charcters are ascii (are bytes)
func MtimesNBytesToBitPoly(input [c.N*c.M]bytes) [c.M]p.Polynom {

	var polys [c.M]p.Polynom

	for i := 0; i < c.M; i++ {
		var polys [c.N/8]p.Polynom
		for j := 0; j < c.N/8; j++ {
			bytes[j]
		}


		polys[i] = Ndiv8BytesToPoly[]
	}

	return polys
}
