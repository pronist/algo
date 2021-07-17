/*
 * Deque
 */
#ifndef __DEQUE__
#define __DEQUE__

#include <stdlib.h>
#include <stdarg.h>

typedef struct __Deque__ {
  size_t length;

  /**
   * @private
   */
  void** _elements;
  int _position;
} Deque;

Deque deque(size_t, ...);

void  dq_push_front(Deque*, void*);
void  dq_pop_front(Deque*);
void  dq_push_back(Deque*, void*);
void  dq_pop_back(Deque*);

void  dq_insert(Deque*, void*, int);
void  dq_erase(Deque*, int);

void* dq_at(Deque*, int);
void  dq_clear(Deque*);

void* dq_begin(Deque*);
void* dq_next(Deque*);
void* dq_end(Deque*);
void* dq_prev(Deque*);

int   dq_done(void*);

/**
 * @private
 */
void __realloc(Deque* deque, int length) {
  deque->_elements = (void**) realloc(deque->_elements, sizeof(void*) * length);
}

/**
 * Create Deque
 *
 * @public
 */
Deque deque(size_t length, ...) {
  Deque dq = { .length = 0, ._elements = malloc(sizeof(void*) * length), ._position = 0 };
  va_list ap;
  va_start(ap, length);
  for (int i = 0; i < length; i++) {
    dq_push_back(&dq, va_arg(ap, void*));
  }
  va_end(ap);
  return dq;
}

/**
 * Push new node at Front
 *
 * @public
 */
void dq_push_front(Deque* deque, void* element) {
  __realloc(deque, ++deque->length);
  for (int i = deque->length-2; i >= 0; i--) {
    deque->_elements[i+1] = deque->_elements[i];
  }
  deque->_elements[0] = element;
}

/**
 * Remove node at Front
 *
 * @public
 */
void dq_pop_front(Deque* deque) {
  for (int i = 1; i < deque->length; i++) {
    deque->_elements[i-1] = deque->_elements[i];
  }
  __realloc(deque, --deque->length);
}

/**
 * Push new node at Back
 *
 * @public
 */
void dq_push_back(Deque* deque, void* element) {
  __realloc(deque, ++deque->length);
  deque->_elements[deque->length-1] = element;
}

/**
 * Remove node at Back
 *
 * @public
 */
void dq_pop_back(Deque* deque) {
  __realloc(deque, --deque->length);
}

/**
 * Push element at Position
 *
 * @public
 */
void dq_insert(Deque* deque, void* element, int position) {
  __realloc(deque, ++deque->length);
  for (int i = deque->length - 2; i >= position; i--) {
    deque->_elements[i + 1] = deque->_elements[i];
  }
  deque->_elements[position] = element;
}

/**
 * Remove element at Position
 *
 * @public
 */
void dq_erase(Deque* deque, int position) {
  for (int i = position + 1; i < deque->length; i++) {
    deque->_elements[i - 1] = deque->_elements[i];
  }
  __realloc(deque, --deque->length);
}


/**
 * Get element with position
 *
 * @public
 */
void* dq_at(Deque* deque, int position) {
  return deque->_elements[position];
}

/**
 * Clear deque
 *
 * @public
 */
void dq_clear(Deque* deque) {
  free(deque->_elements);
}

/**
 * Get first element
 *
 * @public
 */
void* dq_begin(Deque* deque) {
  deque->_position = 0;
  return deque->_elements[deque->_position];
}

/**
 * Get next element
 *
 * @public
 */
void* dq_next(Deque* deque) {
  if (deque->_position < deque->length - 1) {
    return deque->_elements[++deque->_position];
  } else return NULL;
}

/**
 * Get last element
 *
 * @public
 */
void* dq_end(Deque* deque) {
  deque->_position = deque->length - 1;
  return deque->_elements[deque->_position];
}

/**
 * Get prev element
 *
 * @public
 */
void* dq_prev(Deque* deque) {
  if (deque->_position > 0) {
    return deque->_elements[--deque->_position];
  } else return NULL;
}

/**
 * is Done
 *
 * @public
 */
int dq_done(void* iter) {
  if (iter == NULL) {
    return 1;
  }
  return 0;
}

#endif // !__DEQUE__
