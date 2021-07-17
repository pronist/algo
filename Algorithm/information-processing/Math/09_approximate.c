/**
 * In A(M), N (9999 > N) Get min difference
 */
#include <stdio.h>
#include <stdlib.h>

#define ARRAY_SIZE 10

int get_min_diff(int num, int* arr) {
  int diff = 9999;
  for (int i = 0; i < ARRAY_SIZE; i++) {
    diff = abs(arr[i] - num) > diff ? diff: abs(arr[i] - num);
  }
  return diff;
}

void main() {
  int num = 33;
  int arr[ARRAY_SIZE] = { 50, 22, 15, 100, 42, 68, 30, 81, 55, 42 };

  printf("%d", get_min_diff(num, arr));
}
