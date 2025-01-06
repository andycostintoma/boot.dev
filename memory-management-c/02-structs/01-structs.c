#include <assert.h>
#include <stdio.h>

struct Coordinate {
    int x;
    int y;
    int z;
};

int main() {
    // First test case
    struct Coordinate c1;
    c1.x = 1;
    c1.y = 2;
    c1.z = 3;
    assert(c1.x == 1); // should set x
    assert(c1.y == 2); // should set y
    assert(c1.z == 3); // should set z

    // Second test case
    struct Coordinate c2;
    c2.x = 4;
    c2.y = 2;
    c2.z = 0;
    assert(c2.x == 4); // should set x
    assert(c2.y == 2); // should set y
    assert(c2.z == 0); // should set z

    // Third test case
    struct Coordinate c3;
    c3.x = 10;
    c3.y = 20;
    c3.z = 30;
    assert(c3.x == 10); // should set x
    assert(c3.y == 20); // should set y
    assert(c3.z == 30); // should set z

    // Fourth test case
    struct Coordinate c4;
    c4.x = 0;
    c4.y = 6;
    c4.z = 9;
    assert(c4.x == 0); // should set x
    assert(c4.y == 6); // should set y
    assert(c4.z == 9); // should set z

    printf("All tests passed.\n");

    return 0;
}