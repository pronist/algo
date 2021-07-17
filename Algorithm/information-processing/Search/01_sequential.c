/**
 * Sequential search
 */
#include <stdio.h>

#define ARRAY_SIZE 10

int sequential_search(int *arr, int search) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    if (arr[i] == search) {
      return arr[i];
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 14, 17, 20, 22, 26, 48, 50, 90, 95, 100 };

  printf("%d", sequential_search(arr, 26));
}
