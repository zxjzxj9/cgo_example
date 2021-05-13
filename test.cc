#include "test.hh"
#include <torch/extension.h>
#include <iostream>

void print() {
    std::cout<<"Test2"<<std::endl;
}

void addOne(double* a, int sz) {
    for(int i=0; i<sz; i++) a[i] = a[i]*a[i];
}

void fillRandn(double* a, int sz) {
    auto r = torch::randn({sz});
    auto ptr = static_cast<float*>(r.data_ptr());
    for(int i=0; i<sz; i++) {
        a[i] = static_cast<float>(ptr[i]);
    }
}