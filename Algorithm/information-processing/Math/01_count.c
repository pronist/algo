/**
 * Get -> A(M) >= 80
 */
#include <stdio.h>

#define ARRAY_SIZE 10

int get_count(int lowest, int* arr) {
  int cnt = 0;
  for (int i = 0; i < ARRAY_SIZE; i++) {
    if (arr[i] >= lowest) ++cnt;
  }
  return cnt;
}

void main() {
  int arr[ARRAY_SIZE] = { 10, 50, 30, 54, 80, 92, 58, 100, 33, 52 };

  printf("%d", get_count(80, arr));
}
