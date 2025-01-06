#include <assert.h>
#include <stdio.h>

struct Coordinate {
    int x;
    int y;
    int z;
};

struct Coordinate new_coord(int x, int y, int z) {
    struct Coordinate c = {
        .x = x,
        .y = y,
        .z = z
    };
    return c;
}

struct Coordinate scale_coordinate(struct Coordinate coord, int factor) {
    return new_coord(coord.x * factor, coord.y * factor, coord.z * factor);
}

int main() {
    // Zero Initializer
    struct Coordinate c_zero = {0};
    assert(c_zero.x == 0);
    assert(c_zero.y == 0);
    assert(c_zero.z == 0);

    // Positional Initializer
    struct Coordinate c_positional = {1, 2, 3};
    assert(c_positional.x == 1);
    assert(c_positional.y == 2);
    assert(c_positional.z == 3);

    // Designated Initializer
    struct Coordinate c_designated = {.x = 4, .y = 5, .z = 6};
    assert(c_designated.x == 4);
    assert(c_designated.y == 5);
    assert(c_designated.z == 6);

    // Using the new_coord function
    struct Coordinate c_new = new_coord(7, 8, 9);
    assert(c_new.x == 7);
    assert(c_new.y == 8);
    assert(c_new.z == 9);

    printf("All tests passed.\n");

    return 0;
}
