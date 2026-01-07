#include <cstddef>
#include "sizes.h"
#include "util_int16.h"
#include <cstdint>
class SWIFFT {
    
    //Everything is public because I am lazy, bad practice
    public:

    static const int16_t SWIFFT_IN_SIZE = 128;
    static const int16_t SWIFFT_OUT_SIZE = 66;

    static const int16_t n = 64;
    static const int16_t m = 16;
    static const int16_t q = 257;
    //2nth root of unity
    static const int64_t omega = 42;

    static const int16_t ndiv8 = n/8;
    
    int16_t A[m][ndiv8][8];
    int16_t NTT8_TABLE[256][8];
    int16_t MULT_TABLE[ndiv8][8];
    
    SWIFFT();
    void Setup();
    void Hash(uint8_t input[SWIFFT_IN_SIZE],uint8_t out[SWIFFT_OUT_SIZE]);
    void PrintKey(std::string filename);

    void ncc(uint8_t input[ndiv8], int16_t intermed[ndiv8][8]);
    void ntt_sum(uint8_t input[SWIFFT_IN_SIZE], int16_t out[ndiv8][8]);
    void change_base(int16_t val[ndiv8][8], uint8_t hash[SWIFFT_OUT_SIZE]);

};