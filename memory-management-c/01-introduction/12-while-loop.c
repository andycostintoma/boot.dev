#include <stdio.h>

void print_numbers_reverse(int start, int end){
    int i = start;
    while(i >= end){
        printf("%d\n", i);
        i--;
    }
}


int main(){
  print_numbers_reverse(20, 4);
  return 0;
}