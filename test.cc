#include "test.hh"
#include <torch/torch.h>
#include <torch/script.h>
#include <torch/extension.h>
#include <iostream>

auto m = torch::jit::load("./dcgan.pth");

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

// output bs x 64 x 64 x 3
void genImage(unsigned char* output, int sz) {
    auto a = torch::randn({1, 120});
    auto ret = m.forward({a});
    auto b = (ret.toTensor().clamp(0.0, 1.0) * 256).to(torch::kUInt8);
    auto ptr = static_cast<unsigned char*>(b.data_ptr());
    for(int i=0; i<sz; i++) output[i] = ptr[i];
}

double sumVec(double* vec, int sz) {
    double ret = 0.0;
    for(int i=0; i<sz; i++) ret += vec[i];
    return ret;
}