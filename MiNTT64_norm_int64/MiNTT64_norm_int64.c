
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


void MiNTT64_norm_int64::gen8NCCMat(int64_t omega){

    for (size_t i = 0; i < 8; i++){
        for (size_t k = 0; i < 8; i++){
            
        }
        
    }
    



}


void MiNTT64_norm_int64::genMultTable(int64_t omega){


}