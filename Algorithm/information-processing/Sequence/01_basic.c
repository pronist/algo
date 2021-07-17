/**
 * Sum
 */
#include <stdio.h>

int get_sum(int end) {
  int sum = 0;
  for (int i = 1; i <= end; i++) {
    sum += i;
  }
  return sum;
}

void main() {
  printf("%d", get_sum(5));
}
