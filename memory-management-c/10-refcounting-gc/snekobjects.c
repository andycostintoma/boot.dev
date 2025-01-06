#include <assert.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
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
    int refcount;
} snek_object_t;

snek_object_t *_new_snek_object();

void refcount_free(snek_object_t *obj);

void refcount_inc(snek_object_t *obj);

void refcount_dec(snek_object_t *obj);

snek_object_t *new_snek_integer(int value);

snek_object_t *new_snek_float(float value);

snek_object_t *new_snek_string(const char *value);

snek_object_t *new_snek_vector3(snek_object_t *x, snek_object_t *y, snek_object_t *z);

snek_object_t *new_snek_array(size_t size);

bool snek_array_set(snek_object_t *array, size_t index, snek_object_t *value);

snek_object_t *snek_array_get(snek_object_t *array, size_t index);

size_t snek_length(snek_object_t *obj);

snek_object_t *_new_snek_object() {
    snek_object_t *obj = malloc(sizeof(snek_object_t));
    if (obj == NULL) {
        return NULL;
    }
    obj->refcount = 1;
    return obj;
}

void refcount_free(snek_object_t *obj) {
    switch (obj->kind) {
        case INTEGER:
        case FLOAT:
            break;
        case STRING:
            free(obj->data.v_string);
            break;
        case VECTOR3: {
            snek_vector_t vec = obj->data.v_vector3;
            refcount_dec(vec.x);
            refcount_dec(vec.y);
            refcount_dec(vec.z);
        }
        break;
        case ARRAY: {
            snek_array_t snek_array = obj->data.v_array;
            for (size_t i = 0; i < snek_array.size; i++) {
                refcount_dec(snek_array.elements[i]);
            }
            free(snek_array.elements);
        }
        break;
        default:
            assert(false);
    }
    free(obj);
}

void refcount_inc(snek_object_t *obj) {
    if (obj == NULL) {
        return;
    }
    obj->refcount++;
}

void refcount_dec(snek_object_t *obj) {
    if (obj == NULL) {
        return;
    }
    obj->refcount--;
    if (obj->refcount == 0) {
        refcount_free(obj);
    }
}

snek_object_t *new_snek_integer(int value) {
    snek_object_t *obj = _new_snek_object();
    if (obj == NULL) {
        return NULL;
    }
    obj->kind = INTEGER;
    obj->data.v_int = value;
    return obj;
}

snek_object_t *new_snek_float(float value) {
    snek_object_t *obj = _new_snek_object();
    if (obj == NULL) {
        return NULL;
    }
    obj->kind = FLOAT;
    obj->data.v_float = value;
    return obj;
}

snek_object_t *new_snek_string(const char *value) {
    snek_object_t *obj = _new_snek_object();
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

    refcount_inc(x);
    refcount_inc(y);
    refcount_inc(z);

    snek_object_t *obj = _new_snek_object();
    if (obj == NULL) {
        refcount_dec(x);
        refcount_dec(y);
        refcount_dec(z);
        return NULL;
    }

    obj->kind = VECTOR3;
    obj->data.v_vector3 = (snek_vector_t){.x = x, .y = y, .z = z};

    return obj;
}

snek_object_t *new_snek_array(size_t size) {
    snek_object_t *obj = _new_snek_object();
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
    if (array->data.v_array.elements[index] != NULL) {
        refcount_dec(array->data.v_array.elements[index]);
    }
    array->data.v_array.elements[index] = value;
    refcount_inc(value);
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
                char *dst = calloc(strlen(a->data.v_string) + strlen(b->data.v_string) + 1, sizeof(char));
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

void test_snek_integer() {
    snek_object_t *int_obj = new_snek_integer(42);
    assert(int_obj != NULL);
    assert(int_obj->kind == INTEGER);
    assert(int_obj->data.v_int == 42);
    assert(int_obj->refcount == 1);

    refcount_inc(int_obj);
    assert(int_obj->refcount == 2);

    refcount_dec(int_obj);
    assert(int_obj->refcount == 1);

    refcount_dec(int_obj);
    // int_obj should be freed now, accessing it may cause undefined behavior
}

void test_snek_string() {
    snek_object_t *str_obj = new_snek_string("hello");
    assert(str_obj != NULL);
    assert(str_obj->kind == STRING);
    assert(strcmp(str_obj->data.v_string, "hello") == 0);
    assert(str_obj->refcount == 1);

    refcount_inc(str_obj);
    assert(str_obj->refcount == 2);

    refcount_dec(str_obj);
    assert(str_obj->refcount == 1);

    refcount_dec(str_obj);
    // str_obj should be freed now, accessing it may cause undefined behavior
}

void test_snek_vector3() {
    snek_object_t *x = new_snek_integer(1);
    snek_object_t *y = new_snek_integer(2);
    snek_object_t *z = new_snek_integer(3);
    snek_object_t *vec = new_snek_vector3(x, y, z);

    assert(vec != NULL);
    assert(vec->kind == VECTOR3);
    assert(vec->data.v_vector3.x == x);
    assert(vec->data.v_vector3.y == y);
    assert(vec->data.v_vector3.z == z);
    assert(vec->refcount == 1);

    assert(x->refcount == 2);
    assert(y->refcount == 2);
    assert(z->refcount == 2);

    refcount_dec(vec);
    // vec and its elements should be freed now, accessing them may cause undefined behavior
}

void test_snek_array() {
    snek_object_t *arr = new_snek_array(2);
    assert(arr != NULL);
    assert(arr->kind == ARRAY);
    assert(arr->data.v_array.size == 2);
    assert(arr->refcount == 1);

    snek_object_t *val1 = new_snek_string("test");
    snek_object_t *val2 = new_snek_float(3.14f);

    snek_array_set(arr, 0, val1);
    snek_array_set(arr, 1, val2);

    assert(snek_array_get(arr, 0) == val1);
    assert(snek_array_get(arr, 1) == val2);
    assert(val1->refcount == 2);
    assert(val2->refcount == 2);

    refcount_dec(arr);
    // arr and its elements should be freed now, accessing them may cause undefined behavior
}

void test_snek_add() {
    snek_object_t *int1 = new_snek_integer(1);
    snek_object_t *int2 = new_snek_integer(2);

    snek_object_t *sum_int = snek_add(int1, int2);
    assert(sum_int->kind == INTEGER);
    assert(sum_int->data.v_int == 3);
    refcount_dec(sum_int);

    snek_object_t *float1 = new_snek_float(1.5f);
    snek_object_t *sum_float = snek_add(int1, float1);
    assert(sum_float->kind == FLOAT);
    assert(sum_float->data.v_float == 2.5f);
    refcount_dec(sum_float);

    snek_object_t *str1 = new_snek_string("hello");
    snek_object_t *str2 = new_snek_string(" world");
    snek_object_t *concat_str = snek_add(str1, str2);
    assert(concat_str->kind == STRING);
    assert(strcmp(concat_str->data.v_string, "hello world") == 0);
    refcount_dec(concat_str);

    snek_object_t *vec1 = new_snek_vector3(new_snek_integer(1), new_snek_integer(2), new_snek_integer(3));
    snek_object_t *vec2 = new_snek_vector3(new_snek_integer(4), new_snek_integer(5), new_snek_integer(6));
    snek_object_t *sum_vec = snek_add(vec1, vec2);
    assert(sum_vec->kind == VECTOR3);
    assert(sum_vec->data.v_vector3.x->data.v_int == 5);
    assert(sum_vec->data.v_vector3.y->data.v_int == 7);
    assert(sum_vec->data.v_vector3.z->data.v_int == 9);

    refcount_dec(vec1);
    refcount_dec(vec2);
    refcount_dec(sum_vec);

    snek_object_t *arr1 = new_snek_array(2);
    snek_object_t *arr2 = new_snek_array(2);

    snek_array_set(arr1, 0, new_snek_integer(1));
    snek_array_set(arr1, 1, new_snek_integer(2));
    snek_array_set(arr2, 0, new_snek_integer(3));
    snek_array_set(arr2, 1, new_snek_integer(4));

    snek_object_t *concat_arr = snek_add(arr1, arr2);
    assert(concat_arr->kind == ARRAY);
    assert(snek_length(concat_arr) == 4);

    refcount_dec(arr1);
    refcount_dec(arr2);
    refcount_dec(concat_arr);
}

int main() {
    test_snek_integer();
    test_snek_string();
    test_snek_vector3();
    test_snek_array();
    test_snek_add();
    printf("All tests passed!\n");
    return 0;
}
