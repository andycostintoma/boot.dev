#include <stdio.h>

// Define the Coordinate struct
typedef struct Coordinate {
    int x;
    int y;
    int z;
} coordinate_t;

int main() {
    // Declare an array of 3 Coordinate structs and initialize it
    coordinate_t points[3] = {
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9}
    };

    // Access and print values of the second element in the array using array indexing
    printf("Using array indexing:\n");
    printf("points[1].x = %d, points[1].y = %d, points[1].z = %d\n",
        points[1].x, points[1].y, points[1].z);  // Output: 4, 5, 6

    // Create a pointer to the array of Coordinate structs
    coordinate_t *ptr = points;

    // Access and print values of the second element in the array using pointer arithmetic
    printf("\nUsing pointer arithmetic:\n");
    printf("ptr[1].x = %d, ptr[1].y = %d, ptr[1].z = %d\n",
        (ptr + 1)->x, (ptr + 1)->y, (ptr + 1)->z);  // Output: 4, 5, 6


    // Accessing specific elements to understand memory layout
    printf("\nAccessing elements using pointers:\n");
    printf("points + 0 points to address: %p\n", (void*)(points + 0));
    printf("&points[0] points to address: %p\n", (void*)&points[0]);

    printf("points + 1 points to address: %p\n", (void*)(points + 1));
    printf("&points[1] points to address: %p\n", (void*)&points[1]);

    printf("points + 2 points to address: %p\n", (void*)(points + 2));
    printf("&points[2] points to address: %p\n", (void*)&points[2]);
    return 0;
}