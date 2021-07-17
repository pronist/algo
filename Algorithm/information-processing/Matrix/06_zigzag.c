/**
 * ZigZag Matrix
 *
 *  1  2  3  4  5
 * 10  9  8  7  6
 * 11 12 13 14 15
 * 20 19 18 17 16
 * 21 22 23 24 25
 */
#include <stdio.h>

#define MATRIX 5

void zigzag(int arr[MATRIX][MATRIX]) {
  int value = 1;

  int sign = 1;
  int left = 0, right = MATRIX-1;

  for (int i = 0; i < MATRIX; i++) {
    for (int j = left; j != right+sign; j += sign) {
      arr[i][j] = value++;
    }
    int temp = left;
    left = right;
    right = temp;

    sign *= -1;
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  zigzag(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
