#include <stdio.h>

int main() {
    // Array declaration and initialization
    int numbers[5] = {1, 2, 3, 4, 5};

    // Accessing elements using array indexing
    printf("Using array indexing:\n");
    printf("numbers[2] = %d\n", numbers[2]);  // Access the third element (index 2). Output: 3

    // Accessing elements using pointers (pointer arithmetic)
    printf("\nUsing pointers with pointer arithmetic:\n");
    printf("*(numbers + 2) = %d\n", *(numbers + 2));  // Same as numbers[2]. Output: 3

    // Pointer arithmetic
    int *ptr = numbers;  // numbers acts as a pointer to the first element
    printf("\nPointer arithmetic:\n");
    printf("Pointer ptr points to numbers[0]: %d\n", *ptr);  // Output: 1

    ptr += 2;  // Move the pointer to the third element
    printf("Pointer ptr points to numbers[2]: %d\n", *ptr);  // Output: 3

    printf("\nComparison between addresses obtained via array indexing and pointers:\n");
    for (int i = 0; i < 5; i++) {
        printf("Address of numbers[%d]: %p, Same as numbers + %d: %p\n", i, &numbers[i], i, numbers + i);
    }

    return 0;
}