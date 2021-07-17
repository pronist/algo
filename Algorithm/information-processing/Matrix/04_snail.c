/**
 * Snail Matrix
 *
 *  1  2  3  4  5
 * 16 17 18 19  6
 * 15 24 25 20  7
 * 14 23 22 21  8
 * 13 12 11 10  9
 */
#include <stdio.h>
#include <math.h>

#define MATRIX 5

void snail(int arr[MATRIX][MATRIX]) {
  int value = 1;

  /**
   *  1 => RIGHT, DOWN
   * -1 => LEFT, UP
   */
  int sign = 1;

  int row = 0;
  int column = -1;

  int end = MATRIX;

  while (end > 0) {
    for (int i = 0; i < end; i++) {
      column += sign;
      arr[row][column] = value++;
    }
    end--;
    for (int i = 0; i < end; i++) {
      row += sign;
      arr[row][column] = value++;
    }
    sign *= -1;
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  snail(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
