package MiNTT8

const d int16 = 96
const n int16 = 8
const N int16 = n * d
const m int16 = 1728
const q int16 = 257
const ndiv8 int16 = n / 8
const Ndiv8 int16 = N / 8

var A [m][d][ndiv8][8]int16
var bit2ByteTable [256][8]int16
var NTT8_TABLE [256][8]int16
