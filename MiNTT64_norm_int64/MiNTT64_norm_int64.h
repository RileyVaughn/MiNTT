#include <cstddef>
#include <../sizes.h>
#include <../util/util_int64.h>
#include <cstdint>
class MiNTT64_norm_int64 {

    public:

    [OUTPUT_SIZE]std::byte Hash();

    void Setup();

    private:

    static const int64_t n = 64;
    static const int64_t m = 216;
    static const int64_t N = OUTPUT_SIZE/9*8; //assumes q=257
    static const int64_t d = N/n;
    static const int64_t q = 257;

    static const int64_t ndiv8 = n/8;

    int64_t A[m][d][ndiv8][8];
    int64_t NTT8_TABLE[256][8];
    int64_t MULT_TABLE[8][8];


    void genKey();
    void genNTT8Table(int64_t omega);
    void gen8NCCMat(int64_t omega);
    void genMultTable(int64_t omega);


};