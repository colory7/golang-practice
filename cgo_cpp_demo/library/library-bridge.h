#pragma once
#ifdef __cplusplus
extern "C" {
#endif

void* LIB_NewFoo(int value);
void LIB_DestroyFoo(void* foo);
int LIB_FooValue(void* foo);
const char* LIB_TestString(char* params);

#ifdef __cplusplus
}  // extern "C"
#endif
