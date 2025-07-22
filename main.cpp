#include <iostream>
#include "MiNTT64_norm_int64/MiNTT64_norm_int64.h"
#include "util/util_int64.h"

using namespace std;

void print_2darray(int64_t arr[8][8]);
void GenInput(uint8_t input[INPUT_SIZE]);

int main() {

    MiNTT64_norm_int64 hash = MiNTT64_norm_int64();

    uint8_t input[INPUT_SIZE];
    uint8_t output[OUTPUT_SIZE] = {0};
    GenInput(input);

    hash.Hash(input,output);
    // int64_t out[12][8][8];
    // hash.ntt_sum(input,out);
    
    // for (size_t i = 0; i < 12; i++)
    // {
    //     for (size_t j = 0; j < 8; j++)
    //     {
    //         for (size_t k = 0; k < 8; k++)
    //         {
    //             cout << Util64::Mod_257(out[i][j][k]) << " ";
    //         }
    //         cout << "| ";
    //     }
    // }
    // cout << endl;

    for (size_t i = 0; i < OUTPUT_SIZE; i++)
    {
        cout << int(output[i]) << " ";
    }
    cout << endl;



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

void GenInput(uint8_t input[INPUT_SIZE]){


for (size_t i = 0; i < INPUT_SIZE; i++)
{
    input[i] = i % 256;
}

}