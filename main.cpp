#include <iostream>
#include "MiNTT64_norm_int64/MiNTT64_norm_int64.h"
#include "util/util_int64.h"

using namespace std;

void print_2darray(int64_t arr[8][8]);
void GenInput(uint8_t input[INPUT_SIZE]);

int main() {

    MiNTT64_norm_int64 hash = MiNTT64_norm_int64();

    uint8_t input[INPUT_SIZE];
    uint8_t output[OUTPUT_SIZE] = {0};
    GenInput(input);

    hash.Hash(input,output);

    PrintOut(output);



    return 0;
}


void GenInput(uint8_t input[INPUT_SIZE]){

    for (size_t i = 0; i < INPUT_SIZE; i++)
    {
    input[i] = i % 256;
    }

}


void PrintOut(uint8_t output[OUTPUT_SIZE]){

    for (size_t i = 0; i < OUTPUT_SIZE; i++)
    {
        cout << int(output[i]) << " ";
    }
    cout << endl;

}