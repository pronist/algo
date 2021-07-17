/**
 * S = 1, 1, 2, 3, 5, 8, 13 ...
 * Sum of fibonacci
 */
#include <stdio.h>
#include <math.h>

int get_sum_of_fibonacci(int first, int second, int end) {
  int curr = 0, sum = first + second;
  for (int i = 3; i <= end; i++) {
    // An = An-2 + An-1
    curr = first + second;
    sum += curr;
    first = second;
    second = curr;
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_fibonacci(1, 1, 5));
}
