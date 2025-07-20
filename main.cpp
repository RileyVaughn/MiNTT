#include <iostream>
#include "MiNTT64_norm_int64/MiNTT64_norm_int64.h"
#include "util/util_int64.h"

using namespace std;

void print_2darray(int64_t arr[8][8]);

int main() {
    
    uint8_t in[8] = {1,2,3,4,5,6,7,8};
    int64_t out[8][8];


    MiNTT64_norm_int64 hash = MiNTT64_norm_int64();


    hash.ncc(in,out);

    print_2darray(out);
    
    

    return 0;
}

void print_2darray(int64_t arr[8][8]) {
    for (size_t i = 0; i < 8; i++) {
        cout << i << ": ";
        for (size_t j = 0; j < 8; j++) {
            cout <<  Util64::Mod_257(arr[i][j]) << " ";
        }
        std::cout << "\n";
    }
}

// [[1 1 1 1 1 1 1 1]
// [197 137 17 34 68 136 15 30]
// [222 44 187 88 117 176 234 95]
// [44 117 95 165 246 35 169 23]
// [42 72 50 49 84 144 100 98]
// [50 98 79 124 58 52 215 113]
// [72 84 98 200 62 158 13 58]
// [49 200 124 118 104 157 195 198]]

// [1 16 4 64 2 32 8 128]
// [1 -16 64 4 8 -128 -2 32]
// [1 16 -4 -64 32 -2 -128 8]
// [1 -16 -64 -4 128 8 32 2]
// [1 16 4 64 -2 -32 -8 -128]
// [1 -16 64 4 -8 128 2 -32]
// [1 16 -4 -64 -32 2 128 -8]
// [1 -16 -64 -4 -128 -8 -32 -2]