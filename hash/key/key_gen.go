package key

import (
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
)

const d int = 1
const n int = 256
const q int = 65537
const m int = 32

//Seed rng with MiNNT
//returns 2d string slice length M x (n*d)
func TempKeyGen() [][]string {
	seed, _ := strconv.Atoi("MiNNT")
	rand.Seed(int64(seed))

	key := make([][]string, m)
	for i := range key {
		keyrow := make([]string, n*d)
		for j := range keyrow {
			keyrow[j] = strconv.Itoa(rand.Intn(q))
		}
		key[i] = keyrow
	}

	return key
}

func GenWriteKey(filepath string) {

	key := TempKeyGen()

	fo, _ := os.Create(filepath)
	w := csv.NewWriter(fo)
	w.WriteAll([][]string(key))

}
