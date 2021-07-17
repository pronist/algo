/**
 * Row priority Matrix
 *
 * 1  6 11 16 21
 * 2  7 12 17 22
 * 3  8 13 18 23
 * 4  9 14 19 24
 * 5 10 15 20 25
 */
#include <stdio.h>

#define MATRIX 5

void priority(int arr[MATRIX][MATRIX]) {
  int value = 1;
  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      arr[j][i] = value++;
     }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  priority(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
