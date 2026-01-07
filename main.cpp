#include <iostream>
#include <chrono>
#include <stdlib.h>
#include <x86intrin.h>
#include <algorithm>
#include <cmath>
#include <sys/resource.h>
#include <cryptopp/sha.h>
#include <cryptopp/sha3.h>
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
#include "SWIFFT/SWIFFT.h"

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

int64_t CheckRuntimeSHA256();

void CheckRuntimeMeanSTD();
void MeanSTDRuntime(MiNTT * hash, int64_t & mean, int64_t & std);
void MeanSTDRuntimeQF4(MiNTT * hash, int64_t & mean, int64_t & std);
int64_t CalcSTD(int64_t times[INPUT_SIZE], int64_t mean, int64_t & std);
int64_t CalcSTD_QF4(int64_t times[INPUT_SIZE_QF4], int64_t mean, int64_t & std);

// uint64_t MeasureCycles(uint8_t input[INPUT_SIZE], MiNTT * hash);
// int64_t MeanCycles(MiNTT * hash);
// int64_t MedianCycles(MiNTT * hash);

// uint64_t MeasureCyclesSHA256(uint8_t input[INPUT_SIZE]);
// int64_t MeanRuntimeSHA256();

void GenInputSWIFFT(uint8_t input[128]);
int64_t CheckRuntimeSWIFFT(uint8_t input[128], SWIFFT * hash);
int64_t MeanRuntimeSWIFFT(SWIFFT * hash);


const int TEST_SIZE = 100000;


int main() {

    SWIFFT * swifft = new SWIFFT();
    cout << "SWIFFT " << MeanRuntimeSWIFFT(swifft) << endl; 





    //CheckRuntimeMeanSTD();



    return 0;
}

//////////////////////////////// Runtimes /////////////////////////////////////

void CheckRuntimeMeans(){

    MiNTT * norm64_64 = new MiNTT64_norm_int64();
    MiNTT * simd64_64 = new MiNTT64_SIMD_int64();
    MiNTT * norm128_64 = new MiNTT128_norm_int64();
    MiNTT * simd128_64 = new MiNTT128_SIMD_int64();

    MiNTT * norm64_16 = new MiNTT64_norm_int16();
    MiNTT * simd64_16 = new MiNTT64_SIMD_int16();
    MiNTT * norm128_16 = new MiNTT128_norm_int16();
    MiNTT * simd128_16 = new MiNTT128_SIMD_int16();

    // MiNTT * norm8_16 = new MiNTT8_norm_int16();
    // MiNTT * simd8_16 = new MiNTT8_SIMD_int16();

    MiNTT * norm128_64_QF4 = new MiNTT128_norm_int64_QF4();
    MiNTT * simd128_64_QF4 = new MiNTT128_SIMD_int64_QF4();
    MiNTT * norm64_64_QF4 = new MiNTT64_norm_int64_QF4();
    MiNTT * simd64_64_QF4 = new MiNTT64_SIMD_int64_QF4();

    cout << "norm64_64: " << MeanRuntime(norm64_64) << endl;
    cout << "simd64_64: " << MeanRuntime(simd64_64) << endl;
    cout << "norm128_64: " << MeanRuntime(norm128_64) << endl;
    cout << "simd128_64: " << MeanRuntime(simd128_64) << endl;
    
    cout << "norm64_16: " << MeanRuntime(norm64_16) << endl;
    cout << "simd64_16: " << MeanRuntime(simd64_16) << endl;
    cout << "norm128_16: " << MeanRuntime(norm128_16) << endl;
    cout << "simd128_16: " << MeanRuntime(simd128_16) <<  endl;

    // cout << "norm8_16: " << MeanRuntime(norm8_16) << endl;
    // cout << "simd8_16: " << MeanRuntime(simd8_16) << endl;

    cout << "norm128_64_QF4: " << MeanRuntimeQF4(norm128_64_QF4) << endl;
    cout << "simd128_64_QF4: " << MeanRuntimeQF4(simd128_64_QF4) << endl;
    cout << "norm64_64_QF4: " << MeanRuntimeQF4(norm64_64_QF4) << endl;
    cout << "simd64_64_QF4: " << MeanRuntimeQF4(simd64_64_QF4) << endl;


    delete(norm64_64);
    delete(simd64_64);
    delete(norm128_64);
    delete(simd128_64);
    delete(norm64_16);
    delete(simd64_16);
    delete(norm128_16);
    delete(simd128_16);
    // delete(norm8_16);
    // delete(simd8_16);
    delete(norm128_64_QF4);
    delete(simd128_64_QF4);
    delete(norm64_64_QF4);
    delete(simd64_64_QF4);
 

}


void CheckRuntimeMeanSTD(){

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

    MiNTT * norm128_64_QF4 = new MiNTT128_norm_int64_QF4();
    MiNTT * simd128_64_QF4 = new MiNTT128_SIMD_int64_QF4();
    MiNTT * norm64_64_QF4 = new MiNTT64_norm_int64_QF4();
    MiNTT * simd64_64_QF4 = new MiNTT64_SIMD_int64_QF4();

    int64_t mean = 0;
    int64_t std = 0;
    
    MeanSTDRuntime(norm64_64,mean,std);
    cout << "norm64_64: " <<  mean << " " << std << endl;
    MeanSTDRuntime(simd64_64,mean,std);
    cout << "simd64_64: " <<  mean << " " << std << endl;
    MeanSTDRuntime(norm128_64,mean,std);
    cout << "norm128_64: " <<  mean << " " << std << endl;
    MeanSTDRuntime(simd128_64,mean,std);
    cout << "simd128_64: " <<  mean << " " << std << endl;

    MeanSTDRuntime(norm64_16,mean,std);
    cout << "norm64_16: " <<  mean << " " << std << endl;
    MeanSTDRuntime(simd64_16,mean,std);
    cout << "simd64_16: " <<  mean << " " << std << endl;
    MeanSTDRuntime(norm128_16,mean,std);
    cout << "norm128_16: " <<  mean << " " << std << endl;
    MeanSTDRuntime(simd128_16,mean,std);
    cout << "simd128_16: " <<  mean << " " << std << endl;

    MeanSTDRuntime(norm8_16,mean,std);
    cout << "norm8_16: " <<  mean << " " << std << endl;
    MeanSTDRuntime(simd8_16,mean,std);
    cout << "simd8_16: " <<  mean << " " << std << endl;

    MeanSTDRuntimeQF4(norm128_64_QF4,mean,std);
    cout << "norm128_64_QF4: " <<  mean << " " << std << endl;
    MeanSTDRuntimeQF4(simd128_64_QF4,mean,std);
    cout << "simd128_64_QF4: " <<  mean << " " << std << endl;
    MeanSTDRuntimeQF4(norm64_64_QF4,mean,std);
    cout << "norm64_64_QF4: " <<  mean << " " << std << endl;
    MeanSTDRuntimeQF4(simd64_64_QF4,mean,std);
    cout << "simd64_64_QF4: " <<  mean << " " << std << endl;


    delete(norm64_64);
    delete(simd64_64);
    delete(norm128_64);
    delete(simd128_64);
    delete(norm64_16);
    delete(simd64_16);
    delete(norm128_16);
    delete(simd128_16);
    // delete(norm8_16);
    // delete(simd8_16);
    delete(norm128_64_QF4);
    delete(simd128_64_QF4);
    delete(norm64_64_QF4);
    delete(simd64_64_QF4);
 

}


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

// int64_t CheckRuntimeSHA256(uint8_t input[INPUT_SIZE]) {

//     using namespace std::chrono;

//     CryptoPP::SHA256 sha256;
//     CryptoPP::byte digest[CryptoPP::SHA256::DIGESTSIZE];

//     auto start = high_resolution_clock::now();
//     CryptoPP::ArraySource( input, INPUT_SIZE, true,
//         new CryptoPP::HashFilter( sha256,
//             new CryptoPP::ArraySink(digest, CryptoPP::SHA256::DIGESTSIZE)
//         )
//     );
//     auto end = high_resolution_clock::now();
    
//     return duration_cast<nanoseconds>(end - start).count();

// }

int64_t CheckRuntimeSWIFFT(uint8_t input[128], SWIFFT * hash){

    using namespace std::chrono;
    uint8_t output[66] = {0};

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



int64_t MeanRuntimeSWIFFT(SWIFFT * hash){

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);

    int64_t sum = 0;
    uint8_t input[128]; 

    for (size_t i = 0; i < TEST_SIZE; i++)
    {
        GenInputSWIFFT(input);
        sum = sum + CheckRuntimeSWIFFT(input,hash);
    }

    return sum/TEST_SIZE;

}


// int64_t MeanRuntimeSHA256(){

//     //just a seed for random input gen, I init it here so that all funcs have the same input to test from
//     srand(1);

//     int64_t sum = 0;
//     uint8_t input[INPUT_SIZE]; 

//     for (size_t i = 0; i < TEST_SIZE; i++)
//     {
//         GenInput(input);
//         int64_t runtime = CheckRuntimeSHA256(input);
//         sum = sum + runtime;
//         // cout << int(input[0]) << " " <<runtime << endl;
//     }

//     return sum/TEST_SIZE;

// }


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

    std = CalcSTD_QF4(times,mean,std);
    

}

void MeanSTDRuntimeQF4(MiNTT * hash, int64_t & mean, int64_t & std) {

    //just a seed for random input gen, I init it here so that all funcs have the same input to test from
    srand(1);
    mean = 0;
    std = 0;

    uint8_t input[INPUT_SIZE_QF4];
    int64_t times[TEST_SIZE]; 

    for (size_t i = 0; i < TEST_SIZE; i++) {
        GenInputQF4(input);
        times[i] = CheckRuntimeQF4(input,hash);
        mean += times[i];
    }
    mean /= TEST_SIZE;

    std = CalcSTD(times,mean,std);
    

}

// //////////////////////////// Cycles /////////////////////////////////////////

// uint64_t MeasureCycles(uint8_t input[INPUT_SIZE], MiNTT * hash) {

//     unsigned aux;
//     uint8_t output[OUTPUT_SIZE] = {0};

//     _mm_lfence();
//     uint64_t start = __rdtsc();

//     hash->Hash(input, output);

//     uint64_t end = __rdtscp(&aux);
//     _mm_lfence();

//     return end - start;
// }

// uint64_t MeasureCyclesSHA256(uint8_t input[INPUT_SIZE]) {

//     unsigned aux;
//     CryptoPP::SHA256 sha256;
//     CryptoPP::byte digest[CryptoPP::SHA256::DIGESTSIZE];

//     _mm_lfence();
//     uint64_t start = __rdtsc();

//     CryptoPP::ArraySource(input, INPUT_SIZE, true,
//         new CryptoPP::HashFilter( sha256,
//             new CryptoPP::ArraySink(digest, CryptoPP::SHA256::DIGESTSIZE)
//         )
//     );



//     uint64_t end = __rdtscp(&aux);
//     _mm_lfence();

//     return end - start;

// }


// int64_t MeanCycles(MiNTT * hash) {

//     //just a seed for random input gen, I init it here so that all funcs have the same input to test from
//     srand(1);

//     int64_t sum = 0;
//     uint8_t input[INPUT_SIZE]; 

//     for (size_t i = 0; i < TEST_SIZE; i++)
//     {
//         GenInput(input);
//         sum = sum + MeasureCycles(input,hash);
//     }

//     return sum/TEST_SIZE;
// }

// int64_t MedianCycles(MiNTT * hash) {

//     //just a seed for random input gen, I init it here so that all funcs have the same input to test from
//     srand(1);

//     int64_t times[TEST_SIZE];
//     uint8_t input[INPUT_SIZE]; 

//     for (size_t i = 0; i < TEST_SIZE; i++)
//     {
//         GenInput(input);
//         times[i] = MeasureCycles(input,hash);
//     }

//     sort(times,times+TEST_SIZE);

//     return times[TEST_SIZE/2];
// }


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

// Quick and dirty swifft input gen
void GenInputSWIFFT(uint8_t input[128]){

    for (size_t i = 0; i < 128; i++)
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

int64_t CalcSTD_QF4(int64_t times[INPUT_SIZE_QF4], int64_t mean, int64_t & std) {


    for (size_t i = 0; i < TEST_SIZE; i++) {
        int64_t diff = times[i]-mean;
        std += diff*diff;

    }
    std = std/TEST_SIZE;
    std = int64_t(sqrt(std));
    return std;
}