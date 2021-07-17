#ifndef __ARRAY__
#define __ARRAY__

#include <stdlib.h>
#include <stdarg.h>

typedef struct __Array__ {
  size_t length;

  /**
   * @private
   */
  void** _elements;
  int _position;
} Array;

Array array(size_t, ...);

void* ar_at(Array*, int);
void  ar_clear(Array*);

void* ar_begin(Array*);
void* ar_next(Array*);
void* ar_end(Array*);
void* ar_prev(Array*);

int   ar_done(void*);

/**
 * Create Array
 *
 * @public
 */
Array array(size_t length, ...) {
  Array array = { .length = length, ._elements = malloc(sizeof(void*) * length), ._position = 0 };
  va_list ap;
  va_start(ap, length);
  for (int i = 0; i < length; i++) {
    array._elements[i] = va_arg(ap, void*);
  }
  va_end(ap);
  return array;
}

/**
 * Get element with position
 *
 * @public
 */
void* ar_at(Array* array, int position) {
  return array->_elements[position];
}

/**
 * Clear Array
 *
 * @public
 */
void ar_clear(Array* array) {
  free(array->_elements);
}

/**
 * Get first element
 *
 * @public
 */
void* ar_begin(Array* array) {
  array->_position = 0;
  return array->_elements[array->_position];
}

/**
 * Get next element
 *
 * @public
 */
void* ar_next(Array* array) {
  if (array->_position < array->length - 1) {
    return array->_elements[++array->_position];
  } else return NULL;
}

/**
 * Get last element
 *
 * @public
 */
void* ar_end(Array* array) {
  array->_position = array->length - 1;
  return array->_elements[array->_position];
}

/**
 * Get prev element
 *
 * @public
 */
void* ar_prev(Array* array) {
  if (array->_position > 0) {
    return array->_elements[--array->_position];
  } else return NULL;
}

/**
 * is Done
 *
 * @public
 */
int ar_done(void* iter) {
  if (iter == NULL) {
    return 1;
  }
  return 0;
}

#endif // !__ARRAY__
