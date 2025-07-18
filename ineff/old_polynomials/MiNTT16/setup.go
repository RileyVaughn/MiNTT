package MiNTT16

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

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
func ReadKey(filepath string) [m][d * n]int {

	var key [m][d * n]int

	fi, _ := os.Open(filepath)
	r := csv.NewReader(fi)
	keystring, _ := r.ReadAll()

	for i := range keystring {
		for j := range keystring[i] {
			num, _ := strconv.Atoi(keystring[i][j])
			key[i][j] = int(num)
		}
	}
	return key
}
