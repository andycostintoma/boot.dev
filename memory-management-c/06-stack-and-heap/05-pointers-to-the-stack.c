#include <stdio.h>

typedef struct {
    int x;
    int y;
} coord_t;

coord_t *new_coord(int x, int y) {
    coord_t c;
    c.x = x;
    c.y = y;
    return &c;
    // WARNING: Returning address of a stack-allocated variable
    // is problematic since its lifetime ends when the function returns
}

int main() {
    coord_t *c1 = new_coord(10, 20);
    coord_t *c2 = new_coord(30, 40);
    coord_t *c3 = new_coord(50, 60);

    printf("c1: %d, %d\n", c1->x, c1->y);
    printf("c2: %d, %d\n", c2->x, c2->y);
    printf("c3: %d, %d\n", c3->x, c3->y);

    // OUTPUT:
    // You will likely see unexpected values here.
    // This is because c1, c2, and c3 point to memory that is no longer valid.

    // Explanation:
    // The new_coord function creates a coord_t variable on the stack.
    // When the function returns, the stack frame is deallocated, and
    // the memory occupied by the stack frame may be overwritten by subsequent function calls.
    // Therefore, using pointers to this deallocated stack memory leads to undefined behavior.
}
