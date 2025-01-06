#include <assert.h>
#include <stdio.h>

typedef enum SnekObjectKind {
    INTEGER,
    FLOAT,
    BOOL,
} snek_object_kind_t;

typedef struct SnekInt {
    char *name;
    int value;
} snek_int_t;

typedef struct SnekFloat {
    char *name;
    float value;
} snek_float_t;

typedef struct SnekBool {
    char *name;
    unsigned int value;
} snek_bool_t;

void snek_zero_out(void *ptr, snek_object_kind_t kind) {
    switch (kind) {
        case INTEGER:
            ((snek_int_t*)ptr)->value = 0;
        break;
        case FLOAT:
            ((snek_float_t*)ptr)->value = 0.0;
        break;
        case BOOL:
            ((snek_bool_t*)ptr)->value = 0;
        break;
        default:
            break;
    }
}

void test_zero_out_integer() {
    snek_int_t integer;
    integer.value = 42;
    snek_zero_out(&integer, INTEGER);
    assert(integer.value == 0);
}

void test_zero_out_float() {
    snek_float_t float_num;
    float_num.value = 3.14;
    snek_zero_out(&float_num, FLOAT);
    assert(float_num.value == 0.0f);
}

void test_zero_out_bool() {
    snek_bool_t boolean;
    boolean.value = 1;
    snek_zero_out(&boolean, BOOL);
    assert(boolean.value == 0);
}

void test_zero_out_nonzero_values() {
    snek_int_t integer;
    snek_float_t float_num;
    snek_bool_t boolean;

    integer.value = -100;
    float_num.value = -99.99f;
    boolean.value = 255;

    snek_zero_out(&integer, INTEGER);
    snek_zero_out(&float_num, FLOAT);
    snek_zero_out(&boolean, BOOL);

    assert(integer.value == 0);
    assert(float_num.value == 0.0f);
    assert(boolean.value == 0);
}

int main() {
    test_zero_out_integer();
    test_zero_out_float();
    test_zero_out_bool();
    test_zero_out_nonzero_values();

    printf("All tests passed!\n");

    return 0;
}