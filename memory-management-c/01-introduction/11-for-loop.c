#include <stdio.h>
void print_numbers(int start, int end){
    for (int i=start; i <= end; i++){
        printf("%d\n", i);
    }
}

int main(){
  print_numbers(42, 69);
  return 0;
}