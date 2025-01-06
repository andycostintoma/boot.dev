#include <assert.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>

typedef struct Stack {
    size_t count;
    size_t capacity;
    void **data;
} stack_t;

void stack_free(stack_t *stack) {
    if (stack == NULL || stack->data == NULL) {
        return;
    }

    free(stack->data);
    free(stack);
}

void *stack_pop(stack_t *stack) {
    if (stack->count == 0) {
        return NULL;
    }

    stack->count--;
    return stack->data[stack->count];
}

void stack_push(stack_t *stack, void *obj) {
    if (stack->count == stack->capacity) {
        stack->capacity *= 2;
        void **temp = realloc(stack->data, stack->capacity * sizeof(void *));
        if (temp == NULL) {
            free(stack->data);
            exit(1);
        }
        stack->data = temp;
    }
    stack->data[stack->count] = obj;
    stack->count++;
}

stack_t *stack_new(size_t capacity) {
    stack_t *stack = malloc(sizeof(stack_t));
    if (stack == NULL) {
        return NULL;
    }

    stack->count = 0;
    stack->capacity = capacity;
    stack->data = malloc(stack->capacity * sizeof(void *));
    if (stack->data == NULL) {
        free(stack);
        return NULL;
    }
    return stack;
}

void scary_double_push(stack_t *s) {
    stack_push(s, (void *) 1337);
    int *new = malloc(sizeof(int));
    *new = 1024;
    stack_push(s, new);
}

void stack_push_multiple_types(stack_t *s) {
    float *f = malloc(sizeof(float));
    *f = 3.14;
    stack_push(s, f);

    const char *string = "Sneklang is blazingly slow!";
    size_t len = strlen(string) + 1;

    char *c = malloc(sizeof(char) * len);
    strcpy(c, string);
    stack_push(s, c);
}


void test_heterogenous_stack() {
    stack_t *s = stack_new(2);
    assert(s != NULL && "Must allocate a new stack");

    scary_double_push(s);
    assert(s->count == 2 && "Should have two items in the stack");

    int value = (int) s->data[0];
    assert(value == 1337 && "Zero item should be 1337");

    int *pointer = s->data[1];
    assert(*pointer == 1024 && "Top item should be 1024");

    free(pointer);
    stack_free(s);
}

void test_multiple_types_stack() {
    stack_t *s = stack_new(4);
    assert(s != NULL && "Must allocate a new stack");

    stack_push_multiple_types(s);
    assert(s->count == 2 && "Should have two items in the stack");

    float *f = s->data[0];
    assert(*f == 3.14f && "Float is equal");

    char *string = s->data[1];
    assert(strcmp(string, "Sneklang is blazingly slow!") == 0 && "char* is equal");

    free(f);
    free(string);
    stack_free(s);
}

int main() {
    test_heterogenous_stack();
    test_multiple_types_stack();
    return 0;
}
