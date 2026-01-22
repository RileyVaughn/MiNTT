#ifndef MINTT_H
#define MINTT_H

#include<sizes.h>
#include <cstdint>

class MiNTT {

    public:

    virtual void Hash(uint8_t * input,uint8_t * out) = 0;

    virtual ~MiNTT() = default;

    virtual void BenchMark(float & lookup_table_ratio, float & modulo_ratio, float & other_ntt_ratio, float & key_combine_ratio, float & base_change_ratio) {

        return;
    }

};

#endif