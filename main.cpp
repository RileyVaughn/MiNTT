#include <iostream>
#include "MiNTT64_norm_int64.h"
#include "MiNTT64_simd_int64.h"
#include "MiNTT128_norm_int64.h"
#include "MiNTT128_simd_int64.h"
#include "MiNTT64_norm_int16.h"
#include "MiNTT64_simd_int16.h"
#include "util_int64.h"

using namespace std;

void PrintOut(uint8_t output[OUTPUT_SIZE]);
void Print8x8(int64_t out[8][8]);
void Print8x8_2(int16_t out[8][8]);
void GenInput(uint8_t input[INPUT_SIZE]);

int main() {

    //MiNTT64_norm_int64 hash = MiNTT64_norm_int64();
    // MiNTT64_SIMD_int64 hash2 = MiNTT64_SIMD_int64();
    // MiNTT128_norm_int64 hash3 = MiNTT128_norm_int64();
    // MiNTT128_SIMD_int64 hash4 = MiNTT128_SIMD_int64();

    MiNTT64_norm_int16 hash5 = MiNTT64_norm_int16();
    MiNTT64_SIMD_int16 hash6 = MiNTT64_SIMD_int16();

    uint8_t input[INPUT_SIZE];
    uint8_t output[OUTPUT_SIZE] = {0};
    
    // std::fill(std::begin(output), std::end(output), 0);
    // GenInput(input);
    // hash.Hash(input,output);
    // PrintOut(output);


    // std::fill(std::begin(output), std::end(output), 0);
    // GenInput(input);
    // hash2.Hash(input,output);
    // PrintOut(output);

    // std::fill(std::begin(output), std::end(output), 0);
    // GenInput(input);
    // hash3.Hash(input,output);
    // PrintOut(output);

    // std::fill(std::begin(output), std::end(output), 0);
    // GenInput(input);
    // hash4.Hash(input,output);
    // PrintOut(output);

    std::fill(std::begin(output), std::end(output), 0);
    GenInput(input);
    hash5.Hash(input,output);
    PrintOut(output);

    std::fill(std::begin(output), std::end(output), 0);
    GenInput(input);
    hash6.Hash(input,output);
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
        if(i%54==53) {
            cout << endl;
        }

    }
    cout << endl << endl;

}

void Print8x8(int64_t out[8][8]){

    for (size_t i = 0; i < 8; i++)
    {
        for (size_t j = 0; j < 8; j++) {
            cout << out[i][j] % 257 << " ";
        }
        cout << endl;
    }
    cout << endl;

}

void Print8x8_2(int16_t out[8][8]){

    for (size_t i = 0; i < 8; i++)
    {
        for (size_t j = 0; j < 8; j++) {
            cout << out[i][j] % 257 << " ";
        }
        cout << endl;
    }
    cout << endl;

}