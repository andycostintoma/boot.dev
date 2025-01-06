#include <stdio.h>

// Define the Coordinate struct
typedef struct Coordinate {
    int x;
    int y;
    int z;
} coordinate_t;

int main() {
    // Declare and initialize an array of 3 Coordinate structs
    coordinate_t points[3] = {
        {5, 4, 1},
        {7, 3, 2},
        {9, 6, 8}
    };

    // Cast the array of structs to an array of integers
    int *points_start = (int *)points;

    // Iterate over the known number of integers in the array of structs
    // Since we have 3 structs with 3 integers each, we iterate 9 times
    printf("Iterating over the array of integers:\n");
    for (int i = 0; i < 9; i++) {
        printf("points_start[%d] = %d\n", i, points_start[i]);
    }

    // Out-of-bounds access (intentional)
    // Accessing an element outside the bounds of the array
    // This is dangerous and can lead to undefined behavior
    printf("Out-of-bounds access:\n");
    printf("points_start[10] = %d\n", points_start[10]);

    return 0;
}