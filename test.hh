#pragma once
#include <math.h>

#ifdef __cplusplus
extern "C" {
#endif
    void print();
    void addOne(double*, int);
    void fillRandn(double*, int);
    void genImage(unsigned char* output, int sz);
    double sumVec(double*, int);
#ifdef __cplusplus
}
#endif