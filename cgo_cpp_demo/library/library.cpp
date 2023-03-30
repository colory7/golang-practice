#include "library.hpp"
#include <iostream>

Foo::Foo(int value) : m_value(value) {
  std::cout << "[c++] Foo::Foo(" << m_value << ")" << std::endl;
}

Foo::~Foo() { std::cout << "[c++] Foo::~Foo(" << m_value << ")" << std::endl; }

int Foo::value() const {
  std::cout << "[c++] Foo::value() is " << m_value << std::endl;
  return m_value;
}

string Foo::testString(string params) {
  std::cout << "[c++] testString params is " << params << std::endl;
  return params + " world!";
}
