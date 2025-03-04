package key

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

const d uint64 = 1
const n uint64 = 128
const q uint64 = 65537
const m uint64 = 32

//Seed rng with MiNNT
//returns 2d string slice length M x (n*d)
func tempKeyGen() [][]string {
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

func GenWriteKey(filepath string) {

	key := tempKeyGen()

	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll([][]string(key))

}

//Assumes key is eactly M x (n*d)
func ReadKey(filepath string) [m][d * n]uint64 {

	var key [m][d * n]uint64

	fi, _ := os.Open(filepath)
	r := csv.NewReader(fi)
	keystring, _ := r.ReadAll()

	for i := range keystring {
		for j := range keystring[i] {
			num, _ := strconv.Atoi(keystring[i][j])
			key[i][j] = uint64(num)
		}
	}
	return key
}

func OmegaGen() [n]uint64 {

	var omegas [n]uint64
	omegas[0] = 1
	omegas[1] = 44120
	for i := uint64(2); i < n; i++ {
		omegas[i] = omegas[i-1] * omegas[1] % q
	}

	return omegas
}

const OMEGA uint64 = 59963

// Uses omega=59963
func TableGen() [256][2][8]uint64 {

	//Table is indexed first by 8 bit inout x, then i1 mod 2, then i0
	var table [256][2][8]uint64
	var inter_table [2][8][8]int
	bfb_table := bitsFromByteTable()

	for i1 := 0; i1 < 2; i1++ {
		for i0 := 0; i0 < 8; i0++ {
			for k1 := 0; k1 < 8; k1++ {
				pow := uint64Powq(OMEGA, (8*uint64(k1)*(2*uint64(i0)+1))%(2*n))
				neg := intPow(-1, (i1*k1)%2)
				inter_table[i1][i0][k1] = int(pow) * neg
			}
		}
	}

	// fmt.Println(inter_table)

	for x := uint64(0); x < 256; x++ {
		for i1 := 0; i1 < 2; i1++ {
			for i0 := 0; i0 < 8; i0++ {
				for k1 := 0; k1 < 8; k1++ {
					table[x][i1][i0] = table[x][i1][i0] + uint64(Mod(inter_table[i1][i0][k1]*bfb_table[x][k1], int(q)))
				}
			}
		}

	}

	return table
}

func bitsFromByteTable() [256][8]int {

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

func intPow(b int, x int) int {

	var result int = 1

	for i := 0; i < x; i++ {
		result = result * b
	}
	return result
}

func uint64Powq(b uint64, x uint64) uint64 {

	var result uint64 = 1

	for i := uint64(0); i < x; i++ {
		result = (result * b) % q
	}
	return result
}

//Golang mod of a negative returns a negative, this returns the positive
func Mod(a int, b int) int {
	return ((a % b) + b) % b
}
