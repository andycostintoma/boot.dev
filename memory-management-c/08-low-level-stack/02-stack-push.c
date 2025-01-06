#include <assert.h>
#include <stddef.h>
#include <stdlib.h>
#include <string.h>

typedef struct Stack {
    size_t count;
    size_t capacity;
    void **data;
} stack_t;

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
    return;
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

void test_create_stack() {
    stack_t *s = stack_new(10);
    assert(s != NULL && "Failed to allocate new stack");
    assert(s->capacity == 10 && "Sets capacity to 10");
    assert(s->count == 0 && "No elements in the stack yet");
    assert(s->data != NULL && "Allocates the stack data");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

void test_push_stack() {
    stack_t *s = stack_new(2);
    assert(s != NULL && "Failed to allocate a new stack");

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 0 && "No elements in the stack yet");
    assert(s->data != NULL && "Allocates the stack data");

    int a = 1;

    stack_push(s, &a);
    stack_push(s, &a);

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 2 && "2 elements in the stack");
    assert(s->data[0] == &a && "Element inserted into stack");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

void test_push_and_double_capacity() {
    stack_t *s = stack_new(2);
    assert(s != NULL && "Failed to allocate a new stack");

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 0 && "No elements in the stack yet");
    assert(s->data != NULL && "Allocates the stack data");

    int a = 1;

    stack_push(s, &a);
    stack_push(s, &a);

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 2 && "2 elements in the stack");

    stack_push(s, &a);
    assert(s->capacity == 4 && "Capacity is doubled");
    assert(s->count == 3 && "3 elements in the stack");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

int main() {
    test_create_stack();
    test_push_stack();
    test_push_and_double_capacity();

    return 0;
}