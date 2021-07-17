/**
 * Hourglass Matrix
 *
 *  1  2  3  4  5
 *     6  7  8
 *        9
 *    10 11 12
 * 13 14 15 16 17
 */
#include <stdio.h>

#define MATRIX 5

void hourglass(int arr[MATRIX][MATRIX]) {
  int value = 1;
  int left = 0, right = 5;
  for (int i = 0; i < MATRIX; i++) {
    for (int j = left; j < right; j++) {
      arr[i][j] = value++;
    }
    if (i < 2) {
      left++; right--;
    } else {
      left--; right++;
    }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  hourglass(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
