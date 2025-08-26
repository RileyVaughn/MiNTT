#include <cstddef>
#include "sizes.h"
#include "MiNTT.h"
#include "util_int16.h"
#include <cstdint>
class MiNTT128_norm_int16 : public MiNTT{
    
    //Everything is public because I am lazy, bad practice
    public:

    static const int16_t n = 128;
    static const int16_t m = INPUT_SIZE/n*8;
    static const int16_t N = OUTPUT_SIZE/9*8; //assumes q=257
    static const int16_t d = N/n;
    static const int16_t q = 257;
    //2nth root of unity
    static const int64_t omega = 27;

    static const int16_t ndiv8 = n/8;
    
    int16_t A[m][d][ndiv8][8];
    int16_t NTT8_TABLE[256][8];
    int16_t MULT_TABLE[ndiv8][8];
    
    MiNTT128_norm_int16();
    void Setup();
    void Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]);
    void PrintKey(std::string filename);

    void ncc(uint8_t input[ndiv8], int16_t intermed[ndiv8][8]);
    void ntt_sum(uint8_t input[INPUT_SIZE], int16_t out[d][ndiv8][8]);
    void change_base(int16_t val[d][ndiv8][8], uint8_t hash[OUTPUT_SIZE]);

};