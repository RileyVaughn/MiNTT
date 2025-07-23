UTIL = util/util_int64.o
NORM64_64 = MiNTT64_norm_int64/MiNTT64_norm_int64.o
SIMD64_64 = MiNTT64_simd_int64/MiNTT64_simd_int64.o
NORM128_64 = MiNTT128_norm_int64/MiNTT128_norm_int64.o

INCLUDES = -I. -I./util -I./MiNTT64_norm_int64 -I./MiNTT64_simd_int64 -I./MiNTT128_norm_int64
CFLAGS = -mavx2 -mavx512f -mavx512dq -mavx512vl

all: main

main: main.cpp $(UTIL) $(NORM64_64) $(SIMD64_64) $(NORM128_64)
	g++ main.cpp $(UTIL) $(NORM64_64) $(SIMD64_64) $(NORM128_64) -o main $(INCLUDES) $(CFLAGS)

$(UTIL): util/util_int64.cpp
	g++ -c util/util_int64.cpp -o $(UTIL) $(INCLUDES) $(CFLAGS)

$(NORM64_64): MiNTT64_norm_int64/MiNTT64_norm_int64.cpp
	g++ -c MiNTT64_norm_int64/MiNTT64_norm_int64.cpp -o $(NORM64_64) $(INCLUDES) $(CFLAGS)

$(SIMD64_64): MiNTT64_simd_int64/MiNTT64_simd_int64.cpp
	g++ -c MiNTT64_simd_int64/MiNTT64_simd_int64.cpp -o $(SIMD64_64) $(INCLUDES) $(CFLAGS)

$(NORM128_64): MiNTT128_norm_int64/MiNTT128_norm_int64.cpp
	g++ -c MiNTT128_norm_int64/MiNTT128_norm_int64.cpp -o $(NORM128_64) $(INCLUDES) $(CFLAGS)


clean:
	rm -f main $(UTIL) $(NORM64_64) $(SIMD64_64) $(NORM128_64)
