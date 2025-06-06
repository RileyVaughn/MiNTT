package util

import (
	"encoding/csv"
	"math"
	"math/rand"
	"os"
	"strconv"
)

// Creates a table which is indexed by a byte and returns a vector of 8 "bits":
// BitsFromByteTable()[5]->[0][0][0][0][0][1][0][1]
func BitsFromByteTable() [256][8]int16 {

	var table [256][8]int16

	for i := int16(0); i < 256; i++ {
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

// Integer b to the power of postive Integer x, mod q
func IntPow(b int16, x int16, q int16) int16 {

	var result int16 = 1

	for i := int16(0); i < x; i++ {
		result = result * b % q
	}
	return result
}

// Slow version
// Golang mod of a negative returns a negative, this returns the positive
func Mod(a int16, b int16) int16 {
	return ((a % b) + b) % b
}

// Reverse the bits of an int up to a bound
func Bit_Rev(i int16, bound int16) int16 {
	var irev int16 = 0
	for i = i | bound; i > 1; i = i >> 1 {
		irev = (irev << 1) | (i & 1)
	}
	return irev
}

// A quick and dirty keygen using GO's math randomness library
// returns 2d string slice length M x (n*d)
func tempKeyGen(m int16, n int16, d int16, q int16) [][]string {
	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	key := make([][]string, m)
	for i := range key {
		keyrow := make([]string, n*d)
		for j := range keyrow {
			keyrow[j] = strconv.Itoa(rand.Intn(int(q)))
		}
		key[i] = keyrow
	}

	return key
}

// Writes a key gnerated by tempKeyGen() to a file
func GenWriteKey(m int16, n int16, d int16, q int16, filepath string) {

	key := tempKeyGen(m, n, d, q)

	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll([][]string(key))

}

// A fake version of simd add/sub that does the addsub function on each element one at a time.
func Fake_SIMD_AddSub(vec1 *[8]int16, vec2 *[8]int16) {
	for i := 0; i < 8; i++ {
		addSub(&vec1[i], &vec2[i])
	}
}

// A fake version of Add mult. The sum is returned in position vec1
func Fake_SIMD_Add_Mult(vec1 *[8]int16, vec2 *[8]int16, vec3 *[8]int16) {

	for i := 0; i < 8; i++ {
		vec1[i] = vec1[i] + (vec2[i] * vec3[i])
	}

}

// Adds and subtracts inputs in place:
// a+b -> a, a-b -> b
func addSub(a *int16, b *int16) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

// A fake version of a SIMD shift that just shifts each element one at a time
func Fake_SIMD_Shift(vec *[8]int16, shift int16) {

	for i := 0; i < 8; i++ {
		vec[i] = vec[i] << shift
	}

}

// A fake version of SIMD mult that jsut multiplies each element one at a time.
func Fake_SIMD_Mult(vec1 [8]int16, vec2 [8]int16) [8]int16 {

	var product [8]int16
	for i := 0; i < 8; i++ {
		product[i] = vec1[i] * vec2[i]

	}
	return product
}

// A fake version of SIMD mod. Computes the modulo 257 one by one.
func Fake_SIMD_Mod(vec *[8]int16) {
	for i := 0; i < 8; i++ {
		Mod_257(&vec[i])
	}
}

func Fake_SIMD_Q_reduce(vec *[8]int16) {
	for i := 0; i < 8; i++ {
		Q_reduce(&vec[i])
	}
}

// Centers val such that |val| < q/2
func Center(val int16, q int16) int16 {

	val = Mod(val, q)
	floor := int16(math.Floor(float64(q) / 2))
	if val > floor {
		val = val - q
	}

	return val

}
