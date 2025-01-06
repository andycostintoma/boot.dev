#include <stdio.h>

typedef union IntOrErrMessage {
    int data;
    char err[256];
} int_or_err_message_t;

int main() {
    // Display the size of the union
    printf("Size of int_or_err_message_t: %zu bytes\n", sizeof(int_or_err_message_t));

    // Create an array of 1000 int_or_err_message_t objects
    int_or_err_message_t messages[1000];

    // Initialize the first few elements with integer data
    for (int i = 0; i < 1000; i++) {
        messages[i].data = i + 1;
    }

    // Display the size of the entire array
    printf("Size of the array: %zu bytes\n", sizeof(messages));

    return 0;
}