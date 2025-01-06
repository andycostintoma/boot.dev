#include <assert.h>
#include <stdio.h>
#include <string.h>

char* get_temperature_status(const int temp) {
    if (temp < 70) {
        return "too cold";
    }
    if (temp > 90) {
        return "too hot";
    }
    return "just right";
}

int main() {
    // Test cases
    {
        char *result = get_temperature_status(50);
        printf("Test '50 should be too cold': %s\n", result);
        assert(strcmp(result, "too cold") == 0);
    }

    {
        char *result = get_temperature_status(100);
        printf("Test '100 should be too hot': %s\n", result);
        assert(strcmp(result, "too hot") == 0);
    }

    {
        char *result = get_temperature_status(70);
        printf("Test '70 should be just right': %s\n", result);
        assert(strcmp(result, "just right") == 0);
    }

    {
        char *result = get_temperature_status(75);
        printf("Test '75 should be just right': %s\n", result);
        assert(strcmp(result, "just right") == 0);
    }

    printf("All tests passed.\n");

    return 0;
}
