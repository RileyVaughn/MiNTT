package hash

func MiNTT(input []int64) {

	//poly_input = p.Polynom(input)

	//convert to binary
	//convert TempKeyGen()

}

// Little endian
func IntToBits(val int) []int {

	bits := make([]int, 64)

	for i := 0; i < 64; i++ {
		bits[i] = val % 2
		val = val >> 1
	}

	return bits
}
