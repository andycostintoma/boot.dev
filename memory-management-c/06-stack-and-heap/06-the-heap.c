#include <stdio.h>
#include <stdlib.h>

// Function to create a new integer array on the heap
int *new_int_array(int size) {
    int *new_arr = malloc(size * sizeof(int)); // Allocate memory
    if (new_arr == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        exit(1); // Exit if allocation fails
    }
    return new_arr;
}

int main() {
    // Create an array of 6 integers on the heap
    int *arr_of_6 = new_int_array(6);

    // Assign values to the array elements
    arr_of_6[0] = 69;
    arr_of_6[1] = 42;
    arr_of_6[2] = 420;
    arr_of_6[3] = 1337;
    arr_of_6[4] = 7;
    arr_of_6[5] = 0;

    // Print the assigned values
    for (int i = 0; i < 6; i++) {
        printf("arr_of_6[%d] = %d\n", i, arr_of_6[i]);
    }

    // Deallocate the memory
    free(arr_of_6);

    // IMPORTANT: arr_of_6 now becomes a dangling pointer
    return 0;
}