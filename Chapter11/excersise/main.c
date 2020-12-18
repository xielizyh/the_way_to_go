#include "stdint.h"
#include "stdio.h"

void test1(int *p) {
    int b = 0xff;
    p = &b;
}

void test2(int **p) {
    int b = 0xff;
    *p = &b;
}

void main(void) {
    int a = 1;
    int *ptr1 = &a;
    test1(ptr1);
    printf("a=%d\n", a);
    int **ptr2 = &ptr1;
    test2(ptr2);
    printf("a=%d, *ptr1=%d, **ptr2=%d\n", a, *ptr1, **ptr2);
}
