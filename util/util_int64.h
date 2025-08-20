#include <stdint.h>
#include <immintrin.h>


namespace Util64 {

// Efficient mod of 257 
// The tradeoff is that the reduced range is not 0 to 256, but rather the range is shifted left by 8.
int64_t Q_reduce(int64_t val);

// Efficient mod of 2^2^4+1 = 65537 
// The tradeoff is that the reduced range is not 0 to 256, but rather the range is shifted left by 8.
int64_t QF4_reduce(int64_t val);

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
int64_t Mod_257(int64_t val);

// Less Efficient, but still efficient mod of 65537
int64_t Mod_65537(int64_t val);

// Not particulalry efficent: raises Integer b to the power x mod q. 
// Only used in setup so can be slow.
int64_t IntPow(int64_t b, int64_t x, int64_t q);

// Reverse the bits of an int up to a bound
// Taken from "https://github.com/micciancio/SWIFFT"
int64_t Bit_Rev(int64_t i, int64_t bound);

// Takes integers a and b as input, returns a+b and a-b in place respectively
void addSub(int64_t * a, int64_t * b);
// Returns the bit representation of the the byte input

// Extracts the bits from a byte and returns a length 8 array
void BitsFromByte(int64_t byte, int64_t bits[8]);

//Takes an omega, a q, and a matrix to return values
void Gen8NCCMat(int64_t omega, int64_t q, int64_t mat[8][8]);

//This generates the table of possible interim 8-sum NTT values
void GenNTT8Table(int64_t omega, int64_t q, int64_t table[256][8]);

//This generates the possible Mult values table
int64_t * GenMultTable(int64_t omega, int64_t n, int64_t q);

//Generates a random key, and stores it in class variable A
int64_t * GenKey(int64_t m, int64_t n, int64_t d, int64_t q);

// Simulates the SIMD AddSub function manually
void Norm_AddSub(int64_t * vec1, int64_t * vec2);

// Simulates the SIMD AddMult function manually, sum is returned in place of vec1.
void Norm_AddMult(int64_t * vec1, int64_t * vec2, int64_t * vec3);

// Simulates the SIMD Left Shift function manually
void Norm_LShift(int64_t * vec, int64_t shift);

// Simulates the SIMD Mult function manually. Prodcut is returned in postion of prod.
void Norm_Mult(int64_t * vec1, int64_t * vec2, int64_t * prod);

// Simulates the SIMD Mod257 function manually.
void Norm_Mod257(int64_t * vec);

// Simulates the SIMD Mod65537 function manually.
void Norm_Mod65537(int64_t * vec);

// Simulates the SIMD Q_reduce function manually.
void Norm_Q_reduce(int64_t * vec);

// Simulates the SIMD QF4_reduce function manually.
void Norm_QF4_reduce(int64_t * vec);

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
void SIMD_AddSub(int64_t* vec1, int64_t* vec2);

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
void SIMD_LShift(int64_t* vec, int64_t shift);

//Multiples two vectors element-wise using SIMD instrucitons, returns product
void SIMD_Mult(int64_t* vec1, int64_t* vec2, int64_t* product);


void SIMD_AddMult(int64_t* vec1, int64_t* vec2, int64_t* vec3);

// Efficient SIMD mod of 257 (or any q=2^x+1 if modified).
// The tradeoff is that the reduced range is not 0 to 256, but rather the range is shifted left by 8.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
void SIMD_Q_reduce(int64_t* vec1);

//Efficient SIMD mod of 65537, range is not 0 to 65535 but rather shifted by 16
void SIMD_QF4_reduce(int64_t* vec1);

// Less Efficient SIMD mod, but still efficient mod of 257 (or any q=2^x+1 if modified).
// Running q_reduce multiple times minimizes the the domain to -1 to 256, the following code removes the -1.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
void SIMD_Mod257(int64_t* vec1);

// Less Efficient SIMD mod, but still efficient mod of 65537
void SIMD_Mod65537(int64_t* vec1);

};

