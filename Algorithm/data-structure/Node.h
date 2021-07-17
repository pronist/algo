#ifndef __NODE__
#define __NODE__

#include <stdlib.h>
#include <stdarg.h>

typedef struct __Node__ Node;

typedef struct __Node__ {
  Node* prev;
  void* element;
  Node* next;
} Node;

Node* create_node(void*);

/**
 * Create a Node
 *
 * @public
 */
Node* create_node(void* element) {
  Node* node = (Node*) malloc(sizeof(Node));
  if (node != NULL) {
    node->prev = NULL;
    node->element = element;
    node->next = NULL;

    return node;
  }
  return NULL;
}

#endif // !__NODE__
