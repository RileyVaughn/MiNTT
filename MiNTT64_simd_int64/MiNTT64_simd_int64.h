#include <cstddef>
#include "sizes.h"
#include "MiNTT.h"
#include "util_int64.h"
#include <cstdint>
class MiNTT64_SIMD_int64 : public MiNTT{
    
    //Everything is public because I am lazy, bad practice
    public:

    static const int64_t n = 64;
    static const int64_t m = INPUT_SIZE/n*8;
    static const int64_t N = OUTPUT_SIZE/9*8; //assumes q=257
    static const int64_t d = N/n;
    static const int64_t q = 257;

    static const int64_t ndiv8 = n/8;
    
    int64_t A[m][d][ndiv8][8];
    int64_t NTT8_TABLE[256][8];
    int64_t MULT_TABLE[ndiv8][8];
    
    MiNTT64_SIMD_int64();
    void Setup();
    void Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]);
    void PrintKey(std::string filename);

    void ncc(uint8_t input[ndiv8], int64_t intermed[ndiv8][8]);
    void ntt_sum(uint8_t input[INPUT_SIZE], int64_t out[d][ndiv8][8]);
    void change_base(int64_t val[d][ndiv8][8], uint8_t hash[OUTPUT_SIZE]);

};