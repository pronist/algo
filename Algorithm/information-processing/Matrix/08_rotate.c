/**
 * Rotate Matrix
 *
 *  1  2  3  4  5       21 16 11  6  1
 *  6  7  8  9 10       22 17 12  7  2
 * 11 12 13 14 15  ==>  23 18 13  8  3
 * 16 17 18 19 20       24 19 14  9  4
 * 21 22 23 24 25       25 20 15 10  5
 */
#include <stdio.h>

#define MATRIX 5

void rotate(int target[MATRIX][MATRIX], int arr[MATRIX][MATRIX]) {
  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      arr[j][4-i] = target[i][j];
    }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = {
    { 1, 2, 3, 4, 5 },
    { 6, 7, 8, 9, 10 },
    { 11, 12, 13, 14, 15 },
    { 16, 17, 18, 19, 20 },
    { 21, 22, 23, 24, 25 }
  };
  int rotatedArray[MATRIX][MATRIX] = { 0, };

  rotate(arr, rotatedArray);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", rotatedArray[i][j]);
    }
    printf("\n");
  }
}
