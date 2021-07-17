/**
 * In A(M) -> 170 ~ 250, Classification each shoes size
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void get_stock(int* arr, int* output) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    output[arr[i]] += 1;
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 170, 185, 200, 175, 195, 210, 170, 200, 225, 245 };
  int output[251] = { 0 };

  get_stock(arr, output);

  for (int i = 170; i <= 250; i += 5) {
    printf("size: %d, count: %d\n", i, output[i]);
  }
}
