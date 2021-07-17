/*
 * Queue
 */
#ifndef __QUEUE__
#define __QUEUE__

#include "List.h"

typedef struct __Queue__ {
  size_t length;

  /**
   * @private
   */
  List _ls;
} Queue;

Queue queue(size_t, ...);

void  qu_push(Queue*, void*);
void  qu_pop(Queue*);

void  qu_clear(Queue*);

Node* qu_front(Queue*);
Node* qu_rear(Queue*);

/**
 * Create Queue
 *
 * @public
 */
Queue queue(size_t length, ...) {
  List ls = list(0);
  Queue qu = { .length = 0, ._ls = ls };
  va_list ap;
  va_start(ap, length);
  for (int i = 0; i < length; i++) {
    qu_push(&qu, va_arg(ap, void*));
  }
  va_end(ap);
  return qu;
}

/**
 * Push new node at Back
 *
 * @public
 */
void qu_push(Queue* queue, void* element) {
  ls_push_back(&queue->_ls, element);
  queue->length = queue->_ls.length;
}

/**
 * Remove node at Front
 *
 * @public
 */
void qu_pop(Queue* queue) {
  ls_pop_front(&queue->_ls);
}

/**
 * Clear Queue
 *
 * @public
 */
void qu_clear(Queue* queue) {
  ls_clear(&queue->_ls);
}

/**
 * Get front
 *
 * @public
 */
Node* qu_front(Queue* queue) {
  return ls_begin(&queue->_ls);
}

/**
 * Get rear
 *
 * @public
 */
Node* qu_rear(Queue* queue) {
  return ls_end(&queue->_ls);
}

#endif // !__QUEUE__
