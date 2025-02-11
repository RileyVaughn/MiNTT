package constant

const Q int = 7681
const N int = 256
const M int = 16
const D int = 3

//These constant result in 7681^256 ~ 2^3328 = 416 byte output, with a 4096 = 512 byte input
// So there is a compression, but it is not nearly the half that might be desired

//For use of NTT, n and q msut be coprime (if n=2^i). In genral, where q=p1^r1*p2^r2...pj^rj, N | gcd(p1-1,p2-1...pj-i)

// Three ways to keep output regaulr: 1) Use 2^i as q (doesn't work), 2) 2 | m*(2^i) 3) Adjust retropectivley

// Use fermat prime q=65,537, m=32, n=256 (4080^2 mod 65537 = 2)

//257
//769

// 4080 64th root
// 4938 128th root
// 59963 256th root
// 44120 512th root
// 10423 1024th root
// 57968 2048th root
// 43265 4096th root
// 56153 8192th root
// 7348 16384th root
// 927 32768th root
// 40264 65536th root
