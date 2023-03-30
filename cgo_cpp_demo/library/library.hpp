#pragma once
#include <string>
using namespace std;
class Foo {
 public:
  Foo(int value);
  ~Foo();
  int value() const;
  string testString(string params);
 private:
  int m_value;
};
