#include <stdio.h>

int main() {
    int age = 37;
    printf("The address of age is: %p\n", &age);
    printf("Size of age is: %llu\n", sizeof(age));
    printf("Size of address is: %llu\n", sizeof(&age));
    return 0;
}
