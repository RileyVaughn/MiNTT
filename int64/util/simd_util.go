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

*/
import "C"
import "unsafe"

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
func SIMD_AddSub(vec1 *[8]int64, vec2 *[8]int64) {

	C.add_sub_simd((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])))
}

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
func SIMD_Shift(vec *[8]int64, shift int) {

	C.left_shift_simd((*C.int64_t)(unsafe.Pointer(&vec[0])), (C.int)(shift))
}
