
#include <cstddef>
#include <MiNTT64_norm_int64.h>
#include <cstdint>


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