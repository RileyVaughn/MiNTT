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



void q_reduce_16(int16_t* val){
	 *val = (*val & 255) - (*val >> 8);
}

void mod_257_16(int16_t* val) {
	q_reduce_16(val);
	q_reduce_16(val);
  	*val = *val ^ (((*val == -1)*-1) & (-257));
}


void simd_q_reduce_16(int16_t* vec1){
	const __m128i TFF = _mm_set1_epi16(255);

	__m128i l_reg = _mm_loadu_si128((__m128i *)vec1);
	__m128i r_reg = _mm_srai_epi16(l_reg,8);
	l_reg = _mm_and_si128(l_reg,TFF);
	l_reg = _mm_sub_epi16(l_reg, r_reg);

	_mm_storeu_si128((__m128i*)vec1, l_reg);
}

void simd_mod_257_16(int16_t* vec1){
	const __m128i NEG_ONE = _mm_set1_epi16(-1);
	const __m128i NEG_TFS = _mm_set1_epi16(-257);

	simd_q_reduce_16(vec1);
	simd_q_reduce_16(vec1);

	__m128i l_reg = _mm_loadu_si128((__m128i *)vec1);
	__m128i r_reg = _mm_cmpeq_epi16(l_reg,NEG_ONE);
	r_reg = _mm_and_si128(r_reg, NEG_TFS);
	l_reg = _mm_xor_si128(l_reg,r_reg);
	_mm_storeu_si128((__m128i*)vec1, l_reg);
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

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
// Running q_reduce twice miimizes the the domain to -1 to 256, the following code removes the -1.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, although slightly modified
func Mod_257(val *int16) {
	C.mod_257_16((*C.int16_t)(unsafe.Pointer(val)))
}

// Efficient mod of 257 (or any q=2^x+1 if modified).
// The tradeoff is that the reduced range is not 0 to 256, but rather -127 tp 387.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>
func Q_reduce(val *int16) {
	C.q_reduce_16((*C.int16_t)(unsafe.Pointer(val)))
}

func SIMD_Q_Reduce(vec *[8]int16) {

	C.simd_q_reduce_16((*C.int16_t)(unsafe.Pointer(&vec[0])))
}

func SIMD_Mod_257(vec *[8]int16) {

	C.simd_mod_257_16((*C.int16_t)(unsafe.Pointer(&vec[0])))
}
