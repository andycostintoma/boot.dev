#include <stdio.h>
#include <string.h>
#include <assert.h>

void concat_strings(char *str1, const char *str2) {
    char *end = str1;
    while (*end != '\0') {
        end++;
    }

    while (*str2 != '\0') {
        *end = *str2;
        end++;
        str2++;
    }

    *end = '\0';
}

int main() {
    {
        char str1[100] = "";
        const char *str2 = "";
        concat_strings(str1, str2);
        assert(strcmp(str1, "") == 0);
    } {
        char str1[100] = "Hello";
        const char *str2 = "";
        concat_strings(str1, str2);
        assert(strcmp(str1, "Hello") == 0);
    } {
        char str1[100] = "Hello ";
        const char *str2 = "World";
        concat_strings(str1, str2);
        assert(strcmp(str1, "Hello World") == 0);
    } {
        char str1[200] = "This is a longer string that ";
        const char *str2 = "will be concatenated with another long string.";
        concat_strings(str1, str2);
        assert(strcmp(str1, "This is a longer string that will be concatenated with another long string.") == 0);
    } {
        char str1[100] = "";
        const char *str2 = "Hello";
        concat_strings(str1, str2);
        assert(strcmp(str1, "Hello") == 0 && "Concatenating a string to an empty should change the original string");
    }

    printf("All tests passed!\n");
    return 0;
}
