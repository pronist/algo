/**
 * S = (10*1)^2 + (9*2)^2 ... + (1*10)^2
 * Sum of product
 */
#include <stdio.h>
#include <math.h>

int get_sum_of_product(int first, int second) {
  int sum = 0, end = first - second;
  for (int i = 0; i <= end; i++) {
    sum += pow(first-- * second++, 2);
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_product(10, 1));
}
