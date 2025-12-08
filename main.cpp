#include <iostream>
#include <chrono>
#include <stdlib.h>
#include <x86intrin.h>
#include <algorithm>
#include <cmath>
#include <sys/resource.h>
#include <cryptopp/sha.h>
#include <cryptopp/hex.h>
#include <cryptopp/filters.h>



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
#include "MiNTT128_norm_int64_QF4.h"
#include "MiNTT128_simd_int64_QF4.h"
#include "MiNTT64_norm_int64_QF4.h"
#include "MiNTT64_simd_int64_QF4.h"

#include "util_int64.h"

using namespace std;

void PrintOut(uint8_t output[OUTPUT_SIZE]);
void GenInput(uint8_t input[INPUT_SIZE]);
int64_t CheckRuntime(uint8_t input[INPUT_SIZE], MiNTT * hash);
int64_t MeanRuntime(MiNTT * hash);

void PrintOutQF4(uint8_t output[OUTPUT_SIZE]);
void GenInputQF4(uint8_t input[INPUT_SIZE]);
int64_t CheckRuntimeQF4(uint8_t input[INPUT_SIZE], MiNTT * hash);
int64_t MeanRuntimeQF4(MiNTT * hash);

void MeanSTDRuntime(MiNTT * hash, int64_t & mean, int64_t & std);
int64_t CalcSTD(int64_t times[INPUT_SIZE], int64_t mean, int64_t & std);

uint64_t MeasureCycles(uint8_t input[INPUT_SIZE], MiNTT * hash);
int64_t MeanCycles(MiNTT * hash);
int64_t MedianCycles(MiNTT * hash);

int64_t CheckMemory(uint8_t input[INPUT_SIZE], MiNTT * hash);


const int TEST_SIZE = 1000;


int main() {

    
    // MiNTT * norm64_64 = new MiNTT64_norm_int64();


    // // int64_t mean = 0;
    // // int64_t std = 0;
    // // MeanSTDRuntime(norm64_64,mean,std);
    // int64_t med_cycles = MedianCycles(norm64_64);
    // cout << "norm64_64: " << med_cycles << endl;

    // cout << "norm64_64: " << mean << " " << std << endl;
    
    // uint8_t input[INPUT_SIZE];
    // GenInput(input);
    // cout << "norm64_64:" << CheckMemory(input,norm64_64)<< endl;

    std::string input = "Hello World!";
    std::string hash;

    CryptoPP::SHA256 sha256;

    CryptoPP::StringSource(input, true,
        new CryptoPP::HashFilter(sha256,
            new CryptoPP::HexEncoder(
                new CryptoPP::StringSink(hash)
            )
        )
    );

    std::cout << "SHA-256: " << hash << std::endl;


    
    return 0;
}

//////////////////////////////// Runtimes /////////////////////////////////////

// void CheckRuntimeMeans(){

//     MiNTT * norm64_64 = new MiNTT64_norm_int64();
//     MiNTT * simd64_64 = new MiNTT64_SIMD_int64();
//     MiNTT * norm128_64 = new MiNTT128_norm_int64();
//     MiNTT * simd128_64 = new MiNTT128_SIMD_int64();

//     MiNTT * norm64_16 = new MiNTT64_norm_int16();
//     MiNTT * simd64_16 = new MiNTT64_SIMD_int16();
//     MiNTT * norm128_16 = new MiNTT128_norm_int16();
//     MiNTT * simd128_16 = new MiNTT128_SIMD_int16();

//     // MiNTT * norm8_16 = new MiNTT8_norm_int16();
//     // MiNTT * simd8_16 = new MiNTT8_SIMD_int16();

//     MiNTT * norm128_64_QF4 = new MiNTT128_norm_int64_QF4();
//     MiNTT * simd128_64_QF4 = new MiNTT128_SIMD_int64_QF4();
//     MiNTT * norm64_64_QF4 = new MiNTT64_norm_int64_QF4();
//     MiNTT * simd64_64_QF4 = new MiNTT64_SIMD_int64_QF4();

//     cout << "norm64_64: " << MeanRuntime(norm64_64) << endl;
//     cout << "simd64_64: " << MeanRuntime(simd64_64) << endl;
//     cout << "norm128_64: " << MeanRuntime(norm128_64) << endl;
//     cout << "simd128_64: " << MeanRuntime(simd128_64) << endl;
    
//     cout << "norm64_16: " << MeanRuntime(norm64_16) << endl;
//     cout << "simd64_16: " << MeanRuntime(simd64_16) << endl;
//     cout << "norm128_16: " << MeanRuntime(norm128_16) << endl;
//     cout << "simd128_16: " << MeanRuntime(simd128_16) <<  endl;

//     // cout << "norm8_16: " << MeanRuntime(norm8_16) << endl;
//     // cout << "simd8_16: " << MeanRuntime(simd8_16) << endl;

//     cout << "norm128_64_QF4: " << MeanRuntimeQF4(norm128_64_QF4) << endl;
//     cout << "simd128_64_QF4: " << MeanRuntimeQF4(simd128_64_QF4) << endl;
//     cout << "norm64_64_QF4: " << MeanRuntimeQF4(norm64_64_QF4) << endl;
//     cout << "simd64_64_QF4: " << MeanRuntimeQF4(simd64_64_QF4) << endl;


//     delete(norm64_64);
//     delete(simd64_64);
//     delete(norm128_64);
//     delete(simd128_64);
//     delete(norm64_16);
//     delete(simd64_16);
//     delete(norm128_16);
//     delete(simd128_16);
//     // delete(norm8_16);
//     // delete(simd8_16);
//     delete(norm128_64_QF4);
//     delete(simd128_64_QF4);
//     delete(norm64_64_QF4);
//     delete(simd64_64_QF4);
 

// }


// Checks how long Hash() takes to execute in nanoseconds
int64_t CheckRuntime(uint8_t input[INPUT_SIZE], MiNTT * hash) {

    using namespace std::chrono;
    uint8_t output[OUTPUT_SIZE] = {0};

    auto start = high_resolution_clock::now();
    hash->Hash(input, output);
    auto end = high_resolution_clock::now();

    return duration_cast<nanoseconds>(end - start).count();

}

int64_t CheckRuntimeQF4(uint8_t input[INPUT_SIZE_QF4], MiNTT * hash) {

    using namespace std::chrono;
    uint8_t output[OUTPUT_SIZE_QF4] = {0};

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

int64_t MeanRuntimeQF4(MiNTT * hash) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);

    int64_t sum = 0;
    uint8_t input[INPUT_SIZE_QF4]; 

    for (size_t i = 0; i < TEST_SIZE; i++)
    {
        GenInput(input);
        sum = sum + CheckRuntimeQF4(input,hash);
    }

    return sum/TEST_SIZE;
}

void MeanSTDRuntime(MiNTT * hash, int64_t & mean, int64_t & std) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);
    mean = 0;
    std = 0;

    uint8_t input[INPUT_SIZE];
    int64_t times[TEST_SIZE]; 

    for (size_t i = 0; i < TEST_SIZE; i++) {
        GenInput(input);
        times[i] = CheckRuntime(input,hash);
        mean += times[i];
    }
    mean /= TEST_SIZE;

    std = CalcSTD(times,mean,std);
    

}

//////////////////////////// Cycles /////////////////////////////////////////

uint64_t MeasureCycles(uint8_t input[INPUT_SIZE], MiNTT * hash) {

    unsigned aux;
    uint8_t output[OUTPUT_SIZE] = {0};

    _mm_lfence();
    uint64_t start = __rdtsc();

    hash->Hash(input, output);

    uint64_t end = __rdtscp(&aux);
    _mm_lfence();

    return end - start;
}

int64_t MeanCycles(MiNTT * hash) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);

    int64_t sum = 0;
    uint8_t input[INPUT_SIZE]; 

    for (size_t i = 0; i < TEST_SIZE; i++)
    {
        GenInput(input);
        sum = sum + MeasureCycles(input,hash);
    }

    return sum/TEST_SIZE;
}

int64_t MedianCycles(MiNTT * hash) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);

    int64_t times[TEST_SIZE];
    uint8_t input[INPUT_SIZE]; 

    for (size_t i = 0; i < TEST_SIZE; i++)
    {
        GenInput(input);
        times[i] = MeasureCycles(input,hash);
    }

    sort(times,times+TEST_SIZE);

    return times[TEST_SIZE/2];
}


//////////////////////////// Memory /////////////////////////////////////////

int64_t CheckMemory(uint8_t input[INPUT_SIZE], MiNTT * hash){

    uint8_t output[OUTPUT_SIZE] = {0};
    rusage usage_before, usage_after;

    getrusage(RUSAGE_SELF, &usage_before);
    hash->Hash(input, output);
    getrusage(RUSAGE_SELF, &usage_after);

    return usage_after.ru_maxrss - usage_before.ru_maxrss;

}



//////////////////////////// Input Output /////////////////////////////////////////

void GenInput(uint8_t input[INPUT_SIZE]){

    for (size_t i = 0; i < INPUT_SIZE; i++)
    {
        input[i] = rand() % 256;
    }

}

void GenInputQF4(uint8_t input[INPUT_SIZE_QF4]){

    for (size_t i = 0; i < INPUT_SIZE_QF4; i++)
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

void PrintOutQF4(uint8_t output[OUTPUT_SIZE_QF4]){

    for (size_t i = 0; i < OUTPUT_SIZE_QF4; i++)
    {
        cout << int(output[i]) << " ";
        if(i%54==53) {
            cout << endl;
        }

    }
    cout << endl << endl;

}

//////////////////////////// Math /////////////////////////////////////////

int64_t CalcSTD(int64_t times[INPUT_SIZE], int64_t mean, int64_t & std) {


    for (size_t i = 0; i < TEST_SIZE; i++) {
        int64_t diff = times[i]-mean;
        std += diff*diff;

    }
    std = std/TEST_SIZE;
    std = int64_t(sqrt(std));
    return std;
}
