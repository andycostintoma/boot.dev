#include <assert.h>
#include <stdio.h>
#include <math.h>

// Function to compare two floating-point numbers within a small epsilon range
int float_equals(const float a, const float b, const float epsilon) {
    return fabsf(a - b) < epsilon;
}

float snek_score(
    const int num_files,
    const int num_contributors,
    const int num_commits,
    const float avg_bug_criticality
) {
    const int project_size = num_files * num_commits;
    const int project_complexity = project_size + num_contributors;
    return (float) project_complexity * avg_bug_criticality;
}

int main() {
    const float epsilon = 1e-6;

    const float result1 = snek_score(3, 4, 5, 0.1);
    printf("result1: %f\n", result1);
    assert(float_equals(result1, 1.9, epsilon));

    const float result2 = snek_score(10, 10, 10, 0.1);
    printf("result2: %f\n", result2);
    assert(float_equals(result2, 11.0, epsilon));

    const float result3 = snek_score(105, 205, 207, 0.1);
    printf("result3: %f\n", result3);
    assert(float_equals(result3, 2194.0, epsilon));

    printf("All tests passed.\n");

    return 0;
}
