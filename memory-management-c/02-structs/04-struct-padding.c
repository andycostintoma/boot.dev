#include <assert.h>
#include <stddef.h>

typedef struct SneklangVar {
    char *name;         // Pointer, 8 bytes on a 64-bit system
    int scope_level;    // Integer, 4 bytes
    int value;          // Integer, 4 bytes
    char type;          // Char, 1 byte
    char is_constant;   // Char, 1 byte
} sneklang_var_t;

int main() {
    // Assert the offsets of the members relative to the structure start
    assert(offsetof(sneklang_var_t, name) == 0);
    assert(offsetof(sneklang_var_t, scope_level) == 8);
    assert(offsetof(sneklang_var_t, value) == 12);
    assert(offsetof(sneklang_var_t, type) == 16);
    assert(offsetof(sneklang_var_t, is_constant) == 17);

    // To align the overall structure to 8 bytes, the compiler adds 6 bytes of padding after is_constant,
    // making the total size 24 bytes.
    assert(sizeof(sneklang_var_t) == 24);

    return 0;
}