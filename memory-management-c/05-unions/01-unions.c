#include <stdio.h>

// Define the union age_or_name_t
typedef union AgeOrName {
    int age;
    char *name;
} age_or_name_t;

int main() {
    // Create a union instance and set the age field
    age_or_name_t lane = { .age = 29 };

    // Access and print the age field
    printf("age: %d\n", lane.age);  // prints: age: 29

    // Access and print the name field (undefined behavior)
    printf("name: %s\n", lane.name);  // prints undefined or garbage data -> SEGMENTATION FAULT

    return 0;
}