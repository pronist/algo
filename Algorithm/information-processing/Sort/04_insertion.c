/**
 * Insertion sort
 */
#include <stdio.h>

#define ARRAY_SIZE 5

void insertion_sort(int* arr) {
  for (int i = 1; i < ARRAY_SIZE; i++) {
    int temp = arr[i];
    for (int j = 0; j < ARRAY_SIZE; j++) {
      if (arr[i] < arr[j]) {
        for (int q = i; q > j; q--) {
          arr[q] = arr[q-1];
        }
        arr[j] = temp;
        break;
      }
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 15, 11, 1, 3, 8 };

  insertion_sort(arr);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%2d ", arr[i]);
  }
}
