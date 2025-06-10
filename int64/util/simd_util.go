package util

/*
#cgo CFLAGS: -mavx2 -mavx512f --mavx512bw -mavx512dq -mavx512vl
#include <stdint.h>
#include <immintrin.h>

void add_sub_simd_64(int64_t* vec1, int64_t* vec2){

	__m512i reg1 = _mm512_loadu_si512((__m512i *)vec1);
	__m512i reg2 = _mm512_loadu_si512((__m512i *)vec2);

	__m512i reg_result = _mm512_add_epi64(reg1, reg2);
	_mm512_storeu_si512((__m512i *)vec1, reg_result);

	reg_result = _mm512_sub_epi64(reg1, reg2);
	_mm512_storeu_si512((__m512i *)vec2, reg_result);


}

void left_shift_simd_64(int64_t* vec, int shift){

	__m512i reg = _mm512_loadu_si512((__m512i *)vec);
	reg = _mm512_slli_epi64(reg,shift);
	_mm512_storeu_si512((__m512i *)vec, reg);


}

void mult_simd_64(int64_t* vec1, int64_t* vec2, int64_t* product) {
    __m512i reg1 = _mm512_loadu_si512((__m512i *)vec1);
    __m512i reg2 = _mm512_loadu_si512((__m512i *)vec2);

    __m512i result = _mm512_mullo_epi64(reg1, reg2);

    _mm512_storeu_si512((__m512i *)product, result);
}

void add_mult_simd_64(int64_t* vec1, int64_t* vec2, int64_t* vec3) {
    __m512i reg1 = _mm512_loadu_si512((__m512i *)vec2);
    __m512i reg2 = _mm512_loadu_si512((__m512i *)vec3);

    reg2 = _mm512_mullo_epi64(reg1, reg2);

    reg1 = _mm512_loadu_si512((__m512i *)vec1);
    reg1 = _mm512_add_epi64(reg1, reg2);

    _mm512_storeu_si512((__m512i *)vec1, reg1);
}



void q_reduce_64(int64_t* val){
	 *val = (*val & 255) - (*val >> 8);
}

void mod_257_64(int64_t* val) {
	q_reduce_64(val);
	q_reduce_64(val);
  	*val = *val ^ (((*val == -1)*-1) & (-257));
}

void simd_q_reduce_64(int64_t* vec1) {
    const __m512i TFF = _mm512_set1_epi64(255);

    __m512i l_reg = _mm512_loadu_si512((__m512i *)vec1);
    __m512i r_reg = _mm512_srai_epi64(l_reg, 8);
    l_reg = _mm512_and_si512(l_reg, TFF);
    l_reg = _mm512_sub_epi64(l_reg, r_reg);

    _mm512_storeu_si512((__m512i *)vec1, l_reg);
}

#include <immintrin.h>

void simd_mod_257_64(int64_t* vec1){
	const __m512i NEG_ONE = _mm512_set1_epi64(-1);
	const __m512i NEG_TFS = _mm512_set1_epi64(-257);

	simd_q_reduce_64(vec1);
	simd_q_reduce_64(vec1);

	__m512i l_reg = _mm512_loadu_si512((__m512i *)vec1);
	__mmask32 mask = _mm512_cmpeq_epi64_mask(l_reg,NEG_ONE);
	__m512i r_reg = _mm512_maskz_mov_epi64(mask, NEG_ONE);
	r_reg = _mm512_and_si512(r_reg, NEG_TFS);
	l_reg = _mm512_xor_si512(l_reg,r_reg);
	_mm512_storeu_si512((__m512i*)vec1, l_reg);
}


*/
import "C"
import "unsafe"

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
func SIMD_AddSub(vec1 *[8]int64, vec2 *[8]int64) {

	C.add_sub_simd_64((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])))
}

//Mults vec2 and vec3 then adds the prodcut to vec1, sum is returned in vec1
func SIMD_Add_Mult(vec1 *[8]int64, vec2 *[8]int64, vec3 *[8]int64) {

	C.add_mult_simd_64((*C.int64_t)(unsafe.Pointer(&vec1[0])), (*C.int64_t)(unsafe.Pointer(&vec2[0])), (*C.int64_t)(unsafe.Pointer(&vec3[0])))
}

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
func SIMD_Shift(vec *[8]int64, shift int) {

	C.left_shift_simd_64((*C.int64_t)(unsafe.Pointer(&vec[0])), (C.int)(shift))
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
