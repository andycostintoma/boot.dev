#include <stdio.h>
#include <stdlib.h>

int main() {
    // Example using malloc
    int *ptr_malloc = (int *)malloc(4 * sizeof(int));
    if (ptr_malloc == NULL) {
        // Handle memory allocation failure
        printf("Memory allocation with malloc failed\n");
        exit(1);
    }

    // Memory unitialized
    printf("Memory allocated using malloc and uninitialized:\n");
    for (int i = 0; i < 4; i++) {
        printf("%d ", ptr_malloc[i]);
    }
    printf("\n");


    // Memory is uninitialized - let's initialize it manually
    for (int i = 0; i < 4; i++) {
        ptr_malloc[i] = i + 1;
    }

    printf("Memory allocated using malloc and initialized manually:\n");
    for (int i = 0; i < 4; i++) {
        printf("%d ", ptr_malloc[i]);
    }
    printf("\n");

    // Example using calloc
    int *ptr_calloc = (int *)calloc(4, sizeof(int));
    if (ptr_calloc == NULL) {
        // Handle memory allocation failure
        printf("Memory allocation with calloc failed\n");
        free(ptr_malloc); // free previously allocated memory before exiting
        exit(1);
    }

    // Memory is already initialized to zero by calloc
    printf("Memory allocated using calloc:\n");
    for (int i = 0; i < 4; i++) {
        printf("%d ", ptr_calloc[i]);
    }
    printf("\n");

    // Free the allocated memory to avoid memory leaks
    free(ptr_malloc);
    free(ptr_calloc);

    printf("Memory has been freed.\n");

    return 0;
}