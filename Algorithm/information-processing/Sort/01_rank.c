/**
 * Rank
 */
#include <stdio.h>

#define ARRAY_SIZE 5

void get_rank(int* arr, int* rank) {
  for (int i = 0; i < ARRAY_SIZE; i++) {
    rank[i] = 1;
    for (int j = 0; j < ARRAY_SIZE; j++) {
      if (arr[i] < arr[j]) {
        rank[i] += 1;
      }
    }
  }
}

void main() {
  int arr[ARRAY_SIZE] = { 8, 4, 3, 2, 6 };
  int rankedArray[ARRAY_SIZE] = { 0 };

  get_rank(arr, rankedArray);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%2d %2d\n", arr[i], rankedArray[i]);
  }
}
