/**
 * Binary to Decimal
 *
 * Left bit is Sign
 * 0 => Positive, 1 => Negative
 */
#include <stdio.h>
#include <math.h>

#define ARRAY_SIZE 5

int to_decimal(int* input) {
  int decimal = 0;
  for (int i = 1; i < ARRAY_SIZE; i++) {
    decimal += pow(2, ARRAY_SIZE - i - 1) *  input[i];
  }
  return decimal;
}

/**
 * 2complement (don't care sign)
 */
void to_binary(int* input) {
  for (int i = ARRAY_SIZE - 1; i >= 1 ; i--) {
    if ((input[i] - 1) == -1) {
      input[i] = 1;
    } else {
      input[i] = 0; break;
    }
  }
  for (int i = 1; i < ARRAY_SIZE; i++) {
    input[i] = 1 - input[i];
  }
}

void main() {
  // int input[ARRAY_SIZE] = { 0, 1, 0, 1, 1 }; // 11
  int input[ARRAY_SIZE] = { 1, 0, 1, 0, 1 }; // -11
  int sign = 1;

  if (input[0] > 0) {
    sign = -sign;
    to_binary(input);
  };

  printf("%d", to_decimal(input) * sign);
}
