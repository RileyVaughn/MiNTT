#include <iostream>
#include "MiNTT64_norm_int64/MiNTT64_norm_int64.h"

using namespace std;

void print_array(int64_t arr[8][8]);

int main() {
    
    uint8_t in[8] = {1,2,3,4,5,6,7,8};
    int64_t out[8][8];


    MiNTT64_norm_int64 hash = MiNTT64_norm_int64();


    hash.ncc(in,out);

    print_array(out);
    


    return 0;
}

void print_array(int64_t arr[8][8]) {
    for (size_t i = 0; i < 8; i++) {
        for (size_t j = 0; j < 8; j++) {
            std::cout << arr[i][j] << " ";
        }
        std::cout << "\n";
    }
}