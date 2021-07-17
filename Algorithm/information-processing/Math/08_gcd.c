/**
 * GCD
 */
#include <stdio.h>

int get_GCD(int x, int y) {
  if (x < y) {
    get_GCD(y, x);
  } else if (x >= y && x % y == 0) {
    return y;
  } else {
    get_GCD(y, x % y);
  }
}

void main() {
  printf("%d", get_GCD(60, 124));
}
