#include <stdio.h>
#include <assert.h>

typedef struct coordinate {
    int x;
    int y;
    int z;
} coordinate_t;

// This function attempts to update the x value of a coordinate
// However, since structs are passed by value, changes here do not affect the original struct
void coordinate_update_x(coordinate_t coord, int new_x) {
    coord.x = new_x;
}

// This function updates the x value of a coordinate and returns the modified struct
// Again, the changes don't affect the original struct but are reflected in the returned copy
coordinate_t coordinate_update_and_return_x(coordinate_t coord, int new_x) {
    coord.x = new_x;
    return coord;
}

coordinate_t new_coordinate(int x, int y, int z) {
    return (coordinate_t){.x = x, .y = y, .z = z};
}

int main() {
    // Test 1: coordinate_update_x should not change the original coordinate
    {
        coordinate_t old = new_coordinate(1, 2, 3);
        assert(old.x == 1);

        // Attempt to update x value using coordinate_update_x
        // Since coord is passed by value, old.x remains unchanged
        coordinate_update_x(old, 4);
        assert(old.x == 1);  // old.x must still be 1
    }

    // Test 2: coordinate_update_and_return_x should return a new coordinate with updated x
    {
        coordinate_t old = new_coordinate(1, 2, 3);
        assert(old.x == 1);

        // Update x value and get a new struct with the updated x value
        // The original struct old remains unchanged
        coordinate_t new = coordinate_update_and_return_x(old, 4);
        assert(new.x == 4);  // new.x is updated to 4
        assert(old.x == 1);  // old.x must still be 1

        // Ensure that the addresses are different, i.e., old and new are different instances
        assert(&old != &new);
    }

    printf("All assertions passed.\n");
    return 0;
}