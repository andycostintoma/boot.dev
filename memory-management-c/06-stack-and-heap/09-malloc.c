#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

int *allocate_scalar_list(int size, int multiplier) {
    int *ptr = malloc(size * sizeof(int));
    if (ptr == NULL){
        return NULL;
    }
    for (int i = 0; i < size; i++) {
        ptr[i] = i * multiplier;
    }
    return ptr;
}

void test_allocate_scalar_list_size() {
    int size = 5;
    int multiplier = 2;
    int *result = allocate_scalar_list(size, multiplier);
    assert(result != NULL && "Function should return a non-null pointer");
    free(result);
}

void test_allocate_scalar_list_values() {
    int size = 5;
    int multiplier = 2;
    int *result = allocate_scalar_list(size, multiplier);
    int expected[5] = {0, 2, 4, 6, 8};
    for (int i = 0; i < size; i++) {
        assert(result[i] == expected[i] && "Element does not match expected value");
    }
    free(result);
}

void test_allocate_scalar_list_zero_multiplier() {
    int size = 3;
    int multiplier = 0;
    int *result = allocate_scalar_list(size, multiplier);
    for (int i = 0; i < size; i++) {
        assert(result[i] == 0 && "All elements should be 0 with multiplier 0");
    }
    free(result);
}


int main() {
    test_allocate_scalar_list_size();
    test_allocate_scalar_list_values();
    test_allocate_scalar_list_zero_multiplier();

    printf("All tests passed.\n");
    return 0;
}