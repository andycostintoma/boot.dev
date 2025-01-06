#include <stdio.h>
#include <assert.h>

typedef enum {
    EXIT_SUCCESS = 0,
    EXIT_FAILURE = 1,
    EXIT_COMMAND_NOT_FOUND = 127,
} ExitStatus;

// Define the first one and let the compiler fill in the rest
typedef enum {
    LANE_WPM = 200,
    PRIME_WPM, // 201
    TEEJ_WPM,  // 202
} WordsPerMinute;

int main() {
    assert(EXIT_SUCCESS == 0);
    assert(EXIT_FAILURE == 1);
    assert(EXIT_COMMAND_NOT_FOUND == 127);

    assert(LANE_WPM == 200);
    assert(PRIME_WPM == 201);
    assert(TEEJ_WPM == 202);

    printf("All tests passed!\n");

    return 0;
}