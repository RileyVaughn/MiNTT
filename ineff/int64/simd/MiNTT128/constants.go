package MiNTT128

const d int64 = 6
const n int64 = 128
const N int64 = n * d
const m int64 = 108
const q int64 = 257
const ndiv8 int64 = n / 8
const Ndiv8 int64 = N / 8
const OUT_SIZE int64 = N / 8 * 9 // assumes q=257

var A [m][d][ndiv8][8]int64

var NTT8_TABLE [256][8]int64
var MULT_TABLE [16][8]int64
