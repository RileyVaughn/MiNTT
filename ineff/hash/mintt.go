package hash

import (
	c "github.com/RileyVaughn/MiNTT/ineff/constant"
	p "github.com/RileyVaughn/MiNTT/ineff/polynom"
)

func MiNTT(input string) string {

	return mbitPolysToHexStr(stringToBitPoly(input))

	//poly_input = p.Polynom(input)

	//convert to binary
	//convert TempKeyGen()

}

//If string is less than 512(NxM) pad with 0's, if longer truncate
func stringToBitPoly(input string) [c.M]p.Polynom {

	var bytes [c.N * c.M / 8]byte

	//Converts string input into 512bit input
	for i := 0; i < c.N*c.M/8; i++ {
		if i < len(input) {
			bytes[i] = byte(input[i])
		} else {
			bytes[i] = 0
		}
	}

	return mtNBytesToBitPoly(bytes)
}

//512 byte input.
//assume all charcters are ascii (are bytes)
func mtNBytesToBitPoly(input [c.N * c.M / 8]byte) [c.M]p.Polynom {

	var polys [c.M]p.Polynom

	for i := 0; i < c.M; i++ {
		var nd8_bytes [c.N / 8]byte
		for j := 0; j < c.N/8; j++ {
			nd8_bytes[j] = input[(i*c.N/8)+j]
		}

		polys[i] = ndiv8BytesToBitPoly(nd8_bytes)
	}

	return polys
}

//Takes N/8 bytes and returns a polynomial with binary coefficients
func ndiv8BytesToBitPoly(input [c.N / 8]byte) p.Polynom {

	var poly p.Polynom

	for i := 0; i < c.N/8; i++ {
		bytes := byteToBits(input[i])

		for j := 0; j < 8; j++ {
			poly[(i*8)+j] = int(bytes[j])
		}
	}
	return poly
}

// Little endian
// buts must be stores as bytes
func byteToBits(val byte) [8]byte {

	var bits [8]byte

	for i := 0; i < 8; i++ {
		bits[i] = val % 2
		val = val >> 1
	}
	return bits
}

func mbitPolysToHexStr(polys [c.M]p.Polynom) string {

	var hex string

	for i := 0; i < c.M; i++ {
		var bytes [c.N / 8]byte = bitPolyToNdiv8Bytes(polys[i])
		for j := 0; j < c.N/8; j++ {
			hex = hex + string(bytes[j])
		}
	}

	return hex
}

// Turns a bit-poly into 32(N/8) bytes
func bitPolyToNdiv8Bytes(poly p.Polynom) [c.N / 8]byte {

	var bytes [c.N / 8]byte

	for i := 0; i < c.N/8; i++ {
		var sbyte byte = 0
		for j := 7; j >= 0; j-- {
			sbyte = (sbyte << 1) + byte(poly[i*8+j])
		}
		bytes[i] = sbyte
	}

	return bytes
}
