/**
 * Bubble sort
 */
#include <stdio.h>

#define ARRAY_SIZE 5

void bubble_sort(int* arr) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    for (int j = 1; j < ARRAY_SIZE; j++) {
      if (arr[j - 1] > arr[j]) {
        int temp = arr[j - 1];
        arr[j - 1] = arr[j];
        arr[j] = temp;
      }
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 15, 11, 1, 3, 8 };

  bubble_sort(arr);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%2d ", arr[i]);
  }
}
