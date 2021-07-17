/**
 * CM
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void get_CM(int x, int y, int* arr, int* output) {
  int oi = -1;
  for (int i = 0; i < ARRAY_SIZE; i++) {
    if (arr[i] % x == 0 && arr[i] % y == 0) {
      output[++oi] = arr[i];
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 21, 17, 4, 41, 24, 75, 40, 27, 48, 72 };
  int output[ARRAY_SIZE] = { 0 };

  get_CM(3, 4, arr, output);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%d ", output[i]);
  }
}
