
#ifndef SIZES_H
#define SIZES_H

//Just a convient way to switch between input and output sizes
const int SCALE_PARAM = 1;

//Number of bytes in output for mod257 functions
const int OUTPUT_SIZE = ((128*SCALE_PARAM)*9)/8;
// Number of bytes in input for mod257 functions
const int INPUT_SIZE = OUTPUT_SIZE*2;

const int OUTPUT_SIZE_QF4 = ((128*SCALE_PARAM)*17)/8;
const int INPUT_SIZE_QF4 = OUTPUT_SIZE_QF4*2;



#endif

//Default order of how contaants are assigned.
// Define n, d, q
// N = n*d
// m must be greater than d*log_2(q), we say m=2*d*ceil(log_2(q)),
// so that the compression function compresses by a factor of 2.
// INPUT_SIZE = m*n bits or  m*n/8 bytes
// OUTPUT_SIZE = N*ceil(log_2(q)) bits or N*ceil(log_2(q))/8 bytes


// To test the many MiNNT version simultaneously with ease, we do the following:
// Define n, q
// Define INPUT_SIZE, and OUTPUT_SIZE 
// Retrieve m, d, and N from INPUT_SIZE and OUTPUT_SIZE definitions