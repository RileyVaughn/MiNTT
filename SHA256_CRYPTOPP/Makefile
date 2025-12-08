
CXX := g++
CXXFLAGS := -O3 -std=c++17 -SHA256_CRYPTOPP

SRCS := main.cpp \
        cryptopp890/algparam.cpp \
        cryptopp890/cryptlib.cpp \
        cryptopp890/filters.cpp \
        cryptopp890/hex.cpp \
        cryptopp890/integer.cpp \
        cryptopp890/misc.cpp \
        cryptopp890/secblock.cpp \
        cryptopp890/sha.cpp \
        cryptopp890/simple.cpp \
        cryptopp890/smartptr.cpp

# Object files
OBJS := $(SRCS:.cpp=.o)

# ------------------------------
# Build rules
# ------------------------------

all: $(OBJS)

# Compile each .cpp to .o
%.o: %.cpp
	$(CXX) $(CXXFLAGS) -c $< -o $@

# Clean
clean:
	rm -f $(OBJS) 
