/**
 * S = 2, 8, 14, 20, 26
 * Sum of arithmetic
 */
#include <stdio.h>

int get_sum_of_arithmetic(int begin, int end, int diff) {
  int sum = begin;
  for (int i = 2; i <= end; i++) {
    // An = A1 + (N-1) * 6
    sum += begin + (i-1) * diff;
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_arithmetic(2, 5, 6));
}
