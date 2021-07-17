/**
 * Sum, Average
 */
#include <stdio.h>

#define ARRAY_SIZE 10

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
  int arr[ARRAY_SIZE] = { 10, 50, 30, 54, 80, 92, 58, 100, 33, 52 };
  int sum = get_sum(arr);

  printf("sum: %d, avg: %.2f", sum, get_avg(sum));
}
