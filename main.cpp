#include <iostream>
#include <chrono>
#include <stdlib.h>
#include "MiNTT64_norm_int64.h"
#include "MiNTT64_simd_int64.h"
#include "MiNTT128_norm_int64.h"
#include "MiNTT128_simd_int64.h"
#include "MiNTT64_norm_int16.h"
#include "MiNTT64_simd_int16.h"
#include "MiNTT128_norm_int16.h"
#include "MiNTT128_simd_int16.h"
#include "MiNTT8_norm_int16.h"
#include "MiNTT8_simd_int16.h"
#include "util_int64.h"

using namespace std;

void PrintOut(uint8_t output[OUTPUT_SIZE]);
void GenInput(uint8_t input[INPUT_SIZE]);
int64_t CheckRuntime(uint8_t input[INPUT_SIZE], MiNTT * hash);
int64_t MeanRuntime(MiNTT * hash);

const int TEST_SIZE = 10000;


int main() {

    MiNTT * norm64_64 = new MiNTT64_norm_int64();
    MiNTT * simd64_64 = new MiNTT64_SIMD_int64();
    MiNTT * norm128_64 = new MiNTT128_norm_int64();
    MiNTT * simd128_64 = new MiNTT128_SIMD_int64();

    MiNTT * norm64_16 = new MiNTT64_norm_int16();
    MiNTT * simd64_16 = new MiNTT64_SIMD_int16();
    MiNTT * norm128_16 = new MiNTT128_norm_int16();
    MiNTT * simd128_16 = new MiNTT128_SIMD_int16();

    MiNTT * norm8_16 = new MiNTT8_norm_int16();
    MiNTT * simd8_16 = new MiNTT8_SIMD_int16();

    cout << "norm64_64: " << MeanRuntime(norm64_64) << endl;
    cout << "simd64_64: " << MeanRuntime(simd64_64) << endl;
    cout << "norm128_64: " << MeanRuntime(norm128_64) << endl;
    cout << "simd128_64: " << MeanRuntime(simd128_64) << endl;
    
    cout << "norm64_16: " << MeanRuntime(norm64_16) << endl;
    cout << "simd64_16: " << MeanRuntime(simd64_16) << endl;
    cout << "norm128_16: " << MeanRuntime(norm128_16) << endl;
    cout << "simd128_16: " << MeanRuntime(simd128_16) <<  endl;

    cout << "norm8_16: " << MeanRuntime(norm8_16) << endl;
    cout << "simd8_16: " << MeanRuntime(simd8_16) << endl;






    delete(norm64_64);
    delete(simd64_64);
    delete(norm128_64);
    delete(simd128_64);
    delete(norm64_16);
    delete(simd64_16);
    delete(norm128_16);
    delete(simd128_16);
    delete(norm8_16);
    delete(simd8_16);
    return 0;
}

// Cheks how long Hash() takes to execute in nanoseconds
int64_t CheckRuntime(uint8_t input[INPUT_SIZE], MiNTT * hash) {

    using namespace std::chrono;
    uint8_t output[OUTPUT_SIZE] = {0};

    auto start = high_resolution_clock::now();
    hash->Hash(input, output);
    auto end = high_resolution_clock::now();

    return duration_cast<nanoseconds>(end - start).count();

}

int64_t MeanRuntime(MiNTT * hash) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);

    int64_t sum = 0;
    uint8_t input[INPUT_SIZE]; 

    for (size_t i = 0; i < TEST_SIZE; i++)
    {
        GenInput(input);
        sum = sum + CheckRuntime(input,hash);
    }

    return sum/TEST_SIZE;
}


void GenInput(uint8_t input[INPUT_SIZE]){

    for (size_t i = 0; i < INPUT_SIZE; i++)
    {
        input[i] = rand() % 256;
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

