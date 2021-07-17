/**
 * In A(M) -> Max, Min, Sum, Average
 */
#include <stdio.h>

#define ARRAY_SIZE 5

void get_maxmin(int* arr, int* min, int* max) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    *max = *max > arr[i]? *max: arr[i];
    *min = *min < arr[i]? *min: arr[i];
  }
}

int get_sum(int* arr) {
  int sum = 0;
  for (int i = 0; i < ARRAY_SIZE; i++) {
    sum += arr[i];
  }
  return sum;
}

float get_avg(int sum) {
  return (float) sum / ARRAY_SIZE;
}

void main() {
  int arr[ARRAY_SIZE] = { 100, 55, 25, 45, 75 };

  int min = 999999;
  int max = 0;
  int sum = get_sum(arr);
  float avg = get_avg(sum);

  get_maxmin(arr, &min, &max);

  printf("Min: %d, Max: %d, Sum: %d, Avg: %.2f", min, max, sum, avg);
}
