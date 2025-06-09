package util

/*
#cgo CFLAGS: -mavx2
#include <stdint.h>
#include <immintrin.h>

void add_sub_simd(int64_t* vec1, int64_t* vec2){

	__m256i reg1 = _mm256_loadu_si256((__m256i *)vec1);
	__m256i reg2 = _mm256_loadu_si256((__m256i *)vec2);
	__m256i reg_result = _mm256_add_epi64(reg1, reg2);
	_mm256_storeu_si256((__m256i*)vec1, reg_result);

	reg_result = _mm256_sub_epi64(reg1, reg2);
	_mm256_storeu_si256((__m256i*)vec2, reg_result);


	reg1 = _mm256_loadu_si256((__m256i *)(vec1+4));
	reg2 = _mm256_loadu_si256((__m256i *)(vec2+4));
	reg_result = _mm256_add_epi64(reg1, reg2);
	_mm256_storeu_si256((__m256i*)(vec1+4), reg_result);

	reg_result = _mm256_sub_epi64(reg1, reg2);
	_mm256_storeu_si256((__m256i*)(vec2+4), reg_result);

}

void left_shift_simd(int64_t* vec, int shift){

	__m256i reg = _mm256_loadu_si256((__m256i *)vec);
	reg = _mm256_slli_epi64(reg,shift);
	_mm256_storeu_si256((__m256i*)vec, reg);

 	reg = _mm256_loadu_si256((__m256i *)(vec+4));
 	reg = _mm256_slli_epi64(reg,shift);
 	_mm256_storeu_si256((__m256i*)(vec+4), reg);

}

void q_reduce_64(int64_t* val){
	 *val = (*val & 255) - (*val >> 8);
}

void mod_257_64(int64_t* val) {
	q_reduce_64(val);
	q_reduce_64(val);
  	*val = *val ^ (((*val == -1)*-1) & (-257));
}


*/
import "C"
import "unsafe"

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
func SIMD_AddSub(vec1 *[8]int64, vec2 *[8]int64) {

	C.add_sub_simd((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])))
}

//Mults vec2 and vec3 then adds the prodcut to vec1, sum is returned in vec1
func SIMD_Add_Mult(vec1 *[8]int64, vec2 *[8]int64, vec3 *[8]int64) {

	C.add_mult_simd_64((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])), (*C.int64_t)(unsafe.Pointer(&vec3[0])))
}

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
func SIMD_Shift(vec *[8]int64, shift int) {

	C.left_shift_simd((*C.int64_t)(unsafe.Pointer(&vec[0])), (C.int)(shift))
}

//Multiples two vectors element-wise using SIMD instrucitons, returns product
func SIMD_Mult(vec1 *[8]int64, vec2 *[8]int64) [8]int64 {

	var product [8]int64

	C.mult_simd_64((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])), (*C.int64_t)(unsafe.Pointer(&product[0])))

	return product
}

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
// Running q_reduce twice miimizes the the domain to -1 to 256, the following code removes the -1.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
func Mod_257(val *int64) {
	C.mod_257_64((*C.int64_t)(unsafe.Pointer(val)))
}

// Efficient mod of 257 (or any q=2^x+1 if modified).
// The tradeoff is that the reduced range is not 0 to 256, but rather -127 tp 387.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
func Q_reduce(val *int64) {
	C.q_reduce_64((*C.int64_t)(unsafe.Pointer(val)))
}

func SIMD_Q_Reduce(vec *[8]int64) {

	C.simd_q_reduce_64((*C.int64_t)(unsafe.Pointer(&vec[0])))
}

func SIMD_Mod_257(vec *[8]int64) {

	C.simd_mod_257_64((*C.int64_t)(unsafe.Pointer(&vec[0])))
}
