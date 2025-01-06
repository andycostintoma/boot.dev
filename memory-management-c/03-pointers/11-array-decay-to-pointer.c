#include <stdio.h>

// Function to demonstrate pointer decay when an array is passed as an argument
void core_utils_func(int core_utilization[]) {
    printf("sizeof core_utilization in core_utils_func: %zu (size of pointer)\n", sizeof(core_utilization));
}

// Function to demonstrate full array and pointer to the array
void show_array_and_pointer(int arr[5]) {
    printf("Address of arr: %p\n", (void*)arr);
    printf("Address of &arr: %p (address of the whole array)\n", (void*)&arr);
}

int main() {
    // Declare and initialize an array
    int core_utilization[] = {43, 67, 89, 92, 71, 43, 56, 12};

    // Calculate length of the array using sizeof
    int len = sizeof(core_utilization) / sizeof(core_utilization[0]);

    // Print the size of the whole array
    printf("sizeof core_utilization in main: %zu (size of the whole array)\n", sizeof(core_utilization));
    printf("len of core_utilization: %d\n", len);
  
    // Call function to show sizeof within function (pointer decay)
    core_utils_func(core_utilization);

    // Declare another array to demonstrate & operator
    int another_array[5] = {1, 2, 3, 4, 5};

    // Print sizeof the whole array
    printf("sizeof another_array in main: %zu (size of the whole array)\n", sizeof(another_array));

    // Demonstrate the use of & operator on an array
    show_array_and_pointer(another_array);

    return 0;
}