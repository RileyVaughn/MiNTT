package util

/*
#cgo CFLAGS: -mavx2
#include <stdint.h>
#include <immintrin.h>

void add_sub_simd_16(int16_t* vec1, int16_t* vec2){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec1);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec2);
	__m128i reg_result = _mm_add_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec1, reg_result);

	reg_result = _mm_sub_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec2, reg_result);

}

void left_shift_simd_16(int16_t* vec, int shift){

	__m128i reg = _mm_loadu_si128((__m128i *)vec);
	reg = _mm_slli_epi16(reg,shift);
	_mm_storeu_si128((__m128i*)vec, reg);
}

void mult_simd_16(int16_t* vec1, int16_t* vec2, int16_t * product){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec1);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec2);
	reg1 = _mm_mullo_epi16(reg1,reg2);
	_mm_storeu_si128((__m128i*)product, reg1);
}

void add_mult_simd_16(int16_t* vec1, int16_t* vec2, int16_t* vec3){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec2);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec3);

	reg2 = _mm_mullo_epi16(reg1,reg2);
	reg1 = _mm_loadu_si128((__m128i *)vec1);
	reg1 = _mm_add_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec1, reg1);

}


*/
import "C"
import "unsafe"

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
func SIMD_AddSub(vec1 *[8]int16, vec2 *[8]int16) {

	C.add_sub_simd_16((*C.int16_t)(unsafe.Pointer(&vec1[0])), (*C.int16_t)(unsafe.Pointer(&vec2[0])))
}

//Mults vec2 and vec3 then adds the prodcut to vec1, sum is returned in vec1
func SIMD_Add_Mult(vec1 *[8]int16, vec2 *[8]int16, vec3 *[8]int16) {

	C.add_mult_simd_16((*C.int16_t)(unsafe.Pointer(&vec1[0])), (*C.int16_t)(unsafe.Pointer(&vec2[0])), (*C.int16_t)(unsafe.Pointer(&vec3[0])))
}

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
func SIMD_Shift(vec *[8]int16, shift int) {

	C.left_shift_simd_16((*C.int16_t)(unsafe.Pointer(&vec[0])), (C.int)(shift))
}

//Multiples two vectors element-wise using SIMD instrucitons, returns product
func SIMD_Mult(vec1 *[8]int16, vec2 *[8]int16) [8]int16 {

	var product [8]int16

	C.mult_simd_16((*C.int16_t)(unsafe.Pointer(&vec1[0])), (*C.int16_t)(unsafe.Pointer(&vec2[0])), (*C.int16_t)(unsafe.Pointer(&product[0])))

	return product
}
