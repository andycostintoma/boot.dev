#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>

// Definition of the structures and function prototypes

typedef struct SnekObject snek_object_t;

typedef struct {
    snek_object_t *x;
    snek_object_t *y;
    snek_object_t *z;
} snek_vector_t;

typedef struct SnekArray {
    size_t size;
    snek_object_t **elements;
} snek_array_t;

typedef enum SnekObjectKind {
    INTEGER,
    FLOAT,
    STRING,
    VECTOR3,
    ARRAY
} snek_object_kind_t;

typedef union SnekObjectData {
    int v_int;
    float v_float;
    char *v_string;
    snek_vector_t v_vector3;
    snek_array_t v_array;
} snek_object_data_t;

typedef struct SnekObject {
    snek_object_kind_t kind;
    snek_object_data_t data;
} snek_object_t;

snek_object_t *new_snek_integer(int value);

snek_object_t *new_snek_float(float value);

snek_object_t *new_snek_string(char *value);

snek_object_t *new_snek_vector3(snek_object_t *x, snek_object_t *y, snek_object_t *z);

snek_object_t *new_snek_array(size_t size);

bool snek_array_set(snek_object_t *array, size_t index, snek_object_t *value);

snek_object_t *snek_array_get(snek_object_t *array, size_t index);

size_t snek_length(snek_object_t *obj);

// Function implementations

snek_object_t *new_snek_integer(int value) {
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->kind = INTEGER;
    obj->data.v_int = value;
    return obj;
}

snek_object_t *new_snek_float(float value) {
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->kind = FLOAT;
    obj->data.v_float = value;
    return obj;
}

snek_object_t *new_snek_string(char *value) {
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    char *dst = malloc(strlen(value) + 1);
    if (dst == NULL) {
        free(obj);
        return NULL;
    }
    strcpy(dst, value);
    obj->kind = STRING;
    obj->data.v_string = dst;
    return obj;
}

snek_object_t *new_snek_vector3(snek_object_t *x, snek_object_t *y, snek_object_t *z) {
    if (x == NULL || y == NULL || z == NULL) {
        return NULL;
    }
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->kind = VECTOR3;
    obj->data.v_vector3 = (snek_vector_t){.x = x, .y = y, .z = z};
    return obj;
}

snek_object_t *new_snek_array(size_t size) {
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    snek_object_t **array = calloc(size, sizeof(snek_object_t *));
    if (array == NULL) {
        free(obj);
        return NULL;
    }
    obj->kind = ARRAY;
    snek_array_t snek_array = {.size = size, .elements = array};
    obj->data.v_array = snek_array;
    return obj;
}

bool snek_array_set(snek_object_t *array, size_t index, snek_object_t *value) {
    if (array == NULL || value == NULL || array->kind != ARRAY || array->data.v_array.size < index + 1) {
        return false;
    }
    array->data.v_array.elements[index] = value;
    return true;
}

snek_object_t *snek_array_get(snek_object_t *array, size_t index) {
    if (array == NULL || array->kind != ARRAY || array->data.v_array.size < index + 1) {
        return NULL;
    }
    return array->data.v_array.elements[index];
}

size_t snek_length(snek_object_t *obj) {
    if (obj == NULL) {
        return -1;
    }
    switch (obj->kind) {
        case INTEGER:
        case FLOAT:
            return 1;
        case STRING:
            return strlen(obj->data.v_string);
        case VECTOR3:
            return 3;
        case ARRAY:
            return obj->data.v_array.size;
        default:
            return -1;
    }
}


snek_object_t *snek_add(snek_object_t *a, snek_object_t *b) {
    if (a == NULL || b == NULL) {
        return NULL;
    }
    switch (a->kind) {
        case INTEGER:
            switch (b->kind) {
                case INTEGER:
                    return new_snek_integer(a->data.v_int + b->data.v_int);
                case FLOAT:
                    return new_snek_float(a->data.v_int + b->data.v_float);
                default:
                    return NULL;
            }
        case FLOAT:
            switch (b->kind) {
                case INTEGER:
                    return new_snek_float(a->data.v_float + b->data.v_int);
                case FLOAT:
                    return new_snek_float(a->data.v_float + b->data.v_float);
                default:
                    return NULL;
            }
        case STRING:
            if (b->kind == STRING) {
                char *dst = calloc(strlen(a->data.v_string) + strlen(b->data.v_string) + 1, sizeof(snek_object_t *));
                if (dst == NULL) {
                    return NULL;
                }
                strcpy(dst, a->data.v_string);
                strcat(dst, b->data.v_string);
                snek_object_t *new_string = new_snek_string(dst);
                free(dst);
                return new_string;
            }
            return NULL;
        case VECTOR3:
            if (b->kind == VECTOR3) {
                return new_snek_vector3(
                    snek_add(a->data.v_vector3.x, b->data.v_vector3.x),
                    snek_add(a->data.v_vector3.y, b->data.v_vector3.y),
                    snek_add(a->data.v_vector3.z, b->data.v_vector3.z)
                );
            }
            return NULL;
        case ARRAY:
            if (b->kind == ARRAY) {
                snek_object_t *dst = new_snek_array(a->data.v_array.size + b->data.v_array.size);
                if (dst == NULL) {
                    return NULL;
                }
                for (size_t i = 0; i < a->data.v_array.size; i++) {
                    snek_array_set(dst, i, snek_array_get(a, i));
                }
                for (size_t i = 0; i < b->data.v_array.size; i++) {
                    snek_array_set(dst, a->data.v_array.size + i, snek_array_get(b, i));
                }
                return dst;
            }
            return NULL;
        default:
            return NULL;
    }
}

// Test functions

void test_integer_positive() {
    snek_object_t *int_object = new_snek_integer(42);
    assert(int_object != NULL && "Failed to allocate memory for integer object");
    assert(int_object->data.v_int == 42 && "Must allow positive numbers");
    free(int_object);
}

void test_integer_zero() {
    snek_object_t *int_object = new_snek_integer(0);
    assert(int_object != NULL && "Failed to allocate memory for integer object");
    assert(int_object->kind == INTEGER && "Must be INTEGER type");
    assert(int_object->data.v_int == 0 && "Must equal zero");
    free(int_object);
}

void test_integer_negative() {
    snek_object_t *int_object = new_snek_integer(-5);
    assert(int_object != NULL && "Failed to allocate memory for integer object");
    assert(int_object->kind == INTEGER && "Must be INTEGER type");
    assert(int_object->data.v_int == -5 && "Must allow negative numbers");
    free(int_object);
}

void test_float_positive() {
    snek_object_t *obj = new_snek_float(42.0);
    assert(obj != NULL && "Failed to allocate memory for float object");
    assert(obj->data.v_float == 42.0f && "Must accept positive values");
    free(obj);
}

void test_float_zero() {
    snek_object_t *obj = new_snek_float(0.0f);
    assert(obj != NULL && "Failed to allocate memory for float object");
    assert(obj->kind == FLOAT && "Must set type to FLOAT");
    assert(obj->data.v_float == 0.0f && "Must accept 0.0");
    free(obj);
}

void test_float_negative() {
    snek_object_t *obj = new_snek_float(-5.0f);
    assert(obj != NULL && "Failed to allocate memory for float object");
    assert(obj->kind == FLOAT && "Must set type to FLOAT");
    assert(obj->data.v_float == -5.0f && "Must accept negative numbers");
    free(obj);
}

void test_string_copy() {
    const char *input = "Hello, World!";
    snek_object_t *obj = new_snek_string(input);
    assert(obj != NULL && "Failed to allocate memory for string object");
    assert(obj->kind == STRING && "Must be a string!");
    assert(strcmp(obj->data.v_string, input) == 0 && "Should copy string correctly");
    assert(obj->data.v_string != input && "You need to copy the string.");
    free(obj->data.v_string);
    free(obj);
}

void test_vector_returns_null() {
    snek_object_t *vec = new_snek_vector3(NULL, NULL, NULL);
    assert(vec == NULL && "Should return null when input is null");
}

void test_vector_multiple_objects() {
    snek_object_t *x = new_snek_integer(1);
    snek_object_t *y = new_snek_integer(2);
    snek_object_t *z = new_snek_integer(3);
    snek_object_t *vec = new_snek_vector3(x, y, z);
    assert(vec != NULL && "Should allocate a new object");
    assert(vec->data.v_vector3.x == x && "Should reference x");
    assert(vec->data.v_vector3.y == y && "Should reference y");
    assert(vec->data.v_vector3.z == z && "Should reference z");
    assert(vec->data.v_vector3.x->data.v_int == 1 && "Should have correct x");
    assert(vec->data.v_vector3.y->data.v_int == 2 && "Should have correct y");
    assert(vec->data.v_vector3.z->data.v_int == 3 && "Should have correct z");
    free(x);
    free(y);
    free(z);
    free(vec);
}

void test_vector_same_object() {
    snek_object_t *i = new_snek_integer(1);
    snek_object_t *vec = new_snek_vector3(i, i, i);
    assert(vec != NULL && "Should allocate a new object");
    assert(vec->data.v_vector3.x == i && "Should reference x");
    assert(vec->data.v_vector3.y == i && "Should reference y");
    assert(vec->data.v_vector3.z == i && "Should reference z");
    assert(vec->data.v_vector3.x->data.v_int == 1 && "Should have correct x");
    assert(vec->data.v_vector3.y->data.v_int == 1 && "Should have correct y");
    assert(vec->data.v_vector3.z->data.v_int == 1 && "Should have correct z");
    i->data.v_int = 2;
    assert(vec->data.v_vector3.x->data.v_int == 2 && "Should have correct x");
    assert(vec->data.v_vector3.y->data.v_int == 2 && "Should have correct y");
    assert(vec->data.v_vector3.z->data.v_int == 2 && "Should have correct z");
    free(i);
    free(vec);
}

void test_create_empty_array() {
    snek_object_t *obj = new_snek_array(2);
    assert(obj != NULL && "Failed to allocate memory for array object");
    assert(obj->kind == ARRAY && "Must set type to ARRAY");
    assert(obj->data.v_array.size == 2 && "Must set size to 2");
    free(obj->data.v_array.elements);
    free(obj);
}

void test_used_calloc() {
    snek_object_t *obj = new_snek_array(2);
    assert(obj != NULL && "Failed to allocate memory for array object");
    assert(obj->data.v_array.elements[0] == NULL && "Should use calloc");
    assert(obj->data.v_array.elements[1] == NULL && "Should use calloc");
    free(obj->data.v_array.elements);
    free(obj);
}

void test_array() {
    snek_object_t *obj = new_snek_array(2);
    snek_object_t *first = new_snek_string("First");
    snek_object_t *second = new_snek_integer(3);
    assert(snek_array_set(obj, 0, first));
    assert(snek_array_set(obj, 1, second));
    snek_object_t *retrieved_first = snek_array_get(obj, 0);
    assert(retrieved_first != NULL && "Should find the first object");
    assert(retrieved_first->kind == STRING && "Should be a string");
    assert(first == retrieved_first && "Should be the same object");
    snek_object_t *retrieved_second = snek_array_get(obj, 1);
    assert(retrieved_second != NULL && "Should find the second object");
    assert(retrieved_second->kind == INTEGER && "Should be an integer");
    assert(second == retrieved_second && "Should be the same object");
    free(first->data.v_string);
    free(first);
    free(second);
    free(obj->data.v_array.elements);
    free(obj);
}

void test_set_outside_bounds() {
    snek_object_t *obj = new_snek_array(2);
    snek_object_t *outside = new_snek_string("First");
    assert(snek_array_set(obj, 1, outside));
    assert(!snek_array_set(obj, 100, outside) && "Should return false if index is out of bounds");
    free(outside->data.v_string);
    free(outside);
    free(obj->data.v_array.elements);
    free(obj);
}

void test_get_outside_bounds() {
    snek_object_t *obj = new_snek_array(1);
    snek_object_t *first = new_snek_string("First");
    assert(snek_array_set(obj, 0, first));
    assert(snek_array_get(obj, 1) == NULL && "Should not access outside the array");
    free(first->data.v_string);
    free(first);
    free(obj->data.v_array.elements);
    free(obj);
}

void test_snek_length_integer() {
    snek_object_t *obj = new_snek_integer(42);
    assert(snek_length(obj) == 1 && "Integers are length 1");
    free(obj);
}

void test_snek_length_float() {
    snek_object_t *obj = new_snek_float(3.14);
    assert(snek_length(obj) == 1 && "Floats are length 1");
    free(obj);
}

void test_snek_length_string() {
    snek_object_t *shorter = new_snek_string("hello");
    assert(snek_length(shorter) == 5 && "Should return based on input");
    snek_object_t *longer = new_snek_string("hello, world");
    assert(snek_length(longer) == strlen("hello, world") && "Should use strlen");
    free(shorter->data.v_string);
    free(shorter);
    free(longer->data.v_string);
    free(longer);
}

void test_snek_length_vector3() {
    snek_object_t *i = new_snek_integer(1);
    snek_object_t *vec = new_snek_vector3(i, i, i);
    assert(snek_length(vec) == 3 && "Vec3 always has length 3");
    free(i);
    free(vec);
}

void test_snek_length_array() {
    snek_object_t *i = new_snek_integer(1);
    snek_object_t *arr = new_snek_array(4);
    assert(snek_array_set(arr, 0, i));
    assert(snek_array_set(arr, 1, i));
    assert(snek_array_set(arr, 2, i));
    assert(snek_length(arr) == 4 && "Length of array should be its size");
    free(i);
    free(arr->data.v_array.elements);
    free(arr);
}

void test_integer_add() {
    snek_object_t *one = new_snek_integer(1);
    snek_object_t *three = new_snek_integer(3);
    snek_object_t *four = snek_add(one, three);
    assert(four != NULL && "must return an object");
    assert(four->kind == INTEGER && "Kind should be INTEGER");
    assert(four->data.v_int == 4 && "1 + 3 should equal 4");
    free(one);
    free(three);
    free(four);
}

void test_float_add() {
    snek_object_t *one = new_snek_float(1.5);
    snek_object_t *three = new_snek_float(3.5);
    snek_object_t *five = snek_add(one, three);
    assert(five != NULL && "must return an object");
    assert(five->kind == FLOAT && "Kind should be FLOAT");
    assert(five->data.v_float == 5.0f && "1.5 + 3.5 should equal 5.0");
    free(one);
    free(three);
    free(five);
}

void test_string_add() {
    snek_object_t *hello = new_snek_string("hello");
    snek_object_t *world = new_snek_string(", world");
    snek_object_t *greeting = snek_add(hello, world);
    assert(greeting != NULL && "must return an object");
    assert(greeting->kind == STRING && "Kind should be STRING");
    assert(strcmp(greeting->data.v_string, "hello, world") == 0 && "Should concatenate strings");
    free(hello->data.v_string);
    free(hello);
    free(world->data.v_string);
    free(world);
    free(greeting->data.v_string);
    free(greeting);
}

void test_string_add_self() {
    snek_object_t *repeated = new_snek_string("(repeated)");
    snek_object_t *result = snek_add(repeated, repeated);
    assert(result != NULL && "must return an object");
    assert(result->kind == STRING && "Kind should be STRING");
    assert(strcmp(result->data.v_string, "(repeated)(repeated)") == 0 && "Should concatenate strings");
    free(repeated->data.v_string);
    free(repeated);
    free(result->data.v_string);
    free(result);
}

void test_vector3_add() {
    snek_object_t *one = new_snek_float(1.0);
    snek_object_t *two = new_snek_float(2.0);
    snek_object_t *three = new_snek_float(3.0);
    snek_object_t *four = new_snek_float(4.0);
    snek_object_t *five = new_snek_float(5.0);
    snek_object_t *six = new_snek_float(6.0);
    snek_object_t *v1 = new_snek_vector3(one, two, three);
    snek_object_t *v2 = new_snek_vector3(four, five, six);
    snek_object_t *result = snek_add(v1, v2);
    assert(result != NULL && "must return an object");
    assert(result->kind == VECTOR3 && "Kind should be VECTOR3");
    assert(result->data.v_vector3.x->data.v_float == 5.0f && "x component should be 5.0");
    assert(result->data.v_vector3.y->data.v_float == 7.0f && "y component should be 7.0");
    assert(result->data.v_vector3.z->data.v_float == 9.0f && "z component should be 9.0");
    free(one);
    free(two);
    free(three);
    free(four);
    free(five);
    free(six);
    free(v1);
    free(v2);
    free(result->data.v_vector3.x);
    free(result->data.v_vector3.y);
    free(result->data.v_vector3.z);
    free(result);
}

void test_array_add() {
    snek_object_t *one = new_snek_integer(1);
    snek_object_t *ones = new_snek_array(2);
    assert(snek_array_set(ones, 0, one));
    assert(snek_array_set(ones, 1, one));
    snek_object_t *hi = new_snek_string("hi");
    snek_object_t *hellos = new_snek_array(3);
    assert(snek_array_set(hellos, 0, hi));
    assert(snek_array_set(hellos, 1, hi));
    assert(snek_array_set(hellos, 2, hi));
    snek_object_t *result = snek_add(ones, hellos);
    assert(result != NULL && "must return an object");
    assert(result->kind == ARRAY && "Kind should be ARRAY");
    snek_object_t *first = snek_array_get(result, 0);
    assert(first != NULL && "should find the first item");
    assert(first->data.v_int == 1 && "First item should be an int with value 1");
    snek_object_t *third = snek_array_get(result, 2);
    assert(third != NULL && "should find the third item");
    assert(strcmp(third->data.v_string, "hi") == 0 && "third item should be 'hi'");
    free(one);
    free(ones->data.v_array.elements);
    free(ones);
    free(hi->data.v_string);
    free(hi);
    free(hellos->data.v_array.elements);
    free(hellos);
    free(result->data.v_array.elements);
    free(result);
}


int main() {
    test_integer_positive();
    test_integer_zero();
    test_integer_negative();
    test_float_positive();
    test_float_zero();
    test_float_negative();
    test_string_copy();
    test_vector_returns_null();
    test_vector_multiple_objects();
    test_vector_same_object();
    test_create_empty_array();
    test_used_calloc();
    test_array();
    test_set_outside_bounds();
    test_get_outside_bounds();
    test_snek_length_integer();
    test_snek_length_float();
    test_snek_length_string();
    test_snek_length_vector3();
    test_snek_length_array();
    test_integer_add();
    test_float_add();
    test_string_add();
    test_string_add_self();
    test_vector3_add();
    test_array_add();
    return 0;
}
