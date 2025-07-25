
#ifndef SIZES_H
#define SIZES_H

//Number of bytes in output
const int OUTPUT_SIZE = 864;


// Number of bytes in input
const int INPUT_SIZE = OUTPUT_SIZE*2;

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