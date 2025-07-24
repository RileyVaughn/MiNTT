UTIL64 = util/util_int64
UTIL16 = util/util_int16

NORM64_64 = MiNTT64_norm_int64/MiNTT64_norm_int64
SIMD64_64 = MiNTT64_simd_int64/MiNTT64_simd_int64
NORM128_64 = MiNTT128_norm_int64/MiNTT128_norm_int64
SIMD128_64 = MiNTT128_simd_int64/MiNTT128_simd_int64

NORM64_16 = MiNTT64_norm_int16/MiNTT64_norm_int16

INCLUDES = -I. -I./util -I./MiNTT64_norm_int64 -I./MiNTT64_simd_int64 -I./MiNTT128_norm_int64 -I./MiNTT128_simd_int64 -I./MiNTT64_norm_int16
CFLAGS = -mavx2 -mavx512f -mavx512dq -mavx512vl

all: main
 
main: main.cpp $(UTIL64).o $(UTIL16).o $(NORM64_64).o $(SIMD64_64).o $(NORM128_64).o $(SIMD128_64).o $(NORM64_16).o
	g++ main.cpp $(UTIL64).o $(UTIL16).o $(NORM64_64).o $(SIMD64_64).o $(NORM128_64).o $(SIMD128_64).o  $(NORM64_16).o -o main $(INCLUDES) $(CFLAGS)

$(UTIL64).o: $(UTIL64).cpp
	g++ -c $(UTIL64).cpp -o $(UTIL64).o $(INCLUDES) $(CFLAGS)

$(UTIL16).o: $(UTIL16).cpp
	g++ -c $(UTIL16).cpp -o $(UTIL16).o $(INCLUDES) $(CFLAGS)

$(NORM64_64).o: $(NORM64_64).cpp
	g++ -c $(NORM64_64).cpp -o $(NORM64_64).o $(INCLUDES) $(CFLAGS)

$(SIMD64_64).o: $(SIMD64_64).cpp
	g++ -c $(SIMD64_64).cpp -o $(SIMD64_64).o $(INCLUDES) $(CFLAGS)

$(NORM128_64).o: $(NORM128_64).cpp
	g++ -c $(NORM128_64).cpp -o $(NORM128_64).o $(INCLUDES) $(CFLAGS)

$(SIMD128_64).o: $(SIMD128_64).cpp
	g++ -c $(SIMD128_64).cpp -o $(SIMD128_64).o $(INCLUDES) $(CFLAGS)

$(NORM64_16).o: $(NORM64_16).cpp
	g++ -c $(NORM64_16).cpp -o $(NORM64_16).o $(INCLUDES) $(CFLAGS)

clean:
	rm -f main $(UTIL64).o $(UTIL16).o $(NORM64_64).o $(SIMD64_64).o $(NORM128_64).o $(SIMD128_64).o $(NORM64_16).o
