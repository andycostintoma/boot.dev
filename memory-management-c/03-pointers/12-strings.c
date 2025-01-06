#include <stdio.h>
#include <string.h>


int main() {
    // Declaring array and pointer string
    char str1[] = "Hi";
    char *str2 = "Snek";

    printf("%s %s\n", str1, str2); // Output: Hi Snek

    // Using strcat to concatenate strings
    char first[50] = "Snek";
    char *second = "lang!";

    // Concatenating second to first
    // It is mandatory that first has enough space to hold the result
    strcat(first, second);

    // Print the result after concatenation
    printf("Hello, %s\n", first); // Output: Hello, Sneklang!

    // Now demonstrating the `strlen` function
    size_t len_first = strlen(first);
    printf("Sizeof first %llu\n", sizeof(first)); // Output: 50
    printf("Length of concatenated string: %zu\n", len_first); // Output: 9

    return 0;
}
