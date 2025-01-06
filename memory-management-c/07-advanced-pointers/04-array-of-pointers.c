#include <assert.h>
#include <stdlib.h>
#include <string.h>

typedef struct Token {
    char *literal;
    int line;
    int column;
} token_t;


/**
* @brief Creates an array of token pointers on the heap.
*
* This function takes an array of `token_t` structures and the count of tokens,
* allocates memory for an array of token pointers, and copies the contents of each
* token into newly allocated memory on the heap. The function returns a pointer
* to the array of these token pointers.
* */

token_t **create_token_pointer_array(token_t *tokens, size_t count) {
    token_t **token_pointers = malloc(count * sizeof(token_t *));
    if (token_pointers == NULL) {
        exit(1);
    }
    for (size_t i = 0; i < count; ++i) {
        token_t *token = malloc(sizeof(token_t));
        if (token == NULL) {
            exit(1);
        }
        *token = tokens[i]; // Copy the contents of the struct
        token_pointers[i] = token; // Store the address of the newly allocated struct
    }
    return token_pointers;
}

void test_create_token_pointer_array_single() {
    token_t token = {"hello", 1, 1};
    token_t **result = create_token_pointer_array(&token, 1);

    assert(result != NULL);
    assert(result[0] != NULL);
    assert(strcmp(result[0]->literal, "hello") == 0);
    assert(result[0]->line == 1);
    assert(result[0]->column == 1);
    assert(result[0] != &token);

    free(result[0]);
    free(result);
}

void test_create_token_pointer_array_multiple() {
    token_t tokens[3] = {
        {"foo", 1, 1},
        {"bar", 2, 5},
        {"baz", 3, 10}
    };
    token_t **result = create_token_pointer_array(tokens, 3);

    assert(result != NULL);
    for (int i = 0; i < 3; i++) {
        assert(result[i] != NULL);
        assert(strcmp(result[i]->literal, tokens[i].literal) == 0);
        assert(result[i]->line == tokens[i].line);
        assert(result[i]->column == tokens[i].column);
        assert(result[i] != &tokens[i]);
    }

    for (int i = 0; i < 3; i++) {
        free(result[i]);
    }
    free(result);
}

void test_create_token_pointer_array_memory_allocation() {
    token_t tokens[2] = {
        {"test1", 1, 1},
        {"test2", 2, 2}
    };
    token_t **result = create_token_pointer_array(tokens, 2);

    assert(result != NULL);
    assert(result[0] != NULL);
    assert(result[1] != NULL);
    assert(result[0] != result[1]);
    assert(result[0] != &tokens[0]);
    assert(result[1] != &tokens[1]);

    free(result[0]);
    free(result[1]);
    free(result);
}

int main() {
    test_create_token_pointer_array_single();
    test_create_token_pointer_array_multiple();
    test_create_token_pointer_array_memory_allocation();

    return 0;
}
