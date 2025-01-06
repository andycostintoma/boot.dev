#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <stdint.h>
#include <assert.h>

// Function to get a full greeting message
char *get_full_greeting(char *greeting, char *name, int size) {
    char *full_greeting = (char*)malloc(size * sizeof(char));
    if (full_greeting == NULL) {
        fprintf(stderr, "Memory allocation failed\n");
        exit(1);
    }
    snprintf(full_greeting, size, "%s %s", greeting, name);
    return full_greeting;
}

// Helper function to check if a pointer is on the stack
bool is_on_stack(void *ptr) {
    void *stack_top = __builtin_frame_address(0);
    uintptr_t stack_top_addr = (uintptr_t)stack_top;
    uintptr_t ptr_addr = (uintptr_t)ptr;

    // Check within a small threshold for demonstration (this value is arbitrary)
    uintptr_t threshold = 1024;

    return ptr_addr >= (stack_top_addr - threshold) && ptr_addr <= (stack_top_addr + threshold);
}

int main() {
    // Test basic greeting
    printf("Running test: test_basic_greeting\n");
    char *result = get_full_greeting("Hello", "Alice", 20);
    assert(strcmp(result, "Hello Alice") == 0); // Basic greeting should be correct
    assert(!is_on_stack(result)); // Memory should be on the heap
    free(result);
    printf("test_basic_greeting passed\n");

    // Test short buffer
    printf("Running test: test_short_buffer\n");
    result = get_full_greeting("Hey", "Bob", 4);
    assert(strcmp(result, "Hey") == 0); // Should truncate to fit buffer
    assert(!is_on_stack(result)); // Memory should be on the heap
    free(result);
    printf("test_short_buffer passed\n");

    // All tests passed
    printf("All tests passed\n");

    return 0;
}