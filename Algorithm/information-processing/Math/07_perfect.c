/**
 * Perfect number
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void get_perfect(int to, int from, int* output) {
  int oi = -1;
  for (int i = to; i <= from; i++) {
    int sum = 0;
    for (int j = 1; j <= (int) (i / 2); j++) {
      if (i % j == 0) sum += j;
    }
    if (i == sum) output[++oi] = i;
  }
}

void main() {
  int output[ARRAY_SIZE] = { 0 };

  get_perfect(4, 500, output);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%d ", output[i]);
  }
}
