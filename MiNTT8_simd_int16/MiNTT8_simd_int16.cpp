
#include <cstddef>
#include<string>
#include "MiNTT8_simd_int16.h"
#include <cstdint>
#include <iostream>
#include <fstream>

MiNTT8_SIMD_int16::MiNTT8_SIMD_int16(){
    Setup();
}

void MiNTT8_SIMD_int16::Setup(){

    Util16::GenNTT8Table(2,q,NTT8_TABLE);

    int16_t * key = Util16::GenKey(m,n,d,q);
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


void MiNTT8_SIMD_int16::Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]){
    int16_t inter[d][8] = {0};
    ntt_sum(input,inter);
    change_base(inter,out);

}


void MiNTT8_SIMD_int16::PrintKey(std::string filename){

    std::ofstream file(filename);
    if (!file.is_open()) {
        std::cerr << "Failed to open file: " << filename << "\n";
        return;
    }

    for (size_t i = 0; i < m; i++){
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < ndiv8; k++){
                for (size_t l = 0; l < 8; l++){
                    if(j != d-1 || k != ndiv8-1 ||l !=7){
                        file << A[i][j][k][l] << ",";
                    } else {
                        file << A[i][j][k][l] << "\n";
                    }
                }
            }
        }
    }
    
    file.close();

}



void MiNTT8_SIMD_int16::ntt_sum(uint8_t input[INPUT_SIZE], int16_t out[d][8]){

    for (size_t i = 0; i < m; i++){
        for (size_t j = 0; j < d; j++){
                Util16::SIMD_AddMult(out[j],NTT8_TABLE[input[i]],A[i][j][0]);
                Util16::SIMD_Q_reduce(out[j]);
        }
    }
}

void MiNTT8_SIMD_int16::change_base(int16_t val[d][8], uint8_t out[OUTPUT_SIZE]){

    for (size_t i = 0; i < d; i++){
            Util16::SIMD_Mod257(val[i]);
            for (size_t k = 0; k < 8; k++){
                out[i*n+k] = uint8_t(val[i][k]);
                val[i][k] = val[i][k] >> 8;
                out[N+i] = out[N+i] || uint8_t(val[i][k]>>k);
            }
    }
}

