/**
 * Compound interest -> 5%
 */
#include <stdio.h>
#include <math.h>

#define ARRAY_SIZE 10

void get_money(int money, int period, int* output) {
  for (int i = 0; i < period; i++) {
    money = (int) money * (1 + 0.05);
    output[i] = money;
  }
}

void main() {
  int money = 100000;
  int period = 5;
  int output[ARRAY_SIZE] = { 0 };

  get_money(money, period, output);

  for (int i = 0; i < ARRAY_SIZE; i++) {
    printf("%d ", output[i]);
  }
}
