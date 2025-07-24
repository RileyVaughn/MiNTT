#include <util_int16.h>
#include <stdlib.h>
#include <immintrin.h>
#include <iostream>



int16_t Util16::Q_reduce(int16_t val){
	return (val & 255) - (val >> 8);
}


int16_t Util16::Mod_257(int16_t val) {
	val = Q_reduce(val);
	val = Q_reduce(val);

  	return val ^ (((val == -1)*-1) & (-257));
}


int16_t Util16::IntPow(int16_t b, int16_t x, int16_t q) {
    int16_t result = 1;
    for (size_t i = 0; i < x; i++){
        result = (result * b) % q;
    }
    return result;
}


int16_t Util16::Bit_Rev(int16_t i, int16_t bound){
    int16_t irev = 0;
    for (i = i | bound; i > 1; i = i >> 1){
        irev = (irev << 1) | (i & 1);
    }
    return irev;
}


void Util16::addSub(int16_t * a, int16_t * b) {
    int16_t temp = *b;
    *b = *a - *b;
    *a = *a + temp;
}


void Util16::BitsFromByte(int16_t byte, int16_t bits[8]){
    for (size_t i = 0; i < 8; i++)
    {
        bits[i] = byte %2;
        byte = byte >> 1;
    }
} 


void Util16::Gen8NCCMat(int16_t omega, int16_t q, int16_t mat[8][8]){

    int16_t ncc_mat[8][8];
    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; k < 8; k++){
            if ((k*(2*i+1))%(2*8) <= 8) {
                ncc_mat[i][k] = Util16::IntPow(omega,(k*(2*i+1))%8, q);
            } else {
                ncc_mat[i][k] = -1 * Util16::IntPow(omega, ((k*(2*i+1))%8), q);
            }
        }
    }
    
    for (size_t i = 0; i < 8; i++){
        for (size_t j = 0; j < 8; j++){
        mat[j][Util16::Bit_Rev(i, 8)] = ncc_mat[j][i];
        }
    }

}


void Util16::GenNTT8Table(int16_t omega, int16_t q, int16_t table[256][8]){

    int16_t ncc_mat[8][8];
    Gen8NCCMat(omega,q,ncc_mat);

    for (size_t i = 0; i < 256; i++){
        for (size_t j = 0; j < 8; j++){
            int16_t vec[8];
            Util16::BitsFromByte(i, vec);
            table[i][j] = 0;
            for (size_t k = 0; k < 8; k++){
                table[i][j] = Center_257(table[i][j] + ncc_mat[j][k]*vec[k]);
            }
        }
 
    }

}


int16_t * Util16::GenMultTable(int16_t omega, int16_t n, int16_t q){

    int16_t * table;
    table = new int16_t[n];

    for (size_t i = 0; i < n/8; i++){
        for (size_t k = 0; k < 8; k++){
            table[k*8+i] = Center_257(IntPow(omega,(Bit_Rev(k,8)*(2*i+1))%(2*n),q));
        }
    }
    return table;
}

int16_t * Util16::GenKey(int16_t m, int16_t n, int16_t d, int16_t q){

    int16_t * key;
    key = new int16_t[m*d*n];

    //just for the sake of reproducibility srand(1)
    // A better implementation would usea  a "trusted randomoness" such as digits of pi
    srand(1);

    

    for (size_t i = 0; i < m; i++){
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < n/8; k++){
                for (size_t l = 0; l < 8; l++){
                    key[i*(d*n)+j*n+k*8+l] = Center_257((int16_t)(rand() % q));
                }
            } 
        } 
    }

    return key;
}

// Centers val such that |val| < 257/2
int16_t Util16::Center_257(int16_t val){
    val = Mod_257(val);
    if (val > 128) { // floor(257/2) = 128 
        val = val - 257;
    }
    return val;
}

void Util16::Norm_AddSub(int16_t * vec1, int16_t * vec2){
    for (size_t i = 0; i < 8; i++){
        addSub(vec1+i,vec2+i);
    }
}


void Util16::Norm_AddMult(int16_t * vec1, int16_t * vec2, int16_t * vec3){
    for (size_t i = 0; i < 8; i++){
        vec1[i] = vec1[i] + (vec2[i] * vec3[i]);
    }
}


void Util16::Norm_LShift(int16_t * vec, int16_t shift){
    for (size_t i = 0; i < 8; i++){
        vec[i] = vec[i] << shift;
    }
}


void Util16::Norm_Mult(int16_t * vec1, int16_t * vec2, int16_t * prod){
    for (size_t i = 0; i < 8; i++){
        prod[i] = vec1[i] * vec2[i];
    }
}


void Util16::Norm_Mod257(int16_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Mod_257(vec[i]);
    }
}


void Util16::Norm_Q_reduce(int16_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Q_reduce(vec[i]);
    }
}

void Util16::Norm_Center257(int16_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Center_257(vec[i]);
    }
}
    

void Util16::SIMD_AddSub(int16_t* vec1, int16_t* vec2){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec1);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec2);
	__m128i reg_result = _mm_add_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec1, reg_result);

	reg_result = _mm_sub_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec2, reg_result);

}


void Util16::SIMD_LShift(int16_t* vec, int16_t shift){

	__m128i reg = _mm_loadu_si128((__m128i *)vec);
	reg = _mm_slli_epi16(reg,shift);
	_mm_storeu_si128((__m128i*)vec, reg);
}


void Util16::SIMD_Mult(int16_t* vec1, int16_t* vec2, int16_t * product){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec1);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec2);
	reg1 = _mm_mullo_epi16(reg1,reg2);
	_mm_storeu_si128((__m128i*)product, reg1);
}


void Util16::SIMD_AddMult(int16_t* vec1, int16_t* vec2, int16_t* vec3){

	__m128i reg1 = _mm_loadu_si128((__m128i *)vec2);
	__m128i reg2 = _mm_loadu_si128((__m128i *)vec3);

	reg2 = _mm_mullo_epi16(reg1,reg2);
	reg1 = _mm_loadu_si128((__m128i *)vec1);
	reg1 = _mm_add_epi16(reg1, reg2);
	_mm_storeu_si128((__m128i*)vec1, reg1);

}


void Util16::SIMD_Q_reduce(int16_t* vec1){
	const __m128i TFF = _mm_set1_epi16(255);

	__m128i l_reg = _mm_loadu_si128((__m128i *)vec1);
	__m128i r_reg = _mm_srai_epi16(l_reg,8);
	l_reg = _mm_and_si128(l_reg,TFF);
	l_reg = _mm_sub_epi16(l_reg, r_reg);

	_mm_storeu_si128((__m128i*)vec1, l_reg);
}


void Util16::SIMD_Mod257(int16_t* vec1){
	const __m128i NEG_ONE = _mm_set1_epi16(-1);
	const __m128i NEG_TFS = _mm_set1_epi16(-257);

	SIMD_Q_reduce(vec1);
	SIMD_Q_reduce(vec1);

	__m128i l_reg = _mm_loadu_si128((__m128i *)vec1);
	__m128i r_reg = _mm_cmpeq_epi16(l_reg,NEG_ONE);
	r_reg = _mm_and_si128(r_reg, NEG_TFS);
	l_reg = _mm_xor_si128(l_reg,r_reg);
	_mm_storeu_si128((__m128i*)vec1, l_reg);
}