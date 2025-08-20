
#include <cstddef>
#include<string>
#include "MiNTT128_norm_int64_qf4.h"
#include <cstdint>
#include <iostream>
#include <fstream>

MiNTT128_norm_int64_qf4::MiNTT128_norm_int64_qf4(){
    Setup();
}

void MiNTT128_norm_int64_qf4::Setup(){

    Util64::GenNTT8Table(2,q,NTT8_TABLE);


    int64_t * mult_table = Util64::GenMultTable(27,n,q);
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


void MiNTT128_norm_int64_qf4::Hash(uint8_t input[INPUT_SIZE],uint8_t out[OUTPUT_SIZE]){
    int64_t inter[d][ndiv8][8] = {0};
    
    ntt_sum(input,inter);
    change_base(inter,out);

}


void MiNTT128_norm_int64_qf4::PrintKey(std::string filename){

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


void MiNTT128_norm_int64_qf4::ncc(uint8_t input[ndiv8], int64_t intermed[ndiv8][8]){

    for (size_t i = 0; i < ndiv8; i++){
        Util64::Norm_Mult(NTT8_TABLE[input[i]],MULT_TABLE[i], intermed[i]);
    }


    Util64::Norm_AddSub(intermed[0], intermed[1]);
    Util64::Norm_AddSub(intermed[2], intermed[3]);
    Util64::Norm_AddSub(intermed[4], intermed[5]);
    Util64::Norm_AddSub(intermed[6], intermed[7]);
    Util64::Norm_AddSub(intermed[8], intermed[9]);
    Util64::Norm_AddSub(intermed[10], intermed[11]);
    Util64::Norm_AddSub(intermed[12], intermed[13]);
    Util64::Norm_AddSub(intermed[14], intermed[15]);

    Util64::Norm_LShift(intermed[3],4);
    Util64::Norm_LShift(intermed[7],4);
    Util64::Norm_LShift(intermed[11],4);
    Util64::Norm_LShift(intermed[15],4);

    Util64::Norm_AddSub(intermed[0], intermed[2]);
    Util64::Norm_AddSub(intermed[1], intermed[3]);
    Util64::Norm_AddSub(intermed[4], intermed[6]);
    Util64::Norm_AddSub(intermed[5], intermed[7]);
    Util64::Norm_AddSub(intermed[8], intermed[10]);
    Util64::Norm_AddSub(intermed[9], intermed[11]);
    Util64::Norm_AddSub(intermed[12], intermed[14]);
    Util64::Norm_AddSub(intermed[13], intermed[15]);

    Util64::Norm_LShift(intermed[5],2);
    Util64::Norm_LShift(intermed[6],4);
    Util64::Norm_LShift(intermed[7],6);
    Util64::Norm_LShift(intermed[13],2);
    Util64::Norm_LShift(intermed[14],4);
    Util64::Norm_LShift(intermed[15],6);

    Util64::Norm_AddSub(intermed[0], intermed[4]);
    Util64::Norm_AddSub(intermed[1], intermed[5]);
    Util64::Norm_AddSub(intermed[2], intermed[6]);
    Util64::Norm_AddSub(intermed[3], intermed[7]);
    Util64::Norm_AddSub(intermed[8], intermed[12]);
    Util64::Norm_AddSub(intermed[9], intermed[13]);
    Util64::Norm_AddSub(intermed[10], intermed[14]);
    Util64::Norm_AddSub(intermed[11], intermed[15]);

    Util64::Norm_LShift(intermed[9],1);
    Util64::Norm_LShift(intermed[10],2);
    Util64::Norm_LShift(intermed[11],3);
    Util64::Norm_LShift(intermed[12],4);
    Util64::Norm_LShift(intermed[13],5);
    Util64::Norm_LShift(intermed[14],6);
    Util64::Norm_LShift(intermed[15],7);

    Util64::Norm_AddSub(intermed[0], intermed[8]);
    Util64::Norm_AddSub(intermed[1], intermed[9]);
    Util64::Norm_AddSub(intermed[2], intermed[10]);
    Util64::Norm_AddSub(intermed[3], intermed[11]);
    Util64::Norm_AddSub(intermed[4], intermed[12]);
    Util64::Norm_AddSub(intermed[5], intermed[13]);
    Util64::Norm_AddSub(intermed[6], intermed[14]);
    Util64::Norm_AddSub(intermed[7], intermed[15]);
    

}


void MiNTT128_norm_int64_qf4::ntt_sum(uint8_t input[INPUT_SIZE], int64_t out[d][ndiv8][8]){

    for (size_t i = 0; i < m; i++){
        int64_t x[ndiv8][8];
        ncc(input+(ndiv8*i),x);
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < ndiv8; k++){
                Util64::Norm_AddMult(out[j][k],x[k],A[i][j][k]);
            }
        }
    }
}

void MiNTT128_norm_int64_qf4::change_base(int64_t val[d][ndiv8][8], uint8_t out[OUTPUT_SIZE]){

    for (size_t i = 0; i < d; i++){
        for (size_t j = 0; j < ndiv8; j++){
            Util64::Norm_Mod257(val[i][j]);
            for (size_t k = 0; k < 8; k++){
                out[i*n+j*8+k] = uint8_t(val[i][j][k]);
                val[i][j][k] = val[i][j][k] >> 8;
                out[N+i*ndiv8+j] = out[N+i*ndiv8+j] || uint8_t(val[i][j][k]>>k);
            }
        }
    }

}

