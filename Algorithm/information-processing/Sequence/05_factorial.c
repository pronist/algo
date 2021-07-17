/**
 * S = !1 + !2 + !3 ... + 100!
 * Sum of factorial
 */
#include <stdio.h>
#include <math.h>

int get_sum_of_factorial(int end) {
  int sum = 1, An = 1;
  for (int i = 2; i <= end; i++) {
    // An = (n-1)! * i
    An *= i;
    sum += An;
  }
  return sum;
}

void main() {
  printf("%d", get_sum_of_factorial(5));
}
