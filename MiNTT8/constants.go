package MiNTT8

const d int = 96
const n int = 8
const N int = n * d
const m int = 1728
const q int = 257
const ndiv8 int = n / 8
const Ndiv8 int = N / 8

var A [m][d * n]int
var bit2ByteTable [256][8]int
var NTT8_TABLE [256][8]int
