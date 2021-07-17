/**
 * Payment = Basic plan payment + (Calltime * Payment per minute) + (Discount * Discount per minute)
 *
 * Customer N(i)
 *
 * N(i,0) = Basic plan payment
 * N(i,1) = Calltime
 * N(i,2) = Discount
 *
 * Payment per seconds = 50, Discount per seconds = 25.
 */
#include <stdio.h>

#define ROW 5
#define COLUMN 3

void get_max_time(int customers[ROW][COLUMN], int* bill, int* time) {
  for (int i = 0; i < ROW; i++) {
    int tTime = customers[i][1] + customers[i][2];
    int tBill = customers[i][0];
    tBill += customers[i][1] * 50;
    tBill += customers[i][2] * 25;

    if (tBill > *bill) {
      *bill = tBill;
      *time = tTime;
    }
  }
}

void main() {
  int customers[ROW][COLUMN] = {
    { 25000, 10, 25 },
    { 35280, 50, 20 },
    { 40000, 40, 22 },
    { 12500, 25, 43 },
    { 10000, 100, 20 }
  };
  int bill = -999999;
  int time = -999999;

  get_max_time(customers, &bill, &time);

  printf("bill: %d, time: %d", bill, time);
}
