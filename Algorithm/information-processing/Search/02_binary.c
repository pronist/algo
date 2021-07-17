/**
 * Binary search
 */
#include <stdio.h>

#define ARRAY_SIZE 10

int binary_search(int* arr, int search) {
  int high = ARRAY_SIZE;
  int low = 0;
  int current = (int) ARRAY_SIZE / 2;

  while (1) {
    if (low >= high) break;

    if (arr[current] == search) {
      return arr[current];
    } else if (search > arr[current]) {
      low = current + 1;
    } else {
      high = current - 1;
    }
    current = (int) (high+low) / 2;
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 14, 17, 20, 22, 26, 48, 50, 90, 95, 100 };

  printf("%d", binary_search(arr, 95));
}
