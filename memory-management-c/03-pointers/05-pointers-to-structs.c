#include <stdio.h>

typedef struct coordinate {
    int x;
    int y;
    int z;
} coordinate_t;

int main() {
    // Initialize a struct
    coordinate_t point = {10, 20, 30};

    // Accessing struct fields using the dot (.) operator
    printf("Using dot operator:\n");
    printf("X: %d\n", point.x);
    printf("Y: %d\n", point.y);
    printf("Z: %d\n", point.z);

    // Creating a pointer to the struct
    coordinate_t *ptrToPoint = &point;

    // Accessing struct fields using the arrow (->) operator
    printf("\nUsing arrow operator:\n");
    printf("X: %d\n", ptrToPoint->x);
    printf("Y: %d\n", ptrToPoint->y);
    printf("Z: %d\n", ptrToPoint->z);

    // Accessing struct fields using dereference and dot (* and .) operators
    printf("\nUsing dereference and dot operators:\n");
    printf("X: %d\n", (*ptrToPoint).x);
    printf("Y: %d\n", (*ptrToPoint).y);
    printf("Z: %d\n", (*ptrToPoint).z);

    return 0;
}