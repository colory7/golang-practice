#include <iostream>

#include "library-bridge.h"
#include "library.hpp"

void* LIB_NewFoo(int value) {
  std::cout << "[c++ bridge] LIB_NewFoo(" << value << ")" << std::endl;
  Foo* foo = new Foo(value);
  std::cout << "[c++ bridge] LIB_NewFoo(" << value << ") will return pointer "
            << foo << std::endl;
  return foo;
}

// Utility function local to the bridge's implementation

Foo* AsFoo(void* foo) { return reinterpret_cast<Foo*>(foo); }

void LIB_DestroyFoo(void* foo) {
  std::cout << "[c++ bridge] LIB_DestroyFoo(" << foo << ")" << std::endl;
  AsFoo(foo)->~Foo();
}

int LIB_FooValue(void* foo) {
  std::cout << "[c++ bridge] LIB_FooValue(" << foo << ")" << std::endl;
  return AsFoo(foo)->value();
}

const char* LIB_TestString(char* params) {
  std::cout << "[c++ bridge] LIB_TestString params is " << params << std::endl;
  Foo foo(1);
  return foo.testString(params).c_str();
  //return const_cast<char*>(foo.testString(params).c_str());
  //return new Foo()->testString(params);
}
