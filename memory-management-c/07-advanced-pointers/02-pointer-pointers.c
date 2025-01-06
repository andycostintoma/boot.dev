#include <stdlib.h>
#include <assert.h>

/**
 * Allocates memory for an integer and updates the original pointer to point to the allocated memory.
 * If allocation is successful, assigns the specified value to the allocated memory.
 */

void allocate_int(int **pointer_pointer, int value) {
    // Allocating memory and updating the original pointer
    int *pointer = malloc(sizeof(int));
    *pointer_pointer = pointer;

    if (*pointer_pointer == NULL) {
        return;
    }

    // Assigning a value to the allocated memory
    **pointer_pointer = value;
}

void test_allocate() {
    int *pointer = NULL;
    allocate_int(&pointer, 10);

    assert(pointer != NULL && "Should allocate pointer");
    assert(*pointer == 10 && "Should assign value to pointer");

    free(pointer);
}

void test_does_not_overwrite() {
    int value = 5;
    int *pointer = &value;

    allocate_int(&pointer, 20);

    assert(value == 5 && "Should not overwrite original value");

    assert(pointer != NULL && "Should allocate pointer");
    assert(*pointer == 20 && "Should assign value to pointer");

    free(pointer);
}

int main() {
    test_allocate();
    test_does_not_overwrite();

    return 0;
}