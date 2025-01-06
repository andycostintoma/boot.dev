#include <stdio.h>

int main() {
    // Initialize an array with 5 integers
    int numbers[5] = {1, 2, 3, 4, 5};

    // Print the original array elements
    printf("Original array: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");

    // Update values in the array
    numbers[1] = 20;
    numbers[3] = 40;

    // Print the updated array elements
    printf("Updated array: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", numbers[i]);
    }
    printf("\n");

    return 0;
}