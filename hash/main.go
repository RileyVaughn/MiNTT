package main

import (
	"fmt"

	keygen "github.com/RileyVaughn/MiNTT/hash/key"
)

func main() {

	// keygen.GenWriteKey("./key/key.csv")
	// fmt.Println(keygen.ReadKey("./key/key.csv"))
	// var input [32]byte = [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	// fmt.Println(fft(input))
	fmt.Println(keygen.TableGen())
	// keygen.TableGen()
}
