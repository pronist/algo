/**
 * Matrix multiply
 *
 *  3  0  8        5  7 -3  4   => -33 11 33
 * -5  1 -1    X   2 -5  3  6       64 31 51
 *  7  4  4
 *  2  4  3
 */
#include <stdio.h>

#define ARRAY_FIRST_ROW 2
#define ARRAY_FIRST_COLUMN 4

#define ARRAY_SECOND_ROW 4
#define ARRAY_SECOND_COLUMN 3

void multiply(
  int first[ARRAY_FIRST_ROW][ARRAY_FIRST_COLUMN],
  int second[ARRAY_SECOND_ROW][ARRAY_SECOND_COLUMN],
  int arr[ARRAY_FIRST_ROW][ARRAY_SECOND_COLUMN]
) {
  for (int i = 0; i < ARRAY_FIRST_ROW; i++) {
    for (int j = 0; j < ARRAY_SECOND_COLUMN; j++) {
      int sum = 0;
      for (int q = 0; q < ARRAY_SECOND_ROW; q++) {
        sum += first[i][q] * second[q][j];
      }
      arr[i][j] = sum;
    }
  }
}

void main() {
  int first[ARRAY_FIRST_ROW][ARRAY_FIRST_COLUMN] = {
    { 5, 7, -3, 4 },
    { 2, -5, 3, 6 }
  };
  int second[ARRAY_SECOND_ROW][ARRAY_SECOND_COLUMN] = {
    { 3, 0, 8 },
    { -5 , 1, -1 },
    { 7, 4, 4 },
    { 2, 4, 3}
  };
  int multipliedArray[ARRAY_FIRST_ROW][ARRAY_SECOND_COLUMN] = { 0, };

  multiply(first, second, multipliedArray);

  for (int i = 0; i < ARRAY_FIRST_ROW; i++) {
    for (int j = 0; j < ARRAY_SECOND_COLUMN; j++) {
      printf("%2d ", multipliedArray[i][j]);
    }
    printf("\n");
  }
}
