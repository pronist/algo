/**
 * 0 < N -> Destructuring
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void destructuring(int num, int* output) {
  int i = 2, oi = -1;
  while (num != 1) {
    if (num % i == 0) {
      output[++oi] = i;
      num = num / i;
      i = 2;
    } else i++;
  }
}

void main() {
  int num = 108;
  int output[ARRAY_SIZE] = { 0 };

  destructuring(108, output);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%d ", output[i]);
  }
}
