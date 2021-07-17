/**
 * N >= Get -> Prime number
 */
#include <stdio.h>
#include <math.h>

#define ARRAY_SIZE 50

void get_prime(int limit, int* output) {
  int s = -1, oi = -1;
  for (int i = 2; i <= limit; i++) {
    s = -1;
    for (int j = 2; j <= (int) sqrt((double) i); j++) {
      if (i % j == 0 || i == j) {
        s = -s; break;
      };
    }
    if (s != 1) {
      output[++oi] = i;
      s = -s;
    }
  }
}

void main() {
  int num = 200;
  int output[ARRAY_SIZE] = { 0 };

  get_prime(num, output);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%d ", output[i]);
  }
}
