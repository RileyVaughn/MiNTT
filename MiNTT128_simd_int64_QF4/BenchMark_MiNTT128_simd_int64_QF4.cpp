
#include <cstddef>
#include<string>
#include "BenchMark_MiNTT128_simd_int64_QF4.h"
#include <cstdint>
#include <iostream>
#include <fstream>
#include <chrono>

using namespace std::chrono;

BenchMark_MiNTT128_simd_int64_QF4::BenchMark_MiNTT128_simd_int64_QF4(){
    Setup();
}

void BenchMark_MiNTT128_simd_int64_QF4::Setup(){

    Util64::GenNTT8Table(Util64::IntPow(omega,8,q),q,NTT8_TABLE);


    int64_t * mult_table = Util64::GenMultTable(omega,n,q);
    for (size_t i = 0; i < ndiv8; i++){
        for (size_t j = 0; j < 8; j++){
            MULT_TABLE[i][j] = mult_table[i*8+j];
        }
    } 
    delete[] mult_table;

    int64_t * key = Util64::GenKey(m,n,d,q);
    for (size_t i = 0; i < m; i++){
       for (size_t j = 0; j < d; j++){
        for (size_t k = 0; k < ndiv8; k++){
            for (size_t l = 0; l < 8; l++){
                A[i][j][k][l] = key[i*(d*n)+j*n+k*8+l];
            }
        }
       }
    }
    delete[] key;
    
}   


void BenchMark_MiNTT128_simd_int64_QF4::Hash(uint8_t input[INPUT_SIZE_QF4],uint8_t out[OUTPUT_SIZE_QF4]){
    int64_t inter[d][ndiv8][8] = {0};
    lookup_table_time = 0;
    modulo_time = 0;
    other_ntt_time = 0;
    key_combine_time = 0;
    base_change_time = 0;


    ntt_sum(input,inter);
    change_base(inter,out);

    PrintBenchMark();

}



void BenchMark_MiNTT128_simd_int64_QF4::ncc(uint8_t input[ndiv8], int64_t intermed[ndiv8][8]){

    for (size_t i = 0; i < ndiv8; i++){
        auto table_start = high_resolution_clock::now();
        Util64::SIMD_Mult(NTT8_TABLE[input[i]],MULT_TABLE[i], intermed[i]);
        auto table_end = high_resolution_clock::now();
        lookup_table_time += duration_cast<nanoseconds>(table_end - table_start).count();
    }

    auto ntt_start = high_resolution_clock::now();
    Util64::SIMD_AddSub(intermed[0], intermed[1]);
    Util64::SIMD_AddSub(intermed[2], intermed[3]);
    Util64::SIMD_AddSub(intermed[4], intermed[5]);
    Util64::SIMD_AddSub(intermed[6], intermed[7]);
    Util64::SIMD_AddSub(intermed[8], intermed[9]);
    Util64::SIMD_AddSub(intermed[10], intermed[11]);
    Util64::SIMD_AddSub(intermed[12], intermed[13]);
    Util64::SIMD_AddSub(intermed[14], intermed[15]);

    Util64::SIMD_LShift(intermed[3],4);
    Util64::SIMD_LShift(intermed[7],4);
    Util64::SIMD_LShift(intermed[11],4);
    Util64::SIMD_LShift(intermed[15],4);

    Util64::SIMD_AddSub(intermed[0], intermed[2]);
    Util64::SIMD_AddSub(intermed[1], intermed[3]);
    Util64::SIMD_AddSub(intermed[4], intermed[6]);
    Util64::SIMD_AddSub(intermed[5], intermed[7]);
    Util64::SIMD_AddSub(intermed[8], intermed[10]);
    Util64::SIMD_AddSub(intermed[9], intermed[11]);
    Util64::SIMD_AddSub(intermed[12], intermed[14]);
    Util64::SIMD_AddSub(intermed[13], intermed[15]);

    Util64::SIMD_LShift(intermed[5],2);
    Util64::SIMD_LShift(intermed[6],4);
    Util64::SIMD_LShift(intermed[7],6);
    Util64::SIMD_LShift(intermed[13],2);
    Util64::SIMD_LShift(intermed[14],4);
    Util64::SIMD_LShift(intermed[15],6);

    Util64::SIMD_AddSub(intermed[0], intermed[4]);
    Util64::SIMD_AddSub(intermed[1], intermed[5]);
    Util64::SIMD_AddSub(intermed[2], intermed[6]);
    Util64::SIMD_AddSub(intermed[3], intermed[7]);
    Util64::SIMD_AddSub(intermed[8], intermed[12]);
    Util64::SIMD_AddSub(intermed[9], intermed[13]);
    Util64::SIMD_AddSub(intermed[10], intermed[14]);
    Util64::SIMD_AddSub(intermed[11], intermed[15]);

    Util64::SIMD_LShift(intermed[9],1);
    Util64::SIMD_LShift(intermed[10],2);
    Util64::SIMD_LShift(intermed[11],3);
    Util64::SIMD_LShift(intermed[12],4);
    Util64::SIMD_LShift(intermed[13],5);
    Util64::SIMD_LShift(intermed[14],6);
    Util64::SIMD_LShift(intermed[15],7);

    Util64::SIMD_AddSub(intermed[0], intermed[8]);
    Util64::SIMD_AddSub(intermed[1], intermed[9]);
    Util64::SIMD_AddSub(intermed[2], intermed[10]);
    Util64::SIMD_AddSub(intermed[3], intermed[11]);
    Util64::SIMD_AddSub(intermed[4], intermed[12]);
    Util64::SIMD_AddSub(intermed[5], intermed[13]);
    Util64::SIMD_AddSub(intermed[6], intermed[14]);
    Util64::SIMD_AddSub(intermed[7], intermed[15]);
    
    auto ntt_end = high_resolution_clock::now();
    other_ntt_time += duration_cast<nanoseconds>(ntt_end - ntt_start).count();

}


void BenchMark_MiNTT128_simd_int64_QF4::ntt_sum(uint8_t input[INPUT_SIZE_QF4], int64_t out[d][ndiv8][8]){

    for (size_t i = 0; i < m; i++){
        int64_t x[ndiv8][8];
        ncc(input+(ndiv8*i),x);
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < ndiv8; k++){
                auto key_start = high_resolution_clock::now();
                Util64::SIMD_AddMult(out[j][k],x[k],A[i][j][k]);
                auto key_end = high_resolution_clock::now();
                key_combine_time += duration_cast<nanoseconds>(key_end - key_start).count();
            }
        }
    }
}

void BenchMark_MiNTT128_simd_int64_QF4::change_base(int64_t val[d][ndiv8][8], uint8_t out[OUTPUT_SIZE_QF4]){
    for (size_t i = 0; i < d; i++){
        for (size_t j = 0; j < ndiv8; j++){
            auto mod_start = high_resolution_clock::now();
            Util64::SIMD_Mod65537(val[i][j]);
            auto mod_end = high_resolution_clock::now();
            modulo_time += duration_cast<nanoseconds>(mod_end - mod_start).count();
            
            auto base_start = high_resolution_clock::now();
            for (size_t k = 0; k < 8; k++){
                out[2*(i*n+j*8+k)] = uint8_t(val[i][j][k]);
                val[i][j][k] = val[i][j][k] >> 8;
                out[2*(i*n+j*8+k)+1] = uint8_t(val[i][j][k]);
                val[i][j][k] = val[i][j][k] >> 8;
                out[2*N+i*ndiv8+j] = out[2*N+i*ndiv8+j] || uint8_t(val[i][j][k]>>k);
            }
            auto base_end = high_resolution_clock::now();
            base_change_time += duration_cast<nanoseconds>(base_end - base_start).count();
        }
    }

}

void BenchMark_MiNTT128_simd_int64_QF4::PrintBenchMark(){

    using namespace std;

    // cout << "lookup_table_time: " << lookup_table_time << endl;
    // cout << "modulo_time: " << modulo_time << endl;
    // cout << "other_ntt_time: " << other_ntt_time << endl;
    // cout << "key_combine_time: " << key_combine_time << endl;
    // cout << "base_change_time: " << base_change_time << endl;

    float total = lookup_table_time + modulo_time + other_ntt_time+ key_combine_time + base_change_time;

    cout << "lookup_table: " << lookup_table_time/total << endl;
    cout << "modulo: " << modulo_time/total << endl;
    cout << "other_ntt: " << other_ntt_time/total << endl;
    cout << "key_combine: " << key_combine_time/total << endl;
    cout << "base_change: " << base_change_time/total << endl;

    cout << endl;


}