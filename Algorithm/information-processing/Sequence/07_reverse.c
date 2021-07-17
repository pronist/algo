/**
 * S = 1 - 2 + 3 - 4 + 5 - 6 ... - 10
 * Sum of reverse
 */
#include <stdio.h>
#include <math.h>

int get_sum_of_reverse(int begin, int end) {
  int sum = 0, s = 1;
  for (int i = begin; i <= end; i++) {
    sum += (i * s);
    s = -s;
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_reverse(1, 5));
}
