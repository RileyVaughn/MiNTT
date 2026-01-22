#include <cstddef>
#include "sizes.h"
#include "MiNTT.h"
#include "util_int64.h"
#include <cstdint>
#include <chrono>
class BenchMark_MiNTT128_norm_int64_QF4 : public MiNTT{
    
    //Everything is public because I am lazy, bad practice
    public:

    static const int64_t n = 128;
    static const int64_t m = INPUT_SIZE_QF4/n*8;
    static const int64_t N = OUTPUT_SIZE_QF4/17*8; //assumes q=65537
    static const int64_t d = N/n;
    static const int64_t q = 65537;
    //2nth root of unity
    static const int64_t omega = 59963;

    static const int64_t ndiv8 = n/8;
    
    int64_t A[m][d][ndiv8][8];
    int64_t NTT8_TABLE[256][8];
    int64_t MULT_TABLE[ndiv8][8];

    //BenchMark params
    int64_t lookup_table_time = 0;
    int64_t modulo_time = 0;
    int64_t other_ntt_time = 0;
    int64_t key_combine_time = 0;
    int64_t base_change_time = 0;

    
    BenchMark_MiNTT128_norm_int64_QF4();
    void Setup();
    void Hash(uint8_t input[INPUT_SIZE_QF4],uint8_t out[OUTPUT_SIZE_QF4]);
    void PrintKey(std::string filename);

    void ncc(uint8_t input[ndiv8], int64_t intermed[ndiv8][8]);
    void ntt_sum(uint8_t input[INPUT_SIZE_QF4], int64_t out[d][ndiv8][8]);
    void change_base(int64_t val[d][ndiv8][8], uint8_t hash[OUTPUT_SIZE_QF4]);

    void PrintBenchMark();
    void BenchMark(float & lookup_table_ratio, float & modulo_ratio, float & other_ntt_ratio, float & key_combine_ratio, float & base_change_ratio);
};

