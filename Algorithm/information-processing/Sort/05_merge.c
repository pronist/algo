/**
 * Merge sort
 */
#include <stdio.h>

#define ARRAY_SIZE_A 5
#define ARRAY_SIZE_B 4

void merge_sort(int* first, int* second, int* arr) {
  int index = -1;
  int i = 0;
  int j = ARRAY_SIZE_B - 1;

  int done = 0;

  for (;;) {
    if (first[i] < second[j]) {
      index++;
      arr[index] = first[i];
      i++;
      if (i > ARRAY_SIZE_A-1) {
        if (done == 0) {
          first[ARRAY_SIZE_A-1] = 99999;
          i = ARRAY_SIZE_A-1;
          done = 1;
        } else break;
      }
    } else {
      index++;
      arr[index] = second[j];
      j--;
      if (j < 0) {
        if (done == 0) {
          second[0] = 99999;
          j = 0;
          done = 1;
        } else break;
      }
    }
  }
}

void main() {
  int first[ARRAY_SIZE_A] = { 2, 5, 10, 17, 20 };
  int second[ARRAY_SIZE_B] = { 11, 9, 8, 7 };
  int mergedArray[ARRAY_SIZE_A + ARRAY_SIZE_B] = { 0 };

  merge_sort(first, second, mergedArray);

  for (int i = 0; i < ARRAY_SIZE_A + ARRAY_SIZE_B; i++) {
    printf("%2d ", mergedArray[i]);
  }
}
