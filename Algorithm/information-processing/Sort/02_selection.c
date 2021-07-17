/**
 * Selection sort
 */
#include <stdio.h>

#define ARRAY_SIZE 5

void selection_sort(int* arr) {
  for (int i = 0; i < ARRAY_SIZE - 1; i++) {
    int mi = -1, min = arr[i];
    for (int j = i + 1; j < ARRAY_SIZE; j++) {
      if (min > arr[j]) {
        min = arr[j];
        mi = j;
      };
    }
    if (mi != -1) {
      int temp = arr[i];
      arr[i] = arr[mi];
      arr[mi] = temp;
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 15, 11, 1, 3, 8 };

  selection_sort(arr);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%2d ", arr[i]);
  }
}
