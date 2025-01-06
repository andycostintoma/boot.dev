#include <stdio.h>
#include <stdlib.h>

int main() {
    // Allocating memory with malloc
    int *ptr = (int *)malloc(5 * sizeof(int));
    if (ptr == NULL) {
        printf("Memory allocation failed\n");
        exit(1);
    }

    // Using the allocated memory
    for (int i = 0; i < 5; i++) {
        ptr[i] = i * 10;
    }

    // Display the allocated memory contents
    printf("Allocated memory contents: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");

    // Freeing the allocated memory
    free(ptr);

    // Setting the pointer to NULL to avoid undefined behavior
    ptr = NULL;

    // Trying to access the freed memory (uncommenting the below lines will cause undefined behavior)
    /*
    printf("Trying to access freed memory: ");
    for (int i = 0; i < 5; i++) {
        printf("%d ", ptr[i]);
    }
    printf("\n");
    */

    printf("Memory has been freed and pointer set to NULL.\n");

    return 0;
}