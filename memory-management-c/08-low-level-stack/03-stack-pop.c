#include <assert.h>
#include <stddef.h>
#include <stdlib.h>

typedef struct Stack {
    size_t count;
    size_t capacity;
    void **data;
} stack_t;

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

    stack_push(s, &a);
    assert(s->capacity == 4 && "Capacity is doubled");
    assert(s->count == 3 && "3 elements in the stack");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

void test_pop_stack() {
    stack_t *s = stack_new(2);
    assert(s != NULL && "Failed to allocate a new stack");

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 0 && "No elements in the stack yet");
    assert(s->data != NULL && "Allocates the stack data");

    int one = 1;
    int two = 2;
    int three = 3;

    stack_push(s, &one);
    stack_push(s, &two);

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 2 && "2 elements in the stack");

    stack_push(s, &three);
    assert(s->capacity == 4 && "Capacity is doubled");
    assert(s->count == 3 && "3 elements in the stack");

    int *popped = stack_pop(s);
    assert(*popped == three && "Should pop the last element");

    popped = stack_pop(s);
    assert(*popped == two && "Should pop the last element");

    popped = stack_pop(s);
    assert(*popped == one && "Should pop the only remaining element");

    popped = stack_pop(s);
    assert(popped == NULL && "No remaining elements");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

void test_pop_stack_empty() {
    stack_t *s = stack_new(2);
    assert(s != NULL && "Failed to allocate a new stack");

    assert(s->capacity == 2 && "Sets capacity to 2");
    assert(s->count == 0 && "No elements in the stack yet");
    assert(s->data != NULL && "Allocates the stack data");

    int *popped = stack_pop(s);
    assert(popped == NULL && "Should return null when popping an empty stack");

    // Clean up our allocated data.
    free(s->data);
    free(s);
}

int main() {
    test_create_stack();
    test_push_stack();
    test_pop_stack();
    test_pop_stack_empty();

    return 0;
}