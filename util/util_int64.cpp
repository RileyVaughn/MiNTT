#include <util_int64.h>
#include <stdlib.h>
#include <iostream>



int64_t Util64::Q_reduce(int64_t val){
	return (val & 255) - (val >> 8);
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
    Util64::Gen8NCCMat(omega,q,ncc_mat);

    for (size_t i = 0; i < 256; i++){
        for (size_t j = 0; j < 8; j++){
            int64_t vec[8];
            Util64::BitsFromByte(i, vec);
            for (size_t k = 0; k < 8; k++){
                table[i][j] = Mod_257(table[i][j] + ncc_mat[j][k]*vec[k]);
            }
        }
    }

}


void Util64::GenMultTable(int64_t omega, int64_t n, int64_t q, int64_t table[8][8]){

    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; k < 8; k++){
            table[k][i] = IntPow(omega,(Bit_Rev(k,8)*(2*i+1))%(2*n),q);
        }
    }
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
    