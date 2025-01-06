#include <stdio.h>
#include <stdlib.h>

int main() {
    // Creating an array of integers on the heap
    int *int_array = malloc(sizeof(int) * 3);
    if (int_array == NULL) {
        fprintf(stderr, "Failed to allocate memory for int_array\n");
        return EXIT_FAILURE;
    }

    int_array[0] = 1;
    int_array[1] = 2;
    int_array[2] = 3;

    printf("Integer array:\n");
    for (int i = 0; i < 3; i++) {
        printf("int_array[%d] = %d\n", i, int_array[i]);
    }

    // Free the allocated memory for the integer array
    free(int_array);

    // Creating an array of strings on the heap
    char **string_array = malloc(sizeof(char *) * 3);
    if (string_array == NULL) {
        fprintf(stderr, "Failed to allocate memory for string_array\n");
        return EXIT_FAILURE;
    }

    string_array[0] = "foo";
    string_array[1] = "bar";
    string_array[2] = "baz";

    printf("\nString array:\n");
    for (int i = 0; i < 3; i++) {
        printf("string_array[%d] = %s\n", i, string_array[i]);
    }

    // Free the allocated memory for the string array
    // Note that we do not need to free each string since
    // they are string literals and not dynamically allocated.
    free(string_array);

    return EXIT_SUCCESS;
}