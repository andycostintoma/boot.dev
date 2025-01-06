#include <assert.h>
#include <stdio.h>

typedef struct CodeFile {
    int lines;
    int filetype;
} codefile_t;

/*
 * Passing by value
 * Pros: Simple to use and understand; ensures the original struct remains unchanged.
 * Cons: Involves copying the entire struct, which can be inefficient for large structs.
 * Use Case: When you need to ensure the original data is not modified and the cost of copying is acceptable.
*/
codefile_t change_filetype_by_value(codefile_t f, int new_filetype) {
    codefile_t new_f = f;  // Make a copy of the struct
    new_f.filetype = new_filetype;  // Modify the copy
    return new_f;  // Return the modified copy
}

/*
 * Passing by reference
 * Pros: Efficient memory usage; changes directly affect the original struct.
 * Cons: Requires careful handling of pointers; potential for accidental modification of the original data.
 * Use Case: When you need to modify the original data directly and efficiently.
*/
void change_filetype_by_reference(codefile_t *f, int new_filetype) {
    f->filetype = new_filetype;  // Directly modify the original struct
}

/*
 * Passing a pointer and dereferencing inside the function
 * Pros: Combines efficient memory usage with the safety of working on a copy.
 * Cons: Slightly more complex than the other methods; still involves making a copy inside the function.
 * Use Case: When you want efficient parameter passing and to avoid modifying the original data
*/
codefile_t change_filetype_by_dereferencing(codefile_t *f, int new_filetype) {
    codefile_t new_f = *f;  // Dereference pointer and make a copy
    new_f.filetype = new_filetype;  // Modify the copy
    return new_f;  // Return the modified copy
}

int main() {
    // Original struct
    codefile_t original;
    original.lines = 100;
    original.filetype = 1;

    // Test "Passing by value"
    {
        codefile_t result = change_filetype_by_value(original, 2);
        assert(result.filetype == 2);  // The new struct's filetype should be changed
        assert(result.lines == 100);   // The number of lines should remain unchanged
        assert(original.filetype == 1);  // The original struct should remain unchanged
    }

    // Test "Passing by reference"
    {
        change_filetype_by_reference(&original, 3);
        assert(original.filetype == 3);  // The original struct's filetype should change
        assert(original.lines == 100);   // The number of lines should remain unchanged
    }

    // Reset the original struct
    original.filetype = 1;

    // Test "Passing a pointer and dereferencing inside the function"
    {
        codefile_t result = change_filetype_by_dereferencing(&original, 4);
        assert(result.filetype == 4);  // The new struct's filetype should be changed
        assert(result.lines == 100);   // The number of lines should remain unchanged
        assert(original.filetype == 1);  // The original struct should remain unchanged
    }

    printf("All assertions passed.\n");
    return 0;
}