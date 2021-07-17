/**
 * Magic square Matrix
 *
 * 17 24  1  8 15
 * 23  5  7 14 16
 *  4  6 13 20 22
 * 10 12 19 21  3
 * 11 18 25  2  9
 */
#include <stdio.h>

#define MATRIX 5

void magic_square(int arr[MATRIX][MATRIX]) {
  int row = 0;
  int column = (int) MATRIX/2;

  for (int i = 1; i <= MATRIX*MATRIX; i++) {
    arr[row][column] = i;

    int nextRow = row - 1;
    int nextColumn = column + 1;

    if (nextRow == -1) {
      nextRow = MATRIX-1;
    }
    if (nextColumn == MATRIX) {
      nextColumn = 0;
    }

    if (arr[nextRow][nextColumn] != 0) {
      row++;
    } else {
      row = nextRow;
      column = nextColumn;
    }
  }
}

void main() {
  int arr[MATRIX][MATRIX] = { 0, };

  magic_square(arr);

  for (int i = 0; i < MATRIX; i++) {
    for (int j = 0; j < MATRIX; j++) {
      printf("%2d ", arr[i][j]);
    }
    printf("\n");
  }
}
