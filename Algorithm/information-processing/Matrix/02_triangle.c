/**
 * Triangle Matrix
 *
 * 1  2  3  4  5
 *    6  7  8  9
 *      10 11 12
 *         13 14
 *            15
 */
#include <stdio.h>

#define MATRIX 5

void triangle(int arr[MATRIX][MATRIX]) {
  int value = 1;
  for (int i = 0; i < MATRIX; i++) {
    for (int j = i; j < MATRIX; j++) {
      arr[i][j] = value++;
     }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  triangle(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
