#include <stdio.h>
#include <stdint.h>

typedef union Color {
    struct {
        uint8_t r;
        uint8_t g;
        uint8_t b;
        uint8_t a;
    } components;
    uint32_t rgba;
} color_t;

void print_color_as_components(color_t color) {
    printf("Color components - R: %u, G: %u, B: %u, A: %u\n",
           color.components.r,
           color.components.g,
           color.components.b,
           color.components.a);
}

void print_color_as_rgba(color_t color) {
    printf("Color as RGBA value: 0x%08X\n", color.rgba);
}

int main() {
    color_t color;

    // Setting individual components
    color.components.r = 255;
    color.components.g = 165;
    color.components.b = 0;
    color.components.a = 255;

    // Print the color through components
    print_color_as_components(color);

    // Print the color as a single RGBA value
    print_color_as_rgba(color);

    // Set the RGBA value directly
    color.rgba = 0x6495EDFF;

    // Print again to see the updated values
    print_color_as_components(color);
    print_color_as_rgba(color);

    return 0;
}