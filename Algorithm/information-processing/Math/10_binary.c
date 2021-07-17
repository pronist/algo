/**
 * Decimal to Binary
 */
#include <stdio.h>
#include <stdlib.h>

#define ARRAY_SIZE 10

int to_binary(int num, int* output) {
  int cnt = 0;
  num = abs(num);
  for (int i = ARRAY_SIZE - 1;; i--) {
    if (num == 1) {
      output[i] = 1; break;
    }
    cnt++;
    output[i] = num % 2;
    num = (int) num / 2;
  }
  return cnt;
}

void to_2complement(int* input, int cnt) {
  for (int i = ARRAY_SIZE - cnt - 1; i < ARRAY_SIZE; i++) {
    input[i] = 1 - input[i];
  }
  for (int i = ARRAY_SIZE - 1; i >= ARRAY_SIZE - cnt - 1; i--) {
    if ((input[i] + 1) == 2) {
      input[i] = 0;
    } else {
      input[i]++; break;
    }
  }
}

void main() {
  // int num = 11; // 1011
  int num = -11; // 0101
  int output[ARRAY_SIZE] = { 0 };

  int cnt = to_binary(num, output);
  if (num < 0) to_2complement(output, cnt);

  for (int i = ARRAY_SIZE - cnt - 1; i < ARRAY_SIZE; i++) {
    printf("%d", output[i]);
  }
}
