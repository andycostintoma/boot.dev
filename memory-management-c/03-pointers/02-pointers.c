#include <stdio.h>

int main() {

    // In C, all data types are passed by value

    int x = 5;           // Step 1: x is initialized to 5
    int y = x;           // Step 2: y is assigned the value of x, so y becomes 5. y is a copy of the value of x.

    x = 7;               // Step 3: x is now changed to 7. This does not affect y because y is a separate copy
                         // of the initial value of x.

    int *x_ptr = &x;     // Step 4: x_ptr is a pointer that stores the address of x (x_ptr points to x)
    int z = *x_ptr;      // Step 5: dereferencing pointer, z is assigned the value pointed to by x_ptr
                         // (which is the current value of x, 7). z is another copy of the value.

    *x_ptr = 12;         // Step 6: The value of x (pointed to by x_ptr) is changed to 12

    printf("x: %d\n", x);          // x is now 12
    printf("y: %d\n", y);          // y remains 5. It was never changed after its initial copy.
    printf("x_ptr: %p\n", (void *)x_ptr);  // x_ptr prints the address of x
    printf("z: %d\n", z);          // z remains 7. It was a copy of the value of x at the time of assignment.

    return 0;
}