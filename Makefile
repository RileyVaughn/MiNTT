all: main

main: main.cpp MiNTT64_norm_int64/MiNTT64_norm_int64.o util/util_int64.o
	g++ main.cpp MiNTT64_norm_int64/MiNTT64_norm_int64.o util/util_int64.o -o main -I./MiNTT64_norm_int64 -I./util

MiNTT64_norm_int64/MiNTT64_norm_int64.o: MiNTT64_norm_int64/MiNTT64_norm_int64.cpp
	g++ -c MiNTT64_norm_int64/MiNTT64_norm_int64.cpp -o MiNTT64_norm_int64/MiNTT64_norm_int64.o -I./MiNTT64_norm_int64

util/util_int64.o: util/util_int64.cpp
	g++ -c util/util_int64.cpp -o util/util_int64.o -I./util

clean:
	rm -f main MiNTT64_norm_int64/*.o util/*.o