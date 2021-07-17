/*
 * Stack
 */
#ifndef __STACK__
#define __STACK__

#include "List.h"

typedef struct __Stack__ {
  size_t length;

  /**
   * @private
   */
  List _ls;
} Stack;

Stack stack(size_t, ...);

void  st_push(Stack*, void*);
void  st_pop(Stack*);

void  st_clear(Stack*);

Node* st_top(Stack*);

/**
 * Create Stack
 *
 * @public
 */
Stack stack(size_t length, ...) {
  List ls = list(0);
  Stack st = { .length = 0, ._ls = ls };
  va_list ap;
  va_start(ap, length);
  for (int i = 0; i < length; i++) {
    st_push(&st, va_arg(ap, void*));
  }
  va_end(ap);
  return st;
}

/**
 * Push new node at Back
 *
 * @public
 */
void st_push(Stack* stack, void* element) {
  ls_push_back(&stack->_ls, element);
  stack->length = stack->_ls.length;
}

/**
 * Remove node at Back
 *
 * @public
 */
void st_pop(Stack* stack) {
  ls_pop_back(&stack->_ls);
}

/**
 * Clear Stack
 *
 * @public
 */
void st_clear(Stack* stack) {
  ls_clear(&stack->_ls);
}

/**
 * Get top element
 *
 * @public
 */
Node* st_top(Stack* stack) {
  return ls_end(&stack->_ls);
}

#endif // !__STACK__
