/**
 * Quick sort
 */
#include <stdio.h>

#define ARRAY_SIZE 10

void quick_sort(int start, int end, int* arr) {
  if (start >= end) return;

  int i = start + 1;
  int j = end;

  while (1) {
    while (arr[start] > arr[i] && i < end) i++;
    while (arr[start] < arr[j] && j > start) j--;
    if (i >= j) break;

    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
  }

  int temp = arr[j];
  arr[j] = arr[start];
  arr[start] = temp;

  quick_sort(start, j-1, arr); quick_sort(j+1, end, arr);
}

void main() {
  int arr[ARRAY_SIZE] = { 26, 14, 100, 95, 22, 17, 48, 20, 50, 90 };

  quick_sort(0, 9, arr);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%2d ", arr[i]);
  }
}
