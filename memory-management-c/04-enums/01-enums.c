#include <stdio.h>
#include <assert.h>

typedef enum Color {
    RED, GREEN, BLUE
} color_t;

int main() {
    assert(RED == 0);
    assert(GREEN == 1);
    assert(BLUE == 2);

    printf("All tests passed!\n");

    return 0;
}