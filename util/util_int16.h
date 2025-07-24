#include <stdint.h>
#include <immintrin.h>


namespace Util16 {

int16_t Q_reduce(int16_t val);

// Less Efficient, but still efficient mod of 257 (or any q=2^x+1 if modified).
int16_t Mod_257(int16_t val);

// Not particulalry efficent: raises Integer b to the power x mod q. 
// Only used in setup so can be slow.
int16_t IntPow(int16_t b, int16_t x, int16_t q);

// Reverse the bits of an int up to a bound
// Taken from "https://github.com/micciancio/SWIFFT"
int16_t Bit_Rev(int16_t i, int16_t bound);

// Takes integers a and b as input, returns a+b and a-b in place respectively
void addSub(int16_t * a, int16_t * b);
// Returns the bit representation of the the byte input

// Extracts the bits from a byte and returns a length 8 array
void BitsFromByte(int16_t byte, int16_t bits[8]);

//Takes an omega, a q, and a matrix to return values
void Gen8NCCMat(int16_t omega, int16_t q, int16_t mat[8][8]);

//This generates the table of possible interim 8-sum NTT values
void GenNTT8Table(int16_t omega, int16_t q, int16_t table[256][8]);

//This generates the possible Mult values table
int16_t * GenMultTable(int16_t omega, int16_t n, int16_t q);

//Generates a random key, and stores it in class variable A
int16_t * GenKey(int16_t m, int16_t n, int16_t d, int16_t q);

// Centers val such that |val| < 257/2
int16_t Center_257(int16_t val);

// Simulates the SIMD AddSub function manually
void Norm_AddSub(int16_t * vec1, int16_t * vec2);

// Simulates the SIMD AddMult function manually, sum is returned in place of vec1.
void Norm_AddMult(int16_t * vec1, int16_t * vec2, int16_t * vec3);

// Simulates the SIMD Left Shift function manually
void Norm_LShift(int16_t * vec, int16_t shift);

// Simulates the SIMD Mult function manually. Prodcut is returned in postion of prod.
void Norm_Mult(int16_t * vec1, int16_t * vec2, int16_t * prod);

// Simulates the SIMD Mod257 function manually.
void Norm_Mod257(int16_t * vec);

// Simulates the SIMD Q_reduce function manually.
void Norm_Q_reduce(int16_t * vec);

// Centers each vlaue around 0.
void Norm_Center257(int16_t * vec);

//Adds and subtracts vec1 and vec2 using SIMD, and returns the sum/difference in place resp.
void SIMD_AddSub(int16_t* vec1, int16_t* vec2);

//Left Shifts each element in vec by shift. Used to efficiently multiply by powers of 2.
void SIMD_LShift(int16_t* vec, int16_t shift);

//Multiples two vectors element-wise using SIMD instrucitons, returns product
void SIMD_Mult(int16_t* vec1, int16_t* vec2, int16_t* product);


void SIMD_AddMult(int16_t* vec1, int16_t* vec2, int16_t* vec3);

// Efficient SIMD mod of 257 (or any q=2^x+1 if modified).
// The tradeoff is that the reduced range is not 0 to 256, but rather the range is shifted left by 8.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
void SIMD_Q_reduce(int16_t* vec1);

// Less Efficient SIMD mod, but still efficient mod of 257 (or any q=2^x+1 if modified).
// Running q_reduce multiple times minimizes the the domain to -1 to 256, the following code removes the -1.
// Originally found in the SWIFFT source code <https://github.com/micciancio/SWIFFT>, modified for 64 bit
void SIMD_Mod257(int16_t* vec1);

// Centers each aluein vec around 0 [-128 to 128] using SIMD instrucitons.
void SIMD_Center257(int16_t * vec);

};