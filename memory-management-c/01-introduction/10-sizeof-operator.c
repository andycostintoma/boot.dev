#include <stdbool.h>
#include <stdio.h>

int main() {
    printf("sizeof(char)   = %zu\n", sizeof(char));
    printf("sizeof(bool)   = %zu\n", sizeof(bool));
    printf("sizeof(int)    = %zu\n", sizeof(int));
    printf("sizeof(float)  = %zu\n", sizeof(float));
    printf("sizeof(double) = %zu\n", sizeof(double));
    printf("sizeof(size_t) = %zu\n", sizeof(size_t));

    return 0;
}