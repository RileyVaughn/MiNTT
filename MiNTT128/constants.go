package MiNTT128

const d int = 6
const n int = 128
const N int = n * d
const m int = 108
const q int = 257
const ndiv8 int = n / 8
const Ndiv8 int = N / 8

var A [m][d * n]int

var NTT8_TABLE [256][8]int
var MULT_TABLE [16][8]int
