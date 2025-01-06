#include <stdio.h>
#include <string.h>

int main() {
    // strcpy: Copies a string to another.
    char src[] = "Hello";
    char dest[6];
    strcpy(dest, src); // dest now contains "Hello"
    printf("strcpy: %s\n", dest);

    // strcat: Concatenates (appends) one string to another.
    char dest2[12] = "Hello";
    char src2[] = " World";
    strcat(dest2, src2); // dest now contains "Hello World"
    printf("strcat: %s\n", dest2);

    // strlen: Returns the length of a string (excluding the null terminator).
    char str[] = "Hello";
    size_t len = strlen(str); // len is 5
    printf("strlen: %zu\n", len);

    // strcmp: Compares two strings lexicographically.
    char str1[] = "Hello";
    char str2[] = "World";
    int result = strcmp(str1, str2); // result is negative since "Hello" < "World"
    printf("strcmp: %d (negative if str1 < str2, 0 if equal, positive if str1 > str2)\n", result);

    // strncpy: Copies a specified number of characters from one string to another.
    char src3[] = "Hello";
    char dest3[6];
    strncpy(dest3, src3, 3); // dest now contains "Hel"
    dest3[3] = '\0'; // ensure null termination
    printf("strncpy: %s\n", dest3);

    // strncat: Concatenates a specified number of characters from one string to another.
    char dest4[12] = "Hello";
    char src4[] = " World";
    strncat(dest4, src4, 3); // dest now contains "Hello Wo"
    printf("strncat: %s\n", dest4);

    // strchr: Finds the first occurrence of a character in a string.
    char str3[] = "Hello";
    char *pos = strchr(str3, 'l'); // pos points to the first 'l' in "Hello"
    if (pos != NULL) {
        printf("strchr: Found 'l' at position %lld\n", pos - str3);
    } else {
        printf("strchr: 'l' not found\n");
    }

    // strstr: Finds the first occurrence of a substring in a string.
    char str4[] = "Hello World";
    char *pos2 = strstr(str4, "World"); // pos points to "World" in "Hello World"
    if (pos2 != NULL) {
        printf("strstr: Found 'World' at position %lld\n", pos2 - str4);
    } else {
        printf("strstr: 'World' not found\n");
    }

    return 0;
}