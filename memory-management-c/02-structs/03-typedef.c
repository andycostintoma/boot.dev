#include <assert.h>

typedef struct Coordinate {
    int x;
    int y;
    int z;
} coordinate_t;


coordinate_t new_coord(int x, int y, int z) {
    coordinate_t coord = {.x = x, .y = y, .z = z};
    return coord;
}

coordinate_t scale_coordinate(coordinate_t coord, int factor) {
    coordinate_t scaled = coord;
    scaled.x *= factor;
    scaled.y *= factor;
    scaled.z *= factor;
    return scaled;
}

int main() {
    coordinate_t c1;
    c1.x = 1;
    c1.y = 2;
    c1.z = 3;
    assert(c1.x == 1);
    assert(c1.y == 2);
    assert(c1.z == 3);
}
