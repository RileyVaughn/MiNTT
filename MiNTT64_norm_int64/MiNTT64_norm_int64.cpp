
#include <cstddef>
#include <MiNTT64_norm_int64.h>
#include <cstdint>
#include <stdlib.h>


//Generates a random key, stored in A
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


}


int64_t * MiNTT64_norm_int64::gen8NCCMat(int64_t omega){

    int64_t ncc_mat[8][8];
    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; k < 8; k++){
            if ((k*(2*i+1))%(2*8) <= 8) {
                ncc_mat[i][k] = Util64::IntPow(omega,(k*(2*i+1))%8, q);
            } else {
                ncc_mat[i][k] = -1 * Util64::IntPow(omega, ((k*(2*i+1))%8), q);
            }
        }
    }
    
    int64_t br_ncc_mat[8][8];
    for (size_t i = 0; i < 8; i++){
        for (size_t j = 0; j < 8; j++){
        br_ncc_mat[j][Util64::Bit_Rev(i, 8)] = ncc_mat[j][i];
        }
    }


}


void MiNTT64_norm_int64::genMultTable(int64_t omega){


}