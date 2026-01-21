
#include <cstddef>
#include<string>
#include "SWIFFT.h"
#include <cstdint>
#include <iostream>
#include <fstream>

SWIFFT::SWIFFT(){
    Setup();
}

void SWIFFT::Setup(){

    Util16::GenNTT8Table(Util16::IntPow(omega,8,q),q,NTT8_TABLE);


    int16_t * mult_table = Util16::GenMultTable(omega,n,q);
    for (size_t i = 0; i < ndiv8; i++){
        for (size_t j = 0; j < 8; j++){
            MULT_TABLE[i][j] = mult_table[i*8+j];
        }
    } 
    delete[] mult_table;
    int16_t d = 1;
    int16_t * key = Util16::GenKey(m,n,1,q);
    for (size_t i = 0; i < m; i++){
       for (size_t j = 0; j < d; j++){
        for (size_t k = 0; k < ndiv8; k++){
            for (size_t l = 0; l < 8; l++){
                A[i][k][l] = key[i*(d*n)+j*n+k*8+l];
            }
        }
       }
    }
    delete[] key;
    
}   


void SWIFFT::Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]){
    int16_t inter[ndiv8][8] = {0};
    ntt_sum(input,inter);
    change_base(inter,out);

}



void SWIFFT::ncc(uint8_t input[ndiv8], int16_t intermed[ndiv8][8]){

    for (size_t i = 0; i < ndiv8; i++){
        Util16::SIMD_Mult(NTT8_TABLE[input[i]],MULT_TABLE[i], intermed[i]);
        Util16::SIMD_Q_reduce(intermed[i]);
    }
    
    Util16::SIMD_AddSub(intermed[0], intermed[1]);
    Util16::SIMD_AddSub(intermed[2], intermed[3]);
    Util16::SIMD_AddSub(intermed[4], intermed[5]);
    Util16::SIMD_AddSub(intermed[6], intermed[7]);

    Util16::SIMD_LShift(intermed[3],4);
    Util16::SIMD_LShift(intermed[7],4);

    Util16::SIMD_Q_reduce(intermed[3]);
    Util16::SIMD_Q_reduce(intermed[7]);

    Util16::SIMD_AddSub(intermed[0], intermed[2]);
    Util16::SIMD_AddSub(intermed[1], intermed[3]);
    Util16::SIMD_AddSub(intermed[4], intermed[6]);
    Util16::SIMD_AddSub(intermed[5], intermed[7]);

    Util16::SIMD_LShift(intermed[5],2);
    Util16::SIMD_LShift(intermed[6],4);
    Util16::SIMD_LShift(intermed[7],6);

    Util16::SIMD_Q_reduce(intermed[5]);
    Util16::SIMD_Q_reduce(intermed[6]);
    Util16::SIMD_Q_reduce(intermed[7]);

    Util16::SIMD_AddSub(intermed[0], intermed[4]);
    Util16::SIMD_AddSub(intermed[1], intermed[5]);
    Util16::SIMD_AddSub(intermed[2], intermed[6]);
    Util16::SIMD_AddSub(intermed[3], intermed[7]);

    for(size_t i=0;i < ndiv8;i++){
        Util16::SIMD_Center257(intermed[i]);
    }

}


void SWIFFT::ntt_sum(uint8_t input[INPUT_SIZE], int16_t out[ndiv8][8]){

    for (size_t i = 0; i < m; i++){
        int16_t x[ndiv8][8];
        ncc(input+(ndiv8*i),x);
            for (size_t k = 0; k < ndiv8; k++){
                Util16::SIMD_AddMult(out[k],x[k],A[i][k]);
                Util16::SIMD_Q_reduce(out[k]);
            }
    }

}

void SWIFFT::change_base(int16_t val[ndiv8][8], uint8_t out[OUTPUT_SIZE]){

   
    for (size_t j = 0; j < ndiv8; j++){
        Util16::SIMD_Mod257(val[j]);
        for (size_t k = 0; k < 8; k++){
            out[j*8+k] = uint8_t(val[j][k]);
            val[j][k] = val[j][k] >> 8;
            out[n+j] = out[n+j] || uint8_t(val[j][k]>>k);
        }
    }

}

