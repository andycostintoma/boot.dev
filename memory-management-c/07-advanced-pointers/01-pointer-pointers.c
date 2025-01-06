#include <stdio.h>

int main() {
    int v1 = 1;
    int v2 = 2;
    int checkpoint = 0;

    int *ptr = &v1;
    int **ptr_ptr = &ptr;
    checkpoint = **ptr_ptr;

    printf("v1: %d\n", v1); // Prints 1
    printf("v2: %d\n", v2); // Prints 2
    printf("checkpoint: %d\n", checkpoint); // Prints 1
    printf("ptr: %p\n", (void *)ptr); // Prints address of v1
    printf("ptr_ptr: %p\n", (void *)ptr_ptr); // Prints address of ptr


    ptr = &v2;
    checkpoint = **ptr_ptr;

    printf("v1: %d\n", v1); // Prints 1
    printf("v2: %d\n", v2); // Prints 2
    printf("checkpoint: %d\n", checkpoint); // Prints 2
    printf("ptr: %p\n", (void *)ptr); // Prints address of v2
    printf("ptr_ptr: %p\n", (void *)ptr_ptr); // Prints address of ptr -> same address

    return 0;
}