package main

import (
	"fmt"

	keygen "github.com/RileyVaughn/MiNTT/hash/key"
)

func main() {

	keygen.GenWriteKey("./key/key.csv")
	fmt.Println(keygen.ReadKey("./key/key.csv"))

}
