#include <stdint.h>
#include <immintrin.h>


namespace Util64 {

int64_t Q_reduce(int64_t val);

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
int64_t Mod_257(int64_t val);

// Not particulalry efficent: raises Integer b to the power x mod q. 
// Only used in setup so can be slow.
int64_t IntPow(int64_t b, int64_t x, int64_t q);

// Reverse the bits of an int up to a bound
// Taken from "https://github.com/micciancio/SWIFFT"
int64_t Bit_Rev(int64_t i, int64_t bound);

// Takes integers a and b as input, returns a+b and a-b in place respectively
int64_t addSub(int64_t * a, int64_t * b);
// Returns the bit representation of the the byte input

// Extracts the bits from a byte and returns a length 8 array
void BitsFromByte(int64_t byte, int64_t bits[8]);

//Takes an omega, a q, and a matrix to return values
void Gen8NCCMat(int64_t omega, int64_t q, int64_t mat[8][8]);

// Simulates the SIMD AddSub function manually
void Norm_AddSub(int64_t * vec1, int64_t * vec2);

// Simulates the SIMD AddMult function manually, sum is returned in place of vec1.
void Norm_AddMult(int64_t * vec1, int64_t * vec2, int64_t * vec3);

// Simulates the SIMD Left Shift function manually
void Norm_LShift(int64_t * vec, int64_t shift);

// Simulates the SIMD Mult function manually. Prodcut is returned in postion of vec1.
void Norm_Mult(int64_t * vec1, int64_t * vec2);

// Simulates the SIMD Mod257 function manually.
void Norm_Mod257(int64_t * vec);

// Simulates the SIMD Q_reduce function manually.
void Norm_Q_reduce(int64_t * vec);
    

};