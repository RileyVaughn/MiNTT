#include <util_int64.h>
#include <stdlib.h>
#include <immintrin.h>
#include <iostream>



int64_t Util64::Q_reduce(int64_t val){
	return (val & 255) - (val >> 8);
}


int64_t Util64::QF4_reduce(int64_t val){
    return (val & 65535) - (val >> 16);
}


int64_t Util64::Mod_257(int64_t val) {
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);
	val = Q_reduce(val);

  	return val ^ (((val == -1)*-1) & (-257));
}

int64_t Util64::Mod_65537(int64_t val) {
	val = QF4_reduce(val);
	val = QF4_reduce(val);
	val = QF4_reduce(val);
	val = QF4_reduce(val);

  	return val ^ (((val == -1)*-1) & (-65537));
}


int64_t Util64::IntPow(int64_t b, int64_t x, int64_t q) {
    int64_t result = 1;
    for (size_t i = 0; i < x; i++){
        result = (result * b) % q;
    }
    return result;
}


int64_t Util64::Bit_Rev(int64_t i, int64_t bound){
    int64_t irev = 0;
    for (i = i | bound; i > 1; i = i >> 1){
        irev = (irev << 1) | (i & 1);
    }
    return irev;
}


void Util64::addSub(int64_t * a, int64_t * b) {
    int64_t temp = *b;
    *b = *a - *b;
    *a = *a + temp;
}


void Util64::BitsFromByte(int64_t byte, int64_t bits[8]){
    for (size_t i = 0; i < 8; i++)
    {
        bits[i] = byte %2;
        byte = byte >> 1;
    }
} 


void Util64::Gen8NCCMat(int64_t omega, int64_t q, int64_t mat[8][8]){

    int64_t ncc_mat[8][8];
    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; k < 8; k++){
            if ((k*(2*i+1))%(2*8) <= 8) {
                ncc_mat[i][k] = Util64::IntPow(omega,(k*(2*i+1))%8, q);
            } else {
                ncc_mat[i][k] = -1 * Util64::IntPow(omega, ((k*(2*i+1))%8), q);
            }
        }
    }
    
    for (size_t i = 0; i < 8; i++){
        for (size_t j = 0; j < 8; j++){
        mat[j][Util64::Bit_Rev(i, 8)] = ncc_mat[j][i];
        }
    }

}


void Util64::GenNTT8Table(int64_t omega, int64_t q, int64_t table[256][8]){

    int64_t ncc_mat[8][8];
    Gen8NCCMat(omega,q,ncc_mat);

    for (size_t i = 0; i < 256; i++){
        for (size_t j = 0; j < 8; j++){
            int64_t vec[8];
            Util64::BitsFromByte(i, vec);
            table[i][j] = 0;
            for (size_t k = 0; k < 8; k++){
                table[i][j] = Mod_257(table[i][j] + ncc_mat[j][k]*vec[k]);
            }
        }
 
    }

}


int64_t * Util64::GenMultTable(int64_t omega, int64_t n, int64_t q){

    int64_t * table;
    table = new int64_t[n];

    for (size_t i = 0; i < n/8; i++){
        for (size_t k = 0; k < 8; k++){
            table[k*8+i] = IntPow(omega,(Bit_Rev(k,8)*(2*i+1))%(2*n),q);
        }
    }
    return table;
}

int64_t * Util64::GenKey(int64_t m, int64_t n, int64_t d, int64_t q){

    int64_t * key;
    key = new int64_t[m*d*n];

    //just for the sake of reproducibility srand(1)
    // A better implementation would usea  a "trusted randomoness" such as digits of pi
    srand(1);

    

    for (size_t i = 0; i < m; i++){
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < n/8; k++){
                for (size_t l = 0; l < 8; l++){
                    key[i*(d*n)+j*n+k*8+l] = (int64_t)(rand() % q);
                }
            } 
        } 
    }

    return key;
}


void Util64::Norm_AddSub(int64_t * vec1, int64_t * vec2){
    for (size_t i = 0; i < 8; i++){
        addSub(vec1+i,vec2+i);
    }
}


void Util64::Norm_AddMult(int64_t * vec1, int64_t * vec2, int64_t * vec3){
    for (size_t i = 0; i < 8; i++){
        vec1[i] = vec1[i] + (vec2[i] * vec3[i]);
    }
}


void Util64::Norm_LShift(int64_t * vec, int64_t shift){
    for (size_t i = 0; i < 8; i++){
        vec[i] = vec[i] << shift;
    }
}


void Util64::Norm_Mult(int64_t * vec1, int64_t * vec2, int64_t * prod){
    for (size_t i = 0; i < 8; i++){
        prod[i] = vec1[i] * vec2[i];
    }
}


void Util64::Norm_Mod257(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Mod_257(vec[i]);
    }
}


void Util64::Norm_Q_reduce(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Q_reduce(vec[i]);
    }
}

void Util64::Norm_Mod65537(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Mod_65537(vec[i]);
    }
}


void Util64::Norm_QF4_reduce(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = QF4_reduce(vec[i]);
    }
}
    

void Util64::SIMD_AddSub(int64_t* vec1, int64_t* vec2){

	__m512i reg1 = _mm512_loadu_si512((__m512i *)vec1);
	__m512i reg2 = _mm512_loadu_si512((__m512i *)vec2);

	__m512i reg_result = _mm512_add_epi64(reg1, reg2);
	_mm512_storeu_si512((__m512i *)vec1, reg_result);

	reg_result = _mm512_sub_epi64(reg1, reg2);
	_mm512_storeu_si512((__m512i *)vec2, reg_result);

}


void Util64::SIMD_LShift(int64_t* vec, int64_t shift){

	__m512i reg = _mm512_loadu_si512((__m512i *)vec);
	reg = _mm512_slli_epi64(reg,shift);
	_mm512_storeu_si512((__m512i *)vec, reg);

}


void Util64::SIMD_Mult(int64_t* vec1, int64_t* vec2, int64_t* product) {
    __m512i reg1 = _mm512_loadu_si512((__m512i *)vec1);
    __m512i reg2 = _mm512_loadu_si512((__m512i *)vec2);

    __m512i result = _mm512_mullo_epi64(reg1, reg2);

    _mm512_storeu_si512((__m512i *)product, result);
}


void Util64::SIMD_AddMult(int64_t* vec1, int64_t* vec2, int64_t* vec3) {
    __m512i reg1 = _mm512_loadu_si512((__m512i *)vec2);
    __m512i reg2 = _mm512_loadu_si512((__m512i *)vec3);

    reg2 = _mm512_mullo_epi64(reg1, reg2);

    reg1 = _mm512_loadu_si512((__m512i *)vec1);
    reg1 = _mm512_add_epi64(reg1, reg2);

    _mm512_storeu_si512((__m512i *)vec1, reg1);
}


void Util64::SIMD_Q_reduce(int64_t* vec1) {
    const __m512i TFF = _mm512_set1_epi64(255);

    __m512i l_reg = _mm512_loadu_si512((__m512i *)vec1);
    __m512i r_reg = _mm512_srai_epi64(l_reg, 8);
    l_reg = _mm512_and_si512(l_reg, TFF);
    l_reg = _mm512_sub_epi64(l_reg, r_reg);

    _mm512_storeu_si512((__m512i *)vec1, l_reg);
}


void Util64::SIMD_Mod257(int64_t* vec1){
	const __m512i NEG_ONE = _mm512_set1_epi64(-1);
	const __m512i NEG_TFS = _mm512_set1_epi64(-257);

	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);

	__m512i l_reg = _mm512_loadu_si512((__m512i *)vec1);
	__mmask32 mask = _mm512_cmpeq_epi64_mask(l_reg,NEG_ONE);
	__m512i r_reg = _mm512_maskz_mov_epi64(mask, NEG_ONE);
	r_reg = _mm512_and_si512(r_reg, NEG_TFS);
	l_reg = _mm512_xor_si512(l_reg,r_reg);
	_mm512_storeu_si512((__m512i*)vec1, l_reg);
}