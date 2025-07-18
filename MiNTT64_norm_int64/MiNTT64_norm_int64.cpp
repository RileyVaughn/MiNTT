
#include <cstddef>
#include "MiNTT64_norm_int64.h"
#include <cstdint>


MiNTT64_norm_int64::MiNTT64_norm_int64(){
    Setup();
}

void MiNTT64_norm_int64::Setup(){

    Util64::GenNTT8Table(2,q,NTT8_TABLE);
    Util64::GenMultTable(42,n,q,MULT_TABLE);

    int64_t * key = Util64::GenKey(m,n,d,q);

    for (size_t i = 0; i < m; i++){
       for (size_t j = 0; j < d; j++){
        for (size_t k = 0; k < n; k++){
            for (size_t l = 0; l < 8; l++){
                A[i][j][k][l] = key[i*(d*n)+j*n+k*8+l];
            }
        }
       }
    }
    delete[] key;

}


void MiNTT64_norm_int64::ncc(uint8_t input[ndiv8], int64_t intermed[ndiv8][8]){

    for (size_t i = 0; i < ndiv8; i++){
        Util64::Norm_Mult(NTT8_TABLE[input[i]],MULT_TABLE[i], intermed[i]);
    }
    
    Util64::Norm_AddSub(intermed[0], intermed[1]);
    Util64::Norm_AddSub(intermed[2], intermed[3]);
    Util64::Norm_AddSub(intermed[4], intermed[5]);
    Util64::Norm_AddSub(intermed[6], intermed[7]);

    Util64::Norm_LShift(intermed[3],4);
    Util64::Norm_LShift(intermed[3],7);

    Util64::Norm_AddSub(intermed[0], intermed[2]);
    Util64::Norm_AddSub(intermed[1], intermed[3]);
    Util64::Norm_AddSub(intermed[4], intermed[6]);
    Util64::Norm_AddSub(intermed[5], intermed[7]);

    Util64::Norm_LShift(intermed[3],2);
    Util64::Norm_LShift(intermed[3],4);
    Util64::Norm_LShift(intermed[3],6);

    Util64::Norm_AddSub(intermed[0], intermed[4]);
    Util64::Norm_AddSub(intermed[1], intermed[5]);
    Util64::Norm_AddSub(intermed[2], intermed[6]);
    Util64::Norm_AddSub(intermed[3], intermed[7]);

}