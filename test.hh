#pragma once

#ifdef __cplusplus
extern "C" {
#endif
    void print();
    void addOne(double*, int);
    void fillRandn(double*, int);
    void genImage(unsigned char* output, int sz);
#ifdef __cplusplus
}
#endif