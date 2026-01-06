

SRCS := util/util_int64.cpp \
		util/util_int16.cpp \
		MiNTT64_norm_int64/MiNTT64_norm_int64.cpp \
		MiNTT64_simd_int64/MiNTT64_simd_int64.cpp \
		MiNTT128_norm_int64/MiNTT128_norm_int64.cpp \
		MiNTT128_simd_int64/MiNTT128_simd_int64.cpp \
		MiNTT64_norm_int16/MiNTT64_norm_int16.cpp \
		MiNTT64_simd_int16/MiNTT64_simd_int16.cpp \
		MiNTT128_norm_int16/MiNTT128_norm_int16.cpp \
		MiNTT128_simd_int16/MiNTT128_simd_int16.cpp \
		MiNTT8_norm_int16/MiNTT8_norm_int16.cpp \
		MiNTT8_simd_int16/MiNTT8_simd_int16.cpp \
		MiNTT128_norm_int64_QF4/MiNTT128_norm_int64_QF4.cpp \
		MiNTT128_simd_int64_QF4/MiNTT128_simd_int64_QF4.cpp \
		MiNTT64_norm_int64_QF4/MiNTT64_norm_int64_QF4.cpp \
		MiNTT64_simd_int64_QF4/MiNTT64_simd_int64_QF4.cpp 

OBJS := $(SRCS:.cpp=.o)


INCLUDES = -I. -I./util -I./MiNTT64_norm_int64 -I./MiNTT64_simd_int64 -I./MiNTT128_norm_int64 -I./MiNTT128_simd_int64 -I./MiNTT64_norm_int16 -I./MiNTT64_simd_int16 -I./MiNTT128_norm_int16 -I./MiNTT128_simd_int16 -I./MiNTT8_norm_int16 -I./MiNTT8_simd_int16 -I./MiNTT128_norm_int64_QF4 -I./MiNTT128_simd_int64_QF4 -I./MiNTT64_norm_int64_QF4 -I./MiNTT64_simd_int64_QF4
CFLAGS = -mavx2 -mavx512f -mavx512dq -mavx512vl 
#-fstack-usage
# Linker flag assumes crypto++ is installed system wide
LFLAGS = -lcryptopp

all: main
 
main: main.cpp $(OBJS)
	g++ main.cpp $(OBJS) -o main $(INCLUDES) $(CFLAGS) $(LFLAGS)


%.o: %.cpp
	g++ -c $< -o $@ $(INCLUDES) $(CFLAGS)




clean:
	rm -f main $(OBJS)
