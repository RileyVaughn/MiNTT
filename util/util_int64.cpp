#include <util_int64.h>




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
    for (size_t i = i | bound; i > 1; i = i >> 1){
        irev = (irev << 1) | (i & 1);
    }
    return irev;
}


int64_t Util64::addSub(int64_t * a, int64_t * b) {
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


void Util64::Norm_Mult(int64_t * vec1, int64_t * vec2){
    for (size_t i = 0; i < 8; i++){
        vec1[i] = vec1[i] * vec2[i];
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
    