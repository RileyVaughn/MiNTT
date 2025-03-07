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

//Integer a to the power of postive Integer b
func IntPow(b int, x int) int {

	var result int = 1

	for i := 0; i < x; i++ {
		result = result * b
	}
	return result
}

//Golang mod of a negative returns a negative, this returns the positive
func Mod(a int, b int) int {
	return ((a % b) + b) % b
}
