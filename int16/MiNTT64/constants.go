package MiNTT64

const d int16 = 12
const n int16 = 64
const N int16 = n * d
const m int16 = 216
const q int16 = 257
const ndiv8 int16 = n / 8
const Ndiv8 int16 = N / 8

var A [m][d * n]int16
var NTT8_TABLE [256][8]int16
var MULT_TABLE [8][8]int16
