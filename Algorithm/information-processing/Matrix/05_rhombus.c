/**
 * Rhombus Matrix
 *
 *       1
 *    3  5  7
 * 9 11 13 15 17
 *   19 21 23
 *      25
 */
#include <stdio.h>
#include <math.h>

#define MATRIX 5

void rhombus(int arr[MATRIX][MATRIX]) {
  int value = 1;

  int pos = (int) MATRIX / 2;
  int left = pos, right = pos;

  for (int i = 0; i < MATRIX; i++) {
    for (int j = left; j <= right; j++) {
      arr[i][j] = value;
      value += 2;
    }
    if (i < pos) {
      left--; right++;
    } else {
      left++; right--;
    }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  rhombus(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
