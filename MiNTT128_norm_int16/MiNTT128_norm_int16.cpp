
#include <cstddef>
#include<string>
#include "MiNTT128_norm_int16.h"
#include <cstdint>
#include <iostream>
#include <fstream>

MiNTT128_norm_int16::MiNTT128_norm_int16(){
    Setup();
}

void MiNTT128_norm_int16::Setup(){

    Util16::GenNTT8Table(2,q,NTT8_TABLE);

    int16_t * mult_table = Util16::GenMultTable(42,n,q);
    for (size_t i = 0; i < ndiv8; i++){
        for (size_t j = 0; j < 8; j++){
            MULT_TABLE[i][j] = mult_table[i*8+j];
        }
    }
    delete[] mult_table;

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


void MiNTT128_norm_int16::Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]){
    int16_t inter[d][ndiv8][8] = {0};
    ntt_sum(input,inter);
    change_base(inter,out);

}


void MiNTT128_norm_int16::PrintKey(std::string filename){

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


void MiNTT128_norm_int16::ncc(uint8_t input[ndiv8], int16_t intermed[ndiv8][8]){

    for (size_t i = 0; i < ndiv8; i++){
        Util16::Norm_Mult(NTT8_TABLE[input[i]],MULT_TABLE[i], intermed[i]);
        Util16::Norm_Q_reduce(intermed[i]);
    }
    
    Util16::Norm_AddSub(intermed[0], intermed[1]);
    Util16::Norm_AddSub(intermed[2], intermed[3]);
    Util16::Norm_AddSub(intermed[4], intermed[5]);
    Util16::Norm_AddSub(intermed[6], intermed[7]);
    Util16::Norm_AddSub(intermed[8], intermed[9]);
    Util16::Norm_AddSub(intermed[10], intermed[11]);
    Util16::Norm_AddSub(intermed[12], intermed[13]);
    Util16::Norm_AddSub(intermed[14], intermed[15]);

    Util16::Norm_LShift(intermed[3],4);
    Util16::Norm_LShift(intermed[7],4);
    Util16::Norm_LShift(intermed[11],4);
    Util16::Norm_LShift(intermed[15],4);

    Util16::Norm_AddSub(intermed[0], intermed[2]);
    Util16::Norm_AddSub(intermed[1], intermed[3]);
    Util16::Norm_AddSub(intermed[4], intermed[6]);
    Util16::Norm_AddSub(intermed[5], intermed[7]);
    Util16::Norm_AddSub(intermed[8], intermed[10]);
    Util16::Norm_AddSub(intermed[9], intermed[11]);
    Util16::Norm_AddSub(intermed[12], intermed[14]);
    Util16::Norm_AddSub(intermed[13], intermed[15]);

    Util16::Norm_LShift(intermed[5],2);
    Util16::Norm_LShift(intermed[6],4);
    Util16::Norm_LShift(intermed[7],6);
    Util16::Norm_LShift(intermed[13],2);
    Util16::Norm_LShift(intermed[14],4);
    Util16::Norm_LShift(intermed[15],6);

    Util16::Norm_AddSub(intermed[0], intermed[4]);
    Util16::Norm_AddSub(intermed[1], intermed[5]);
    Util16::Norm_AddSub(intermed[2], intermed[6]);
    Util16::Norm_AddSub(intermed[3], intermed[7]);
    Util16::Norm_AddSub(intermed[8], intermed[12]);
    Util16::Norm_AddSub(intermed[9], intermed[13]);
    Util16::Norm_AddSub(intermed[10], intermed[14]);
    Util16::Norm_AddSub(intermed[11], intermed[15]);

    Util16::Norm_LShift(intermed[9],1);
    Util16::Norm_LShift(intermed[10],2);
    Util16::Norm_LShift(intermed[11],3);
    Util16::Norm_LShift(intermed[12],4);
    Util16::Norm_LShift(intermed[13],5);
    Util16::Norm_LShift(intermed[14],6);
    Util16::Norm_LShift(intermed[15],7);

    Util16::Norm_AddSub(intermed[0], intermed[8]);
    Util16::Norm_AddSub(intermed[1], intermed[9]);
    Util16::Norm_AddSub(intermed[2], intermed[10]);
    Util16::Norm_AddSub(intermed[3], intermed[11]);
    Util16::Norm_AddSub(intermed[4], intermed[12]);
    Util16::Norm_AddSub(intermed[5], intermed[13]);
    Util16::Norm_AddSub(intermed[6], intermed[14]);
    Util16::Norm_AddSub(intermed[7], intermed[15]);
    

    for(size_t i=0;i < ndiv8;i++){
        Util16::Norm_Center257(intermed[i]);
    }

}


void MiNTT128_norm_int16::ntt_sum(uint8_t input[INPUT_SIZE], int16_t out[d][ndiv8][8]){

    for (size_t i = 0; i < m; i++){
        int16_t x[ndiv8][8];
        ncc(input+(ndiv8*i),x);
       
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < ndiv8; k++){
                Util16::Norm_AddMult(out[j][k],x[k],A[i][j][k]);
                Util16::Norm_Mod257(out[j][k]);
            }
        }
    }

}

void MiNTT128_norm_int16::change_base(int16_t val[d][ndiv8][8], uint8_t out[OUTPUT_SIZE]){

    for (size_t i = 0; i < d; i++){
        for (size_t j = 0; j < ndiv8; j++){
            Util16::Norm_Mod257(val[i][j]);
            for (size_t k = 0; k < 8; k++){
                out[i*n+j*8+k] = uint8_t(val[i][j][k]);
                val[i][j][k] = val[i][j][k] >> 8;
                out[N+i*ndiv8+j] = out[N+i*ndiv8+j] || uint8_t(val[i][j][k]>>k);
            }
        }
    }

}

