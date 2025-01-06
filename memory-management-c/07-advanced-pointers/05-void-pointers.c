#include <stdio.h>

// Enum for data types
typedef enum DATA_TYPE {
    INT,
    FLOAT,
    DOUBLE,
    CHAR
} data_type_t;

// Function to print the value of a generic pointer based on its type
void printValue(void *ptr, data_type_t type) {
    switch (type) {
        case INT:
            printf("Value: %d\n", *(int*)ptr);
        break;
        case FLOAT:
            printf("Value: %f\n", *(float*)ptr);
        break;
        default:
            printf("Unknown type!\n");
        break;
    }
}

int main() {
    int number = 42;
    float decimal = 3.14f;

    void *generic_ptr = &number;

    // This doesn't work
    // printf("Value of number: %d\n", *generic_ptr);

    // This works: Cast to appropriate type before dereferencing
    printf("Value of number: %d\n", *(int*)generic_ptr);

    // Demonstrate int
    printf("Demonstrating int:\n");
    printValue(&number, INT);

    // Demonstrate float
    printf("Demonstrating float:\n");
    printValue(&decimal, FLOAT);


    return 0;
}