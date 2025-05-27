package MiNTT8

const d int64 = 96
const n int64 = 8
const N int64 = n * d
const m int64 = 1728
const q int64 = 257
const ndiv8 int64 = n / 8
const Ndiv8 int64 = N / 8

var A [m][d * n]int64
var bit2ByteTable [256][8]int64
var NTT8_TABLE [256][8]int64
