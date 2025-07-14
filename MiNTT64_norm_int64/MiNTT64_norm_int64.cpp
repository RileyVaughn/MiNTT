
#include <cstddef>
#include <MiNTT64_norm_int64.h>
#include <cstdint>
#include <stdlib.h>



void MiNTT64_norm_int64::genKey(){

    //just for the sake of reproducibility srand(1)
    // A better implementation would usea  a "trusted randomoness" such as digits of pi
    srand(1);

    for (size_t i = 0; i < m; i++){
        for (size_t j = 0; j < d; j++){
            for (size_t k = 0; k < ndiv8; k++){
                for (size_t l = 0; l < 8; l++){
                    A[i][j][k][l] = (int64_t)(rand() % q);
                }
            } 
        } 
    }
}

void MiNTT64_norm_int64::genNTT8Table(int64_t omega){

    int64_t ncc_mat[8][8];
    Util64::Gen8NCCMat(omega,q,ncc_mat);

    for (size_t i = 0; i < 256; i++){
        for (size_t j = 0; j < 8; j++){
            int64_t vec[8];
            Util64::BitsFromByte(i, vec);
            for (size_t k = 0; k < 8; k++){
                NTT8_TABLE[i][j] = Util64::Mod_257(NTT8_TABLE[i][j] + ncc_mat[j][k]*vec[k]);
            }
        }
    }

}




void MiNTT64_norm_int64::genMultTable(int64_t omega){

    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; k < 8; k++){
            MULT_TABLE[i][k] = Util64::IntPow(omega,Util64::Bit_Rev(k,8)*(2*i+1)%(2*n),q);
        }
    }
    

}