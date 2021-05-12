#include "test.hh"
#include <iostream>

void print() {
    std::cout<<"Test2"<<std::endl;
}

void addOne(double* a, int sz) {
    for(int i=0; i<sz; i++) a[i] = a[i]*a[i];
}