#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <string.h>

// Forward declaration of the struct Node and a typedef
typedef struct Node node_t;

// Definition of the struct Node
typedef struct Node {
    int value;
    node_t *next;
} node_t;

typedef struct SnekObject snekobject_t;

typedef struct SnekObject {
    char *name;
    snekobject_t *child;
} snekobject_t;

snekobject_t new_node(char *name) {
    snekobject_t node = {
        .name = name,
        .child = NULL
    };
    return node;
}

void test_new_node_simple() {
    snekobject_t node = new_node("root");
    assert(strcmp(node.name, "root") == 0);
    assert(node.child == NULL);
}

void test_new_node_empty_name() {
    snekobject_t node = new_node("");
    assert(strcmp(node.name, "") == 0);
    assert(node.child == NULL);
}

void test_new_node_with_child() {
    snekobject_t child = new_node("child");
    snekobject_t parent = new_node("parent");
    parent.child = &child;

    assert(strcmp(parent.name, "parent") == 0);
    assert(parent.child != NULL);
    assert(strcmp(parent.child->name, "child") == 0);
    assert(parent.child->child == NULL);
}


int main() {
    test_new_node_simple();
    test_new_node_empty_name();
    test_new_node_with_child();

    printf("All tests passed!\n");

    return 0;
}