package util

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

// Creates a table which is indexed by a byte and returns a vector of 8 "bits":
// BitsFromByteTable()[5]->[0][0][0][0][0][1][0][1]
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

// Integer b to the power of postive Integer x, mod q
func IntPow(b int, x int, q int) int {

	var result int = 1

	for i := 0; i < x; i++ {
		result = result * b % q
	}
	return result
}

// Golang mod of a negative returns a negative, this returns the positive
func Mod(a int, b int) int {
	return ((a % b) + b) % b
}

// Adds and subtracts inputs in place:
// a+b -> a, a-b -> b
func AddSub(a *int, b *int) {
	temp := *b
	*b = *a - *b
	*a = *a + temp
}

// Reverse the bits of an int up to a bound
func Bit_Rev(i int, bound int) int {
	var irev int = 0
	for i = i | bound; i > 1; i = i >> 1 {
		irev = (irev << 1) | (i & 1)
	}
	return irev
}

// A quick and dirty keygen using GO's math randomness library
// returns 2d string slice length M x (n*d)
func tempKeyGen(m int, n int, d int, q int) [][]string {
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
func GenWriteKey(m int, n int, d int, q int, filepath string) {

	key := tempKeyGen(m, n, d, q)

	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll([][]string(key))

}

// Fake for now
func SIMD_AddSub(vec1 *[8]int, vec2 *[8]int) {
	for i := 0; i < 8; i++ {
		AddSub(&vec1[i], &vec2[i])
	}
}

// Fake for now
func SIMD_Shift(vec *[8]int, shift int) {

	for i := 0; i < 8; i++ {
		vec[i] = vec[i] << shift
	}

}

// Fake for now
func SIMD_Mult(vec1 [8]int, vec2 [8]int) [8]int {

	var product [8]int
	for i := 0; i < 8; i++ {
		product[i] = vec1[i] * vec2[i]

	}
	return product
}
