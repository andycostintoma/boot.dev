#include <stdlib.h>
#include <stdio.h>

int *allocate_scalar_list(int size, int multiplier) {
    int *lst = malloc(size * sizeof(int));
    if (lst == NULL) {
        return NULL;
    }
    for (int i = 0; i < size; i++) {
        lst[i] = i * multiplier;
    }
    return lst;
}

int main(){
    const int num_lists = 500;
    for (int i = 0; i < num_lists; i++) {
        int *lst = allocate_scalar_list(50000, 2);
        if (lst == NULL) {
            printf("Failed to allocate list\n");
            return 1;
        }
        printf("Allocated list %d\n", i);
        free(lst);
    }
    return 0;
}

