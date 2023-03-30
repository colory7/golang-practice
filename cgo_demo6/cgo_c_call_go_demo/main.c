#include "main.h"

extern void (*table[1]) = {Start};

void main(){
    Start();
}