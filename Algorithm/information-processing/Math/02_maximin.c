/**
 * Max, Min
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void get_maxmin(int* arr, int* min, int* max) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    *max = *max > arr[i]? *max: arr[i];
    *min = *min < arr[i]? *min: arr[i];
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 23, 50, 30, 54, 80, 92, 58, 100, 33, 52 };
  int min = 9999;
  int max = -9999;

  get_maxmin(arr, &min, &max);

  printf("max: %d, min: %d", max, min);
}
