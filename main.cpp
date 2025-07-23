#include <iostream>
#include "MiNTT64_norm_int64/MiNTT64_norm_int64.h"
#include "MiNTT64_simd_int64/MiNTT64_simd_int64.h"
#include "util/util_int64.h"

using namespace std;

void PrintOut(uint8_t output[OUTPUT_SIZE]);
void GenInput(uint8_t input[INPUT_SIZE]);

int main() {

    // MiNTT64_norm_int64 hash = MiNTT64_norm_int64();
    MiNTT64_SIMD_int64 hash2 = MiNTT64_SIMD_int64();

    uint8_t input[INPUT_SIZE];
    uint8_t output[OUTPUT_SIZE] = {0};
    GenInput(input);

    hash2.Hash(input,output);

    PrintOut(output);



    return 0;
}


void GenInput(uint8_t input[INPUT_SIZE]){

    for (size_t i = 0; i < INPUT_SIZE; i++)
    {
    input[i] = i % 256;
    }

}


void PrintOut(uint8_t output[OUTPUT_SIZE]){

    for (size_t i = 0; i < OUTPUT_SIZE; i++)
    {
        cout << int(output[i]) << " ";
    }
    cout << endl;

}