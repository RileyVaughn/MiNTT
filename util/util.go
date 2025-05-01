package util

func BitsFromByteTable() [256][8]int {

	var table [256][8]int

	for i := int(0); i < 256; i++ {
		table[i][0] = i % 2
		table[i][1] = (i >> 1) % 2
		table[i][2] = (i >> 2) % 2
		table[i][3] = (i >> 3) % 2
		table[i][4] = (i >> 4) % 2
		table[i][5] = (i >> 5) % 2
		table[i][6] = (i >> 6) % 2
		table[i][7] = (i >> 7) % 2
	}

	return table
}

//Integer b to the power of postive Integer x
func IntPow(b int, x int, q int) int {

	var result int = 1

	for i := 0; i < x; i++ {
		result = result * b % q
	}
	return result
}

//Golang mod of a negative returns a negative, this returns the positive
func Mod(a int, b int) int {
	return ((a % b) + b) % b
}

func AddSub(a *int, b *int) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

func Bit_Rev(i int, bound int) int {
	var irev int = 0
	for i = i | bound; i > 1; i = i >> 1 {
		irev = (irev << 1) | (i & 1)
	}
	return irev
}
