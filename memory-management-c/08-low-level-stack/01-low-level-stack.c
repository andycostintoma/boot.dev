#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

typedef struct Stack {
  size_t count;
  size_t capacity;
  void **data;
} stack_t;

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

void test_create_stack_small() {
  stack_t *s = stack_new(3);
  assert(s != NULL);
  assert(s->capacity == 3 && "Sets capacity to 3");
  assert(s->count == 0 && "No elements in the stack yet");
  assert(s->data != NULL && "Allocates the stack data");

  free(s->data);
  free(s);
}

void test_create_stack_large() {
  stack_t *s = stack_new(100);
  assert(s != NULL);
  assert(s->capacity == 100 && "Sets capacity to 100");
  assert(s->count == 0 && "No elements in the stack yet");
  assert(s->data != NULL && "Allocates the stack data");

  free(s->data);
  free(s);
}

int main() {
  test_create_stack_small();
  test_create_stack_large();

  printf("All tests passed!\n");

  return 0;
}