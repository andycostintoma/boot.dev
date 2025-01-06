#include <stdio.h>

int main() {
    const int max_recursive_calls = 100;
    const char io_mode = 'w';
    const float throttle_speed = 0.2;

    printf("Max recursive calls: %d\n", max_recursive_calls);
    printf("IO mode: %c\n", io_mode);
    printf("Throttle speed: %f\n", throttle_speed);
    return 0;
}
