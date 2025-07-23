UTIL = util/util_int64.o
NORM = MiNTT64_norm_int64/MiNTT64_norm_int64.o
SIMD = MiNTT64_simd_int64/MiNTT64_simd_int64.o

INCLUDES = -I./util -I./MiNTT64_norm_int64 -I./MiNTT64_simd_int64
CFLAGS = -mavx2 -mavx512f -mavx512dq -mavx512vl

all: main

main: main.cpp $(UTIL) $(NORM) $(SIMD)
	g++ main.cpp $(UTIL) $(NORM) $(SIMD) -o main $(INCLUDES) $(CFLAGS)

$(UTIL): util/util_int64.cpp
	g++ -c util/util_int64.cpp -o $(UTIL) -I./util $(CFLAGS)

$(NORM): MiNTT64_norm_int64/MiNTT64_norm_int64.cpp
	g++ -c MiNTT64_norm_int64/MiNTT64_norm_int64.cpp -o $(NORM) -I./MiNTT64_norm_int64 $(CFLAGS)

$(SIMD): MiNTT64_simd_int64/MiNTT64_simd_int64.cpp
	g++ -c MiNTT64_simd_int64/MiNTT64_simd_int64.cpp -o $(SIMD) -I./MiNTT64_simd_int64 -I./util $(CFLAGS)

clean:
	rm -f main $(UTIL) $(NORM) $(SIMD)
