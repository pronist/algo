/**
 * If jan, first day is Tuesday,
 * Input month, Day -> Week day
 */
#include <stdio.h>

char* get_date(int month, int day) {
  int md[12] = { 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 };
  char* wd[7] = { "Mon", "Tue", "Wen", "Thr", "Fri", "Sat", "Sun" };

  int tDay = 0;
  int tMonth = month - 1;
  int xDay = 0;

  int week = 0;

  for (int i = 0; i < tMonth; i++) {
    tDay += md[i];
  }
  tDay += day-1;

  week = (int) tDay / 7;
  week *= 7;

  week = tDay - week;

  xDay = week + 1;
  if (xDay > 7) {
    xDay -= 7;
  }

  return wd[xDay];
}

void main() {
  int month = 1;
  int day = 31;

  printf("%s", get_date(month, day));
}
