#include <stdint.h>
#include <immintrin.h>

namespace Util64 {

// Efficient mod of 257 (or any q=2^x+1 if modified).
// The tradeoff is that the reduced range is not 0 to 256, but rather ~(the inital range of val >> 8)
int64_t Q_reduce(int64_t val){
	return (val & 255) - (val >> 8);
}

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
int64_t Mod_257(int64_t val) {
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

// Not particulalry efficent: raises Integer b to the power x mod q. 
// Only used in setup so can be slow.
int64_t IntPow(int64_t b, int64_t x, int64_t q) {
    int64_t result = 1;
    for (size_t i = 0; i < x; i++){
        result = (result * b) % q;
    }
    return result;
}

// Reverse the bits of an int up to a bound
// Taken from "https://github.com/micciancio/SWIFFT"
int64_t Bit_Rev(int64_t i, int64_t bound){
    int64_t irev = 0;
    for (size_t i = i | bound; i > 1; i = i >> 1){
        irev = (irev << 1) | (i & 1);
    }
    return irev;
}

// Takes integers a and b as input, returns a+b and a-b in place respectively
int64_t addSub(int64_t * a, int64_t * b) {
    int64_t temp = *b;
    *b = *a - *b;
    *a = *a + temp;
}

// Simulates the SIMD AddSub function manually
void Norm_AddSub(int64_t * vec1, int64_t * vec2){
    for (size_t i = 0; i < 8; i++){
        addSub(vec1+i,vec2+i);
    }
}

// Simulates the SIMD AddMult function manually, sum is returned in place of vec1.
void Norm_AddMult(int64_t * vec1, int64_t * vec2, int64_t * vec3){
    for (size_t i = 0; i < 8; i++){
        vec1[i] = vec1[i] + (vec2[i] * vec3[i]);
    }
}

// Simulates the SIMD Left Shift function manually
void Norm_LShift(int64_t * vec, int64_t shift){
    for (size_t i = 0; i < 8; i++){
        vec[i] = vec[i] << shift;
    }
}

// Simulates the SIMD Mult function manually. Prodcut is returned in postion of vec1.
void Norm_Mult(int64_t * vec1, int64_t * vec2){
    for (size_t i = 0; i < 8; i++){
        vec1[i] = vec1[i] * vec2[i];
    }
}

// Simulates the SIMD Mod257 function manually.
void Norm_Mod257(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Mod_257(vec[i]);
    }
}

// Simulates the SIMD Q_reduce function manually.
void Norm_Q_reduce(int64_t * vec){
    for (size_t i = 0; i < 8; i++){
       vec[i] = Q_reduce(vec[i]);
    }
}
    
};