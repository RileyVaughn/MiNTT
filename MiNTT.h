#ifndef MINTT_H
#define MINTT_H

#include<sizes.h>
#include <cstdint>

class MiNTT {

    public:

    virtual void Hash(uint8_t * input,uint8_t * out) = 0;

    virtual ~MiNTT() = default;

};

#endif