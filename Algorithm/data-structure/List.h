/*
 * Doubly linked list
 */
#ifndef __LINKED_LIST__
#define __LINKED_LIST__

#include "Node.h"

typedef struct __List__ {
  size_t length;

  Node* front;
  Node* back;
} List;

List list(size_t, ...);

void  ls_push_front(List*, void*);
void  ls_pop_front(List*);
void  ls_push_back(List*, void*);
void  ls_pop_back(List*);

void  ls_insert(List*, Node*, void*);
void  ls_erase(List*, Node*);

Node* ls_at(List*, void*);
void  ls_clear(List*);

Node* ls_begin(List*);
Node* ls_end(List*);
Node* ls_next(Node*);
Node* ls_prev(Node*);

int   ls_done(Node*);

/**
 * Create Linked List
 *
 * @public
 */
List list(size_t length, ...) {
  List ls = { .length = 0, .front = NULL, .back = NULL };
  va_list ap;
  va_start(ap, length);
  for (int i = 0; i < length; i++) {
    ls_push_back(&ls, va_arg(ap, void*));
  }
  va_end(ap);
  return ls;
}

/**
 * Push new node at Front
 *
 * @public
 */
void ls_push_front(List* list, void* element) {
  Node* node = create_node(element);
  if (list->front != NULL) {
    node->next = list->front;
    node->prev = NULL;
    list->front->prev = node;
  }
  else if (list->front == NULL && list->back == NULL) {
    list->back = node;
  }
  list->front = node;
  list->length++;
}

/**
 * Remove node at Front
 *
 * @public
 */
void ls_pop_front(List* list) {
  Node* node = list->front;
  if (list->front->next != NULL) {
    list->front = list->front->next;
    list->front->prev = NULL;
  }
  else {
    list->back = NULL; list->front = NULL;
  }
  free(node);
  list->length--;
}

/**
 * Push new node at Back
 *
 * @public
 */
void ls_push_back(List* list, void* element) {
  Node* node = create_node(element);
  if (list->back != NULL) {
    node->prev = list->back;
    node->next = NULL;
    list->back->next = node;
  }
  else if (list->front == NULL && list->back == NULL) {
    list->front = node;
  }
  list->back = node;
  list->length++;
}

/**
 * Remove node at Back
 *
 * @public
 */
void ls_pop_back(List* list) {
  Node* node = list->back;
  if (list->back->prev != NULL) {
    list->back = list->back->prev;
    list->back->next = NULL;
  }
  else {
    list->back = NULL; list->front = NULL;
  }
  free(node);
  list->length--;
}

/**
 * Insert new node at before the node
 *
 * @public
 */
void ls_insert(List* list, Node* iter, void* element) {
  if (iter == list->front) {
    ls_push_front(list, element);
  }
  else {
    Node* node = create_node(element);

    node->prev = iter->prev;
    node->next = iter;
    iter->prev->next = node;
    iter->prev = node;
    list->length++;
  }
}

/**
 * Remove node
 *
 * @public
 */
void ls_erase(List* list, Node* iter) {
  if (iter == list->front) {
    ls_pop_front(list);
  }
  else if (iter == list->back) {
    ls_pop_back(list);
  }
  else {
    iter->prev->next = iter->next;
    iter->next->prev = iter->prev;

    free(iter);
    list->length--;
  }
}

/**
 * Get node with element
 *
 * @public
 */
Node* ls_at(List* list, void* element) {
  for (Node* iter = ls_begin(list); 1 != ls_done(iter); iter = ls_next(iter)) {
    if (iter->element == element) {
      return iter;
    }
  }
  return NULL;
}

/**
 * Clear Linked List
 *
 * @public
 */
void ls_clear(List* list) {
  size_t s = list->length;
  for (size_t i = 0; i < s; i++) {
    ls_pop_back(list);
  }
}

/**
 * Get front node
 *
 * @public
 */
Node* ls_begin(List* list) {
  return list->front;
}

/**
 * Get next node
 *
 * @public
 */
Node* ls_next(Node* iter) {
  return iter->next;
}

/**
 * Get back node
 *
 * @public
 */
Node* ls_end(List* list) {
  return list->back;
}

/**
 * Get prev node
 *
 * @public
 */
Node* ls_prev(Node* iter) {
  return iter->prev;
}

/**
 * is Done
 *
 * @public
 */
int ls_done(Node* iter) {
  if (iter == NULL) {
    return 1;
  }
  return 0;
}

#endif // !__LINKED_LIST__
