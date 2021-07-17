/**
 * S = 2, 6, 18, 54, 162 ...
 * Sum of geometric
 */
#include <stdio.h>
#include <math.h>

int get_sum_of_geometric(int begin, int end, int diff) {
  int sum = begin;
  for (int i = 2; i <= end; i++) {
    // An = A1 + R^(n-1)
    sum += begin * pow(diff, i-1);
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_geometric(2, 5, 3));
}
