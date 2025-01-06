#include <assert.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>

void swap(void *vp1, void *vp2, size_t size) {
    void *tmp = malloc(size);
    if (tmp == NULL) {
        return;
    }
    memcpy(tmp, vp1, size);
    memcpy(vp1, vp2, size);
    memcpy(vp2, tmp, size);
    free(tmp);
}

typedef struct CoffeeShop {
    uint64_t quality;
    uint64_t taste;
    uint64_t branding;
} coffee_shop_t;


void test_generic_ints() {
    int i1 = 1234;
    int i2 = 5678;

    swap(&i1, &i2, sizeof(int));

    assert(i1 == 5678 && "i1 should be i2's original value");
    assert(i2 == 1234 && "i2 should be i1's original value");
}

void test_generic_strings() {
    char *s1 = "dax";
    char *s2 = "adam";

    swap(&s1, &s2, sizeof(char *));
    assert(strcmp(s1, "adam") == 0 && "s1 should be s2's original value");
    assert(strcmp(s2, "dax") == 0 && "s2 should be s1's original value");
}

void test_generic_structs() {
    coffee_shop_t sbucks = {2, 3, 4};
    coffee_shop_t terminalshop = {10, 10, 10};

    swap(&sbucks, &terminalshop, sizeof(coffee_shop_t));

    assert(sbucks.quality == 10 && "sbucks.quality should be terminalshop.quality");
    assert(sbucks.taste == 10 && "sbucks.taste should be terminalshop.taste");
    assert(sbucks.branding == 10 && "sbucks.branding should be terminalshop.branding");

    assert(terminalshop.quality == 2 && "terminalshop.quality should be sbucks.quality");
    assert(terminalshop.taste == 3 && "terminalshop.taste should be sbucks.taste");
    assert(terminalshop.branding == 4 && "terminalshop.branding should be sbucks.branding");
}

int main() {
    test_generic_ints();
    test_generic_strings();
    test_generic_structs();

    printf("All tests passed!\n");

    return 0;
}