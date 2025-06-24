package main

import (
	"fmt"
	"math/rand"
	"time"

	m128_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT128"
	m64_norm_int64 "github.com/RileyVaughn/MiNTT/hash/int64/normal/MiNTT64"

	m128_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT128"
	m64_simd_int64 "github.com/RileyVaughn/MiNTT/hash/int64/simd/MiNTT64"

	m128_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT128"
	m64_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT64"
	m8_simd_int16 "github.com/RileyVaughn/MiNTT/hash/int16/simd/MiNTT8"

	m128_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT128"
	m64_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT64"
	m8_norm_int16 "github.com/RileyVaughn/MiNTT/hash/int16/normal/MiNTT8"
)

const IN_SIZE = 1728
const OUT_SIZE = IN_SIZE/2
const TEST_SIZE = 10000


func main() {

	Setup()

	TestAll()
	//MeanRuntimeAll()

	
}

// // Sets up all hash variants
func Setup() {

	m8_norm_int16.SetupM8()
	m64_norm_int16.SetupM64()
	m128_norm_int16.SetupM128()

	m8_simd_int16.SetupM8()
	m64_simd_int16.SetupM64()
	m128_simd_int16.SetupM128()

	m128_norm_int64.SetupM128()
	m64_norm_int64.SetupM64()

	m128_simd_int64.SetupM128()
	m64_simd_int64.SetupM64()

}

func GenInput()[IN_SIZE]byte{
	var input [IN_SIZE]byte

	for i := 0; i < 1728; i++ {
		input[i] = byte(rand.Intn(256))
	}
	return input
}

//outputs are in nano-seconds
func CheckRuntime(f func([IN_SIZE]byte)[OUT_SIZE]byte) int64{

	start := time.Now()

	f(GenInput())
	
	end := time.Now()
	return end.Sub(start).Nanoseconds()

}

func MeanRuntime(f func([IN_SIZE]byte)[OUT_SIZE]byte) int64{

	var mean int64 = 0
	for i := 0; i < TEST_SIZE; i++ {
		mean = mean + CheckRuntime(f)

	}
	return mean/TEST_SIZE
}

func MeanRuntimeAll(){

	mean := MeanRuntime(m64_norm_int64.MiNTT64)
	fmt.Println(mean, "m64_norm_int64")
	mean = MeanRuntime(m64_simd_int64.MiNTT64)
	fmt.Println(mean, "m64_simd_int64")
	mean = MeanRuntime(m64_norm_int16.MiNTT64)
	fmt.Println(mean, "m64_norm_int16")
	mean = MeanRuntime(m64_simd_int16.MiNTT64)
	fmt.Println(mean, "m64_simd_int16")


}


func TestOut(f func([IN_SIZE]byte)[OUT_SIZE]byte){

	fmt.Println(f(GenInput()))

}

func TestAll(){

	TestOut(m64_norm_int64.MiNTT64)
	TestOut(m64_simd_int64.MiNTT64)
	TestOut(m64_norm_int16.MiNTT64)
	TestOut(m64_simd_int16.MiNTT64)
}