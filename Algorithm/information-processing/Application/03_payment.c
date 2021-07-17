/**
 * papers -> 10000, 5000, 1000, 500, 100, 50, 10, 5, 1
 * Input name, payment -> Classification each paper from payment
 */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define PAYMENT 9
#define PERSON 3

typedef struct __Person__ {
  char name[PERSON];
  int payment;
  int papers[PAYMENT];
} Person;

void get_papers(Person* persons, int* payments) {
  for (int i = 0; i < PERSON; i++) {
    int payment = persons[i].payment;
    for (int j = 0 ; j < PAYMENT; j++) {
      persons[i].papers[j] = (int) payment / payments[j];
      payment = (int) payment % payments[j];
    }
  }
}

void main() {
  int payments[PAYMENT] = { 10000, 5000, 1000, 500, 100, 50, 10, 5, 1 };
  int totalPayments[PAYMENT] = { 0 };

  Person person1 = { .name = "A", .payment = 532263 };
  Person person2 = { .name = "B", .payment = 307349 };
  Person person3 = { .name = "C", .payment = 152830 };

  Person persons[PERSON] = { person1, person2, person3 };

  get_papers(persons, payments);

  printf("Name  Payment  ");
  for (int i = 0; i < PAYMENT; i++) {
    printf("%-5d  ", payments[i]);
  }
  printf("\n-----------------------------------------------------------------------------\n");
  for (int i = 0; i < PERSON; i++) {
    printf("%-4s  %-7d  ", persons[i].name, persons[i].payment);
    for (int j = 0; j < PAYMENT; j++) {
      printf("%-5d  ", persons[i].papers[j]);
      totalPayments[j] += persons[i].papers[j];
    }
    printf("\n");
  }
  printf("-----------------------------------------------------------------------------\n");
  printf("%-15s", "Payments:");
  for (int i = 0; i < PAYMENT; i++) {
    printf("%-5d  ", totalPayments[i]);
  }
}
